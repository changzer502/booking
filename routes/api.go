package routes

import (
	"github.com/gin-gonic/gin"
	"registration-booking/app/handler"
	"registration-booking/app/middleware"
	"registration-booking/app/services"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", handler.Register)
	router.POST("/auth/login", handler.Login)
	router.POST("/auth/wx_login", handler.WxLogin)
	authRouter := router.Group("", middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", handler.Info)
		authRouter.POST("/auth/logout", handler.Logout)
		authRouter.POST("/image_upload", handler.ImageUpload)

		card := authRouter.Group("/card")
		{
			card.POST("/create", handler.CreateCard)
			card.POST("/list", handler.GetCardList)
			card.GET("/detail/:id", handler.GetCardById)
			card.POST("/update/:id", handler.UpdateCard)
			card.POST("/delete/:id", handler.DeleteCard)
		}
	}

	department := router.Group("/department")
	{
		department.GET("/list", handler.GetDepartmentList)
		department.POST("/page", handler.GetDepartmentPage)
		department.GET("/:id", handler.GetDepartmentById)
		authDepartment := department.Group("/", middleware.JWTAuth(services.AppGuardName))
		{
			authDepartment.POST("/create", handler.CreateDepartment)
		}
	}

	admin := authRouter.Group("/admin", middleware.JWTAuth(services.AppGuardName))
	{
		admin.POST("/create_doctor", handler.CreateDoctor)
	}
}
