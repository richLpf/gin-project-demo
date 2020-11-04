package role

import (
	"net/http"

	"myapp/dbs"
	"myapp/model"

	"github.com/gin-gonic/gin"
)

// 看下结构怎么关联
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
	if err := dbs.DB.Find(&list).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Data": list})
}

// 添加角色，创建接收模型，角色信息和资源信息
// 角色入库，资源信息循环入库
func AddRole(c *gin.Context) {
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
}

// 更新角色信息，同创建一致
func UpdateRole(c *gin.Context) {
	var req model.Roles
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