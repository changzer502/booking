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
		authRouter.POST("/auth/refreshToken", handler.RefreshToken)
		authRouter.POST("/auth/info", handler.Info)
		authRouter.POST("/auth/logout", handler.Logout)
		authRouter.GET("/user", handler.UserInfoAndRole)
		authRouter.POST("/image_upload", handler.ImageUpload)

		card := authRouter.Group("/card")
		{
			card.POST("/create", handler.CreateCard)
			card.POST("/list", handler.GetCardList)
			card.POST("/all_list", handler.GetAllCardList)
			card.GET("/detail/:id", handler.GetCardById)
			card.POST("/update/:id", handler.UpdateCard)
			card.POST("/delete/:id", handler.DeleteCard)
		}
	}

	department := router.Group("/department")
	{
		department.GET("/list", handler.GetDepartmentList)
		department.POST("/page", handler.GetDepartmentPage)
		department.POST("/parent/page", handler.GetParentDepartmentPage)
		department.GET("/:id", handler.GetDepartmentById)
		authDepartment := department.Group("/", middleware.JWTAuth(services.AppGuardName))
		{
			authDepartment.POST("/create", handler.CreateDepartment)
			authDepartment.POST("/update", handler.UpdateDepartment)
			authDepartment.POST("/delete/:department_id", handler.DeleteDepartment)
		}
	}

	admin := authRouter.Group("/admin", middleware.JWTAuth(services.AppGuardName))
	admin.POST("/create_doctor", handler.CreateDoctor)
	admin.POST("/disease_list", handler.GetDiseaseList)
	doctor := router.Group("/doctor")
	doctor.GET("/list/:department_id", handler.GetDoctorList)
	doctor.POST("/list", handler.GetAllDoctorList)
	doctor.GET("/:id", handler.GetDoctorById)
	user := authRouter.Group("/user")
	user.POST("/update", handler.UpdateUser)
	user.POST("/update/doctor", handler.UpdateDoctor)
	user.POST("/delete/:id", handler.DeleteUser)

	// 预约
	scheduleAuth := authRouter.Group("/schedule", middleware.JWTAuth(services.AppGuardName))
	scheduleAuth.POST("/create", handler.CreateSchedule)
	scheduleAuth.POST("/ticket/create", handler.CreateTicket)
	scheduleAuth.POST("/ticket/booking", handler.Booking)
	scheduleAuth.GET("/ticket/info/:ticket_id", handler.GetInfoByTicketId)

	scheduleAuth.POST("/booking/history", handler.BookingHistory)
	scheduleAuth.POST("/booking/:department_id", handler.BookingHistoryByDept)
	scheduleAuth.GET("/booking/history/:booking_id", handler.GetBookingHistoryById)
	scheduleAuth.GET("/booking/statistics/:department_id", handler.GetBookingStatisticsByDept)
	scheduleAuth.POST("/list/all", handler.GetScheduleListByDept)
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

	messageAuth := authRouter.Group("/letter", middleware.JWTAuth(services.AppGuardName))
	messageAuth.POST("/list", handler.GetLetterList)
	messageAuth.GET("/unread_count", handler.UnreadCount)
	messageAuth.POST("/detail/:conversationId", handler.GetConversationDetail)
	messageAuth.POST("/send_letter", handler.SendLetter)
	messageAuth.POST("/notice/list", handler.GetNoticeList)
	messageAuth.POST("/notice/send_notice", handler.SendNotice)
}
