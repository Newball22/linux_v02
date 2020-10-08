package routers

import (
	"github.com/gin-gonic/gin"
	v1 "goFlow/api/v1"
	"goFlow/middleware"
	"goFlow/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Cors())

	//使用中间件
	//r.Use(middleware.Logger())

	router := r.Group("api/v1")
	{
		//测试接口网络
		router.GET("ping", v1.TestNet)

		//登录路由
		router.POST("login", v1.Login)

		//用户路由
		router.GET("user/:id", v1.GetUser)
		router.GET("users", v1.GetUsers)
		router.POST("user/add", v1.AddUser)

		//文章路由
		router.GET("article/:id", v1.GetArticle)       //查询一篇文章
		router.GET("articles/:cid", v1.GetCateArticle) //查询某分类下的文章
		router.GET("articles", v1.GetArticles)         //查询所有文章

		//分类路由
		router.GET("category", v1.GetCate)
	}

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		auth.POST("upload", v1.UploadData) //文件上传

		auth.POST("category/add", v1.AddCate)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

	}

	_ = r.Run(utils.HttpPort)
}
