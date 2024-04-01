package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/models"
	"registration-booking/app/services"
	"strconv"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	mobile, err := strconv.Atoi(form.Mobile)
	if err != nil {
		return
	}
	form.Mobile = strconv.Itoa(mobile % 100000)
	if err, user := services.UserService.Login(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.Fail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

func RefreshToken(c *gin.Context) {
	if err, user := services.UserService.GetUserById(c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.Fail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, user)
}

func UserInfoAndRole(c *gin.Context) {
	err, user := services.UserService.UserInfoAndRole(c.Keys["id"].(string))
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, user)
}

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.Fail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}

func WxLogin(c *gin.Context) {
	var form request.WxLogin
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.WxLogin(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.Fail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// Register 用户注册
func CreateDoctor(c *gin.Context) {
	var form request.CreateDoctorReq
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.CreateDoctor(form, c.Keys["id"].(string)); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

func GetDoctorList(c *gin.Context) {
	if err, list := services.UserService.GetDoctorList(c.Param("department_id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, list)
	}
}

func GetAllDoctorList(c *gin.Context) {
	var form request.GetDoctorListReq
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, total, res := services.UserService.GetAllDoctorList(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, response.PageData{
			PageData: res,
			Total:    total,
		})
	}
}

func GetDoctorById(c *gin.Context) {
	if err, list := services.UserService.GetDoctorById(c.Param("id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, list)
	}
}

func GetDiseaseList(c *gin.Context) {
	var form request.Page
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, total, list := services.UserService.GetDiseaseList(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, response.PageData{
			PageData: list,
			Total:    total,
		})
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(c, request.GetErrorMsg(user, err))
		return
	}
	if err := services.UserService.UpdateUser(user); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}

func UpdateDoctor(c *gin.Context) {
	var user request.DoctorReq
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(c, request.GetErrorMsg(user, err))
		return
	}
	if err := services.UserService.UpdateDoctor(user); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}

func DeleteUser(c *gin.Context) {
	if err := services.UserService.DeleteUser(c.Param("id")); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}
