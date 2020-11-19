package role

import (
	"net/http"

	"myapp/dbs"
	"myapp/model"

	"github.com/gin-gonic/gin"
	"fmt"
)

// 获取角色，同时要获取角色关联的资源信息
func GetRole(c *gin.Context) {
	var info model.Roles
	if err := dbs.DB.First(&info).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Data": info})
}

func GetRoleList(c *gin.Context) {
	var list []model.Roles
	if err := dbs.DB.Where("is_deleted = 0").Order("created_at desc").Find(&list).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Data": list})
}

// 添加角色，创建接收模型，角色信息和资源信息
// 角色入库，资源信息循环入库
/*func AddRole(c *gin.Context) {
	var req model.Roles
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	if err := dbs.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}*/

//ReqUserRoles
type ReqUserRoles struct {
	ID        uint `json:"id"` 
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Role      string `json:"role"`
	Describe   string  `json:"describe"`
	Permission []uint  `json:"permission"`
	Status    uint   `json:"status"`
	CreatedBy string `json:"created_by"`
}

//AddRole 添加角色
func AddRole(c *gin.Context) {
	curUser := c.MustGet("submitUser").(string)
	if curUser == "" {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "Message": "no user info"})
		return
	}
	var req ReqUserRoles
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "Message": err.Error()})
		return
	}
	tx := dbs.DB.Begin()
	// 首先插入role角色列表
	roleReq := model.Roles{
		Name: req.Name,
		Namespace: req.Namespace,
		Role: req.Role,
		Describe: req.Describe,
		CreatedBy: curUser,
	}
	if err := tx.Save(&roleReq).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{"RetCode": 2, "Message": err.Error()})
		return
	}
	fmt.Println("插入的RoleReq", roleReq)
	// 插入role_resources列表
	for _, v := range req.Permission {
		info := model.RoleResources{
			Namespace: req.Namespace,
			RoleID: roleReq.ID,
			ResourceID: uint(v),
			Describe: req.Describe,
			CreatedBy: curUser,
		}
		if err := tx.Save(&info).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{"RetCode": 3, "Message": err.Error()})
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

// 更新角色信息，同创建一致
func UpdateRole(c *gin.Context) {
	curUser := c.MustGet("submitUser").(string)
	if curUser == "" {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "Message": "no user info"})
		return
	}
	var req ReqUserRoles
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 2, "message": err.Error()})
		return
	}
	tx := dbs.DB.Begin()
	roleReq := model.Roles{
		ID: req.ID,
		Name: req.Name,
		Namespace: req.Namespace,
		Role: req.Role,
		Describe: req.Describe,
		CreatedBy: curUser,
	}
	if err := tx.Save(&roleReq).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{"RetCode": 3, "message": err.Error()})
		return
	}
	// 插入role_resources列表
	for _, v := range req.Permission {
		info := model.RoleResources{
			Namespace: req.Namespace,
			RoleID: req.ID,
			ResourceID: uint(v),
			Describe: req.Describe,
			CreatedBy: curUser,
		}
		if err := tx.Save(&info).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{"RetCode": 4, "Message": err.Error()})
			return
		}
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

func DelRole(c *gin.Context) {
	var req model.Roles
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	if err := dbs.DB.Model(&req).Where("is_deleted = 0").Update("is_deleted", 1).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

// RelateResource 角色关联资源
func RelateResource(c *gin.Context) {
	var req model.ReqRoleResources
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	tx := dbs.DB.Begin()
	for _, v := range req.RolePermission {
		var info model.RoleResources
		info.RoleID = req.RoleID
		info.Namespace = req.Namespace
		info.ResourceID = v.ResourceID
		info.Describe = v.Describe
		info.CreatedBy = req.CreatedBy
		if err := tx.Create(&info).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

// GetRelateResource 获取所有的资源
func GetRelateResource(c *gin.Context) {
	var res []model.RoleResources
	id := c.Param("id")
	if err := dbs.DB.Where("is_deleted = 0 AND role_id = ?", id).Find(&res).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Data": res})
}
