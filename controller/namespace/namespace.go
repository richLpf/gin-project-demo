package namespace

import (
	"net/http"

	"myapp/dbs"
	"myapp/model"

	"github.com/gin-gonic/gin"
)

// 定义通用的返回值，并写好uuid

func GetNamespaceList(c *gin.Context) {
	var list []model.Namespaces
	if err := dbs.DB.Where("is_deleted = 0").Find(&list).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "data": list})
}
func AddNamespace(c *gin.Context) {
	var req model.Namespaces
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	if err := dbs.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 2, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Data": req})
}
func UpdateNamespace(c *gin.Context) {
	var req model.Namespaces
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 2, "message": err.Error()})
		return
	}
	if err := dbs.DB.Where("is_deleted = 0").Save(&req).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 3, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}
func DelNamespace(c *gin.Context) {
	var req model.Namespaces
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
