package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ResponseMessage 返回结构体
type ResponseMessage struct {
	RetCode int                    `json:"RetCode"`
	Message string                 `json:"Message"`
	Data    map[string]interface{} `json:"Data,omitempty"`
}

func response(RetCode int, Message string, Data map[string]interface{}, c *gin.Context) {
	r := ResponseMessage{RetCode, Message, Data}
	c.JSON(http.StatusOK, r)
}

//SuccessResponse 成果返回
func SuccessResponse(data map[string]interface{}, c *gin.Context) {
	response(0, "success", data, c)
}

//ErrorCodeResponse 自定义错误
func ErrorCodeResponse(code int, c *gin.Context) {
	message := GetMsg(code)
	response(code, message, make(map[string]interface{}), c)
}

//FailResponse 错误码
func FailResponse(code int, message string, c *gin.Context) {
	response(code, message, make(map[string]interface{}), c)
}
