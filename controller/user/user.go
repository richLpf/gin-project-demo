package user

import (
	"fmt"
	"net/http"

	"myapp/dbs"
	"myapp/model"

	"github.com/gin-gonic/gin"
)

//GetUserList get user
func GetUserList(c *gin.Context) {
	var info []model.Users
	if err := dbs.DB.First(&info).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Data": info})
}

//GetDetail update user
func GetDetail(c *gin.Context) {
	id := c.Param("id")
	var info model.Users
	if err := dbs.DB.Where("is_deleted = 0 and id = ?", id).First(&info).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Data": info})
}

//UpdateUser
func UpdateUser(c *gin.Context) {
	var req model.Users
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	if err := dbs.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

//AddUser
func AddUser(c *gin.Context) {
	var req model.Users
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	if err := dbs.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "data": req})
}

// DelUser
func DelUser(c *gin.Context) {
	var req model.Users
	id := c.Param("id")
	fmt.Println("id", id)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	if err := dbs.DB.Debug().Model(&req).Where("id = ? and is_deleted = 0", id).Update("is_deleted", 1).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

//ReqUserRoles
type ReqUserRoles struct {
	Namespace string `json:"namespace"`
	User      string `json:"user"`
	RoleID    []uint `json:"role_id"`
	Status    uint   `json:"status"`
	CreatedBy string `json:"created_by"`
}

//AddRole 给用户添加并给权限
func AddRole(c *gin.Context) {
	var req ReqUserRoles
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	tx := dbs.DB.Begin()
	for _, v := range req.RoleID {
		var info model.UserRoles
		info.Namespace = req.Namespace
		info.User = req.User
		info.RoleID = v
		info.Status = req.Status
		info.CreatedBy = req.CreatedBy
		if err := tx.Save(&info).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

//GetUserPermissioin 获取当前用户的权限信息
func GetUserPermission(c *gin.Context) {
	user := c.DefaultQuery("user", "")
	namespace := c.DefaultQuery("namespace", "")
	category := c.DefaultQuery("category", "")
	fmt.Println("user", user, namespace, category)
	if user == "" || namespace == "" || category == "" {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": "参数不能为空"})
		return
	}
	// 获取当前用户所有的角色id
	var roleIds []int
	if err := dbs.DB.Table("user_roles").Where("is_deleted = 0 AND user = ? AND status = 1", user).Pluck("role_id", &roleIds).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 2, "message": err.Error()})
		return
	}

	// 获取当前角色下所有的资源id
	var resourceIds []int
	if err := dbs.DB.Table("role_resources").Where("is_deleted = 0 AND role_id in (?)", roleIds).Pluck("resource_id", &resourceIds).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 3, "message": err.Error()})
		return
	}
	var res []model.Resources
	// 获取当前用户下所有的资源id信息
	if err := dbs.DB.Debug().Table("resources").Where("is_deleted = 0 AND id in (?)", resourceIds).Find(&res).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 4, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "message": "success", "Data": res})
}
