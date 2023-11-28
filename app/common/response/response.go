package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"registration-booking/global"
)

// 响应结构体
type Response struct {
	Code int         `json:"code"` // 自定义错误码
	Data interface{} `json:"data"` // 数据
	Mes  string      `json:"mes"`  // 信息
}

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		data,
		"ok",
	})
}

// AuthorizationFail 认证失败 ErrorCode 不为 0 表示失败
func AuthorizationFail(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, Response{
		1,
		nil,
		msg,
	})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		1,
		nil,
		msg,
	})
}

func ServerError(c *gin.Context, err interface{}) {
	msg := "Internal Server Error"
	// 非生产环境显示具体错误信息
	if global.App.Config.App.Env != "production" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
		if _, ok := err.(error); ok {
			msg = err.(error).Error()
		}
	}
	c.JSON(http.StatusInternalServerError, Response{
		1,
		nil,
		msg,
	})
	c.Abort()
}
