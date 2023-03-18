package router

import (
	"myfile/api"
	"myfile/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api")
	{
		v1.GET("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		//获取用户信息
		v1.GET("user/info", api.UserInfo)

		//文件相关路由
		v1.POST("file/file", api.FileUpload)
		v1.GET("file/file", api.FileList)
		v1.GET("file/file/:id", api.FileInfo)

		//文件夹相关路由
		v1.POST("file/folder", api.FolderCreate)
		v1.PUT("file/folder/:id", api.FolderUpdate)
		v1.GET("file/folder/:id", api.FolderInfo)
		v1.GET("file/folder", api.FolderList)

		// 需要登录保护的
		//auth := v1.Group("")
		//auth.Use(middleware.AuthRequired())
		//{
		// User Routing
		//auth.GET("user/me", api.UserMe)
		//auth.DELETE("user/logout", api.UserLogout)
		//}
	}
	return r
}
