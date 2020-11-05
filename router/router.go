package router

import (
	"myapp/config"
	"myapp/controller/namespace"
	"myapp/controller/passage"
	"myapp/controller/resource"
	"myapp/controller/role"
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

	//acl
	aclRouter := router.Group("/acl")
	//所有用户的users和权限没有关系
	aclRouter.POST("/user/add", user.AddUser)
	aclRouter.POST("/user/delete/:id", user.DelUser)
	aclRouter.POST("/user/update", user.UpdateUser)
	aclRouter.GET("/user/detail/:id", user.GetDetail)
	aclRouter.GET("/user/list", user.GetUserList)
	// 用户权限
	//aclRouter.POST("/user/addrole", user.AddRole)
	aclRouter.GET("/user/permission", user.GetUserPermission)
	aclRouter.GET("/user/rolelist", user.GetUserRole)
	aclRouter.POST("/user/role/add", user.AddUserRole)
	//role
	aclRouter.POST("/role/add", role.AddRole)
	aclRouter.POST("/role/delete", role.DelRole)
	aclRouter.POST("/role/update", role.UpdateRole)
	aclRouter.GET("/role/detail/:id", role.GetRole)
	aclRouter.GET("/role/list", role.GetRoleList)
	aclRouter.POST("/role/permission", role.RelateResource)
	aclRouter.GET("/role/permission/:id", role.GetRelateResource)
	//resource
	aclRouter.POST("/resource/add", resource.AddResource)
	aclRouter.POST("/resource/delete", resource.DelResource)
	aclRouter.POST("/resource/update", resource.UpdateResource)
	//aclRouter.GET("/resource/detail/:id", resource.GetResource)
	aclRouter.GET("/resource/list", resource.GetResourceList)
	//namespace
	aclRouter.POST("/namespace/add", namespace.AddNamespace)
	aclRouter.POST("/namespace/delete", namespace.DelNamespace)
	aclRouter.POST("/namespace/update", namespace.UpdateNamespace)
	//aclRouter.GET("/namespace/detail/:id", namespace.GetNamespace)
	aclRouter.GET("/namespace/list", namespace.GetNamespaceList)
	return router
}
