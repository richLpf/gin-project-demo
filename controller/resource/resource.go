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
func DelResource(c *gin.Context) {
	var req model.Resources
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	if err := dbs.DB.Debug().Delete(&req).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}
