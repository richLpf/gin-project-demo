package passage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetPassageList get passage
// List godoc
// @Summary 列表实例
// @Description 描述信息
// @Tags 文章
// @Accept json
// @Produce json
// @Param limit query string false  "20"
// @Param offset query string false  "0"
// @Success 200 {string} string "ok"
// @Router /ucloud/list [get]
func GetPassageList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

// AddPassage update passage
// AddUser godoc
// @Summary 添加文章
// @Description 添加文章
// @Tags 文章
// @Accept json
// @Produce json
// @Param body body model.Passage true  "请求参数"
// @Success 200 {string} string "ok"
// @Router /passage/add [post]
func AddPassage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}
