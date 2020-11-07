package namespace

import (
	"net/http"

	"myapp/dbs"
	"myapp/model"

	"github.com/gin-gonic/gin"
)

// 定义通用的返回值，并写好uuid
//GetNamespaceList get namespace
// @Summary acl项目
// @Description 获取acl项目
// @Tags ACL
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /acl/namespace/list [get]
func GetNamespaceList(c *gin.Context) {
	var list []model.Namespaces
	if err := dbs.DB.Where("is_deleted = 0").Find(&list).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "data": list})
}

// AddNamespace add
// @Summary 添加ACL项目
// @Description 添加ACL项目
// @Tags ACL
// @Accept json
// @Produce json
// @Param body body model.Namespaces true  "请求参数"
// @Success 200 {string} string "ok"
// @Router /acl/namespace/add [POST]
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

// UpdateNamespace Update
// @Summary 更新ACL项目
// @Description 更新ACL项目
// @Tags ACL
// @Accept json
// @Produce json
// @Param body body model.Namespaces true  "请求参数"
// @Success 200 {string} string "ok"
// @Router /acl/namespace/update [POST]
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

// DelNamespace delete
// @Summary 删除ACL项目
// @Description 删除ACL项目
// @Tags ACL
// @Accept json
// @Produce json
// @Param body body model.Namespaces true  "请求参数"
// @Success 200 {string} string "ok"
// @Router /acl/namespace/delete [POST]
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
