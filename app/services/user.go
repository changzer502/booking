package services

import (
	"errors"
	"github.com/goccy/go-json"
	"registration-booking/app/common/request"
	"registration-booking/app/models"
	"registration-booking/global"
	"registration-booking/utils"
	"strconv"
	"time"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Nickname: params.Nickname, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("mobile = ?", params.Mobile).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

func (userService *userService) WxLogin(params request.WxLogin) (err error, user *models.User) {
	// 获取openId
	openIdRes, err := global.App.Config.Wx.GetOpenId(params.Code)
	var result = global.App.DB.Where("open_id = ?", openIdRes.Openid).First(&user)
	if result.RowsAffected == 0 {
		// 进行注册
		user = &models.User{OpenId: openIdRes.Openid, Nickname: "微信用户" + openIdRes.Openid[len(openIdRes.Openid)-5:len(openIdRes.Openid)],
			AvatarUrl: "https://thirdwx.qlogo.cn/mmopen/vi_32/POgEwh4mIHO4nibH0KlMECNjjGxQUq24ZEaGT4poC6icRiccVGKSyXwibcPq4BWmiaIGuG1icwxaQX6grC9VemZoJ8rg/132"}
		err = global.App.DB.Create(&user).Error
		if err != nil {
			return err, nil
		}
	}

	return
}

func (userService *userService) CreateDoctor(params request.CreateDoctorReq, id uint) (err error, user models.User) {
	introduce, _ := json.Marshal(params.Introduce)
	user = models.User{
		Nickname:     params.Nickname,
		AvatarUrl:    params.AvatarUrl,
		DepartmentID: params.DepartmentID,
		Password:     utils.BcryptMake([]byte("123456")),
		Introduce:    string(introduce),
		Timestamps: models.Timestamps{
			CreatedAt: time.Now(),
			CreatedBy: id,
			UpdatedAt: time.Now(),
			UpdatedBy: id,
		},
	}
	err = global.App.DB.Create(&user).Error
	return
}
