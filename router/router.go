package router

import (
	"myapp/config"
	"myapp/controller/passage"
	"myapp/controller/user"
	"myapp/controller/wechat"
	"myapp/middleware"

	//docs
	_ "myapp/docs"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//Router 路由
func Router(app config.App) *gin.Engine {
	router := gin.New()
	pprof.Register(router)
	// 使用中间件
	// 使用Logger中间件
	router.Use(gin.Logger())
	// 使用Recovery中间件
	router.Use(gin.Recovery())

	if app.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Use(middleware.Cors())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userRouter := router.Group("/user")
	userRouter.GET("list", user.GetUserList)
	userRouter.GET("detail/:id", user.GetDetail)

	passageRouter := router.Group("/passage")
	passageRouter.GET("list", passage.GetPassageList)
	passageRouter.POST("add", passage.AddPassage)

	wechatRouter := router.Group("/wechat")
	wechatRouter.GET("/session", wechat.GetAccession)

	return router
}
