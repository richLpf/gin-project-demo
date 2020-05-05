package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUserList get user
func GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

//GetDetail update user
func GetDetail(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success", "Id": id})
}
