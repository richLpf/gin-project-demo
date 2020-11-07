package resource

import (
	"net/http"

	"myapp/dbs"
	"myapp/model"

	"github.com/gin-gonic/gin"
)

// 获取某个用户的所有资源
// 合并统一用户的所有的资源信息
// 要个列表就行了，不需要详情----
// GetResource Get
// @Summary 获取权限资源
// @Description 获取权限资源
// @Tags ACL
// @Accept json
// @Produce json
// @Param body body model.Resources true  "请求参数"
// @Success 200 {string} string "ok"
// @Router /acl/resource/list [GET]
func GetResource(c *gin.Context) {
	var info model.Resources
	if err := dbs.DB.First(&info).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Data": info})
}

func GetResourceList(c *gin.Context) {
	var list []model.Resources
	if err := dbs.DB.Find(&list).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Data": list})
}

// AddResource Add
// @Summary 添加ACL资源
// @Description 添加ACL资源
// @Tags ACL
// @Accept json
// @Produce json
// @Param body body model.Resources true  "请求参数"
// @Success 200 {string} string "ok"
// @Router /acl/resource/add [POST]
func AddResource(c *gin.Context) {
	var req model.Resources
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

// UpdateResource update
// @Summary 更新ACL资源
// @Description 更新ACL资源
// @Tags ACL
// @Accept json
// @Produce json
// @Param body body model.Resources true  "请求参数"
// @Success 200 {string} string "ok"
// @Router /acl/resource/update [POST]
func UpdateResource(c *gin.Context) {
	var req model.Resources
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

// DelResource delete
// @Summary 删除ACL资源
// @Description 删除ACL资源
// @Tags ACL
// @Accept json
// @Produce json
// @Param body body model.Resources true  "请求参数"
// @Success 200 {string} string "ok"
// @Router /acl/resources/delete [POST]
func DelResource(c *gin.Context) {
	var req model.Resources
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	if err := dbs.DB.Delete(&req).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}
