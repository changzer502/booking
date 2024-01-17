package services

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/goccy/go-json"
	"net/http"
	"registration-booking/app/common/request"
	"registration-booking/app/common/response"
	"registration-booking/app/models"
	"registration-booking/global"
	"registration-booking/utils"
	"strconv"
	"strings"
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
func (userService *userService) GetUserById(id string) (err error, user *models.User) {
	uid, _ := strconv.Atoi(id)
	user, err = models.FindUserById(uint(uid))
	if err != nil {
		err = errors.New("用户不存在")
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

func (userService *userService) UserInfoAndRole(id string) (err error, user response.UserRes) {
	intId, err := strconv.Atoi(id)
	u := models.User{}
	err = global.App.DB.First(&u, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	role, err := models.GetRoleById(u.RoleId)
	if err != nil {
		return err, response.UserRes{}
	}
	user.ID = u.ID.ID
	user.Nickname = u.Nickname
	user.Role = append(user.Role, role.RoleKey)
	user.AvatarUrl = u.AvatarUrl
	user.DepartmentId = u.DepartmentID
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

func (userService *userService) CreateDoctor(params request.CreateDoctorReq, id string) (err error, user models.User) {
	introduce, _ := json.Marshal(params.Introduce)
	uid, _ := strconv.Atoi(id)
	user = models.User{
		Nickname:     params.Nickname,
		AvatarUrl:    params.AvatarUrl,
		DepartmentID: params.DepartmentID,
		Password:     utils.BcryptMake([]byte("123456")),
		Introduce:    string(introduce),
		RoleId:       2,
		Timestamps: models.Timestamps{
			CreatedAt: time.Now(),
			CreatedBy: uint(uid),
			UpdatedAt: time.Now(),
			UpdatedBy: uint(uid),
		},
	}
	err = global.App.DB.Create(&user).Error
	return
}

func (userService *userService) GetDoctorListBySpider(url string, id uint) (err error, users []models.User) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("解析HTML失败：", err)
	}
	doc.Find(".subPage").Find("li").Each(func(i int, selection *goquery.Selection) {
		link, err := selection.Find("a").Attr("href")
		if err != true {
			fmt.Println("获取链接失败：", err)
		} else {
			user := download(urlJoin(link, url), id)
			err, _ := userService.CreateDoctor(user, "2")
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	})
	return
}
func download(url string, departmentID uint) request.CreateDoctorReq {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败：", err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("解析HTML失败：", err)
	}
	// 头像

	avatarUrl := ""
	doc.Find(".imgResponsive").Find("img").Each(func(i int, selection *goquery.Selection) {
		avatarUrl, _ = selection.Attr("src")
	}).Attr("src")
	avatarUrl = "https://www.btch.edu.cn/" + avatarUrl[12:]

	info := doc.Find(".doctorList01").Text()
	info = strings.ReplaceAll(info, "姓 名：", "")
	info = strings.ReplaceAll(info, "职 务：", "")
	info = strings.ReplaceAll(info, "特 长：", "\n")
	infos := strings.Split(info, "\n")
	user := request.CreateDoctorReq{
		AvatarUrl:    avatarUrl,
		DepartmentID: departmentID,
	}
	if len(infos) > 1 {
		user.Nickname = strings.Trim(infos[1], " ")
	}
	introduce := request.Introduce{}
	if len(infos) > 2 {
		introduce.Duties = strings.Trim(infos[2], " ")
	}
	if len(infos) > 3 {
		introduce.Speciality = strings.Trim(infos[3], " ")
	}
	other := doc.Find(".doctorColumn").Text()
	others := strings.Split(other, "\n")
	if len(others) > 2 {
		introduce.EducationalBackground = strings.Trim(others[2], " ")
	}

	if len(others) > 5 {
		introduce.WorkExperience = strings.Trim(others[5], " ")
	}
	if len(others) > 8 {
		introduce.ResearchDirection = strings.Trim(others[8], " ")
	}
	if len(others) > 11 {
		introduce.AcademicPositions = strings.Trim(others[11], " ")
	}
	user.Introduce = introduce

	fmt.Println(user)
	return user
}

func urlJoin(href, base string) string {
	base = base[:len(base)-9]
	return base + href
}

func (userService *userService) GetDoctorList(departmentId string) (err error, user []models.User) {
	err = global.App.DB.Where("department_id = ? and role_id = ?", departmentId, 2).Find(&user).Error
	return
}

func (userService *userService) GetAllDoctorList(page request.GetDoctorListReq) (err error, total int64, doctors []response.DoctorRes) {
	query := ""
	if page.Query != "" {
		query = " and nickname like '%" + page.Query + "%'"
	}
	deptIds := make([]int, 0)
	if page.Dept > 0 {
		depts, err := models.FindChildrenByDeptId(page.Dept)
		if err != nil {
			return err, 0, nil
		}
		deptIds = append(deptIds, int(page.Dept))
		for _, dept := range depts {
			deptIds = append(deptIds, int(dept.ID.ID))
		}
	}

	doc, total, err := models.FindAllDoctorByPage(page.PageNo, page.PageSize, query, deptIds)
	if err != nil {
		return err, 0, nil
	}
	for _, doctor := range doc {
		dept, err := models.FindDepartmentById(doctor.DepartmentID)
		if err != nil {
			return err, 0, nil
		}
		info := request.Introduce{}
		err = json.Unmarshal([]byte(doctor.Introduce), &info)
		if err != nil {
			return err, 0, nil
		}
		doctors = append(doctors, response.DoctorRes{
			User:           doctor,
			DepartmentName: dept.DeptName,
			Introduce:      info,
		})
	}
	return
}

func (userService *userService) GetDoctorById(id string) (err error, user models.User) {
	err = global.App.DB.Where("id = ? and role_id = ?", id, 2).Find(&user).Error
	if user.ID.ID == 0 {
		err = errors.New("数据不存在")
	}
	return
}
func (userService *userService) GetDiseaseList(page request.Page) (err error, total int64, users []models.User) {
	query := ""
	if page.Query != "" {
		query = " and nickname like '%" + page.Query + "%'"
	}
	users, total, err = models.FindDiseaseByPage(page.PageNo, page.PageSize, query)
	if err != nil {
		return err, 0, nil
	}
	return
}

func (userService *userService) UpdateUser(user models.User) (err error) {
	err = global.App.DB.Updates(&user).Error
	return
}

func (userService *userService) UpdateDoctor(doctor request.DoctorReq) (err error) {
	user := doctor.User
	marshal, err := json.Marshal(doctor.Introduce)
	if err != nil {
		return err
	}
	user.Introduce = string(marshal)
	err = models.UpdateUser(user)
	return
}

func (userService *userService) DeleteUser(id string) (err error) {
	err = global.App.DB.Delete(&models.User{}, id).Error
	return
}
