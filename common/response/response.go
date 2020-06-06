package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Response struct {
	C *gin.Context
}

//ResponseMessage 返回结构体
type ResponseMessage struct {
	RetCode int         `json:"RetCode"`
	Message string      `json:"Message"`
	Data    map[string]interface{} `json:"Data,omitempty"`
}

//ResponseMessageWithTotal 返回数量
type ResponseMessageWithTotal struct {
	ResponseMessage
	Total uint64 `json:"total"`
}

func (r *Response)SuccessResponse() {
	r.C.JSON(http.StatusOK, Response{
		RetCode: 0,
		Message: "success",
		Data: data
	})
}
