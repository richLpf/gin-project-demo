package passage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetPassageList get passage
func GetPassageList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}

//AddPassage update passage
func AddPassage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "Msg": "success"})
}
