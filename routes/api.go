package routes

import (
	"github.com/gin-gonic/gin"
	"registration-booking/app/handler"
	"registration-booking/app/middleware"
	"registration-booking/app/services"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/test/GetDoctorList", handler.GetDoctorList)

	router.POST("/auth/register", handler.Register)
	router.POST("/auth/login", handler.Login)
	router.POST("/auth/wx_login", handler.WxLogin)
	authRouter := router.Group("", middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", handler.Info)
		authRouter.POST("/auth/logout", handler.Logout)
		authRouter.GET("/user", handler.UserInfoAndRole)
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
	admin.POST("/create_doctor", handler.CreateDoctor)
	doctor := router.Group("/doctor")
	doctor.GET("/list/:department_id", handler.GetDoctorList)
	doctor.GET("/:id", handler.GetDoctorById)

	// 预约
	scheduleAuth := authRouter.Group("/schedule", middleware.JWTAuth(services.AppGuardName))
	scheduleAuth.POST("/create", handler.CreateSchedule)
	scheduleAuth.POST("/ticket/create", handler.CreateTicket)
	scheduleAuth.POST("/ticket/booking", handler.Booking)
	scheduleAuth.GET("/ticket/info/:ticket_id", handler.GetInfoByTicketId)

	scheduleAuth.POST("/booking/history", handler.BookingHistory)
	scheduleAuth.GET("/booking/history/:booking_id", handler.GetBookingHistoryById)
	schedule := authRouter.Group("/schedule")
	schedule.POST("/list", handler.GetScheduleList)

	//文章
	article := authRouter.Group("/article")
	article.POST("/list", handler.GetArticleList)
	article.GET("/:id", handler.GetArticleById)
	articleAuth := article.Group("/", middleware.JWTAuth(services.AppGuardName))
	articleAuth.POST("/create", handler.CreateArticle)
	articleAuth.POST("/update", handler.UpdateArticle)
	articleAuth.POST("/delete/:id", handler.DeleteArticle)
}
