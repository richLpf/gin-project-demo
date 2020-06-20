package wechat

import (
	"myapp/common/response"

	"github.com/gin-gonic/gin"
)

//GetAccession 获取accession
func GetAccession(c *gin.Context) {

	//config := config.IniInfo{}
	//config.SetIniInfo()

	res := map[string]interface{}{
		"list": "test",
	}
	response.SuccessResponse(res, c)
}
