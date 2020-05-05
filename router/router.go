package router

import (
	"myapp/controller/passage"
	"myapp/controller/user"
	"myapp/middleware"

	"github.com/gin-gonic/gin"
)

//Router 路由
func Router() *gin.Engine {
	router := gin.New()
	// 使用中间件
	// 使用Logger中间件
	router.Use(gin.Logger())
	// 使用Recovery中间件
	router.Use(gin.Recovery())

	router.Use(middleware.Cors())

	userRouter := router.Group("/user")
	userRouter.GET("list", user.GetUserList)
	userRouter.GET("detail/:id", user.GetDetail)

	passageRouter := router.Group("/passage")
	passageRouter.GET("list", passage.GetPassageList)
	passageRouter.POST("add", passage.AddPassage)

	return router
}
