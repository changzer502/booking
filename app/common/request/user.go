package request

type Register struct {
	Nickname  string `form:"nickname" json:"nickname" binding:"required"`
	Gender    int    `form:"gender" json:"gender"`
	AvatarUrl string `form:"avatar_url" json:"avatar_url" binding:"required"`
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password  string `form:"password" json:"password" binding:"required"`
}

// 自定义错误信息
func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"nickname.required": "用户名称不能为空",
		//"gender.required":     "性别不能为空",
		"mobile.required":     "手机号码不能为空",
		"mobile.mobile":       "手机号码格式不正确",
		"avatar_url.required": "用户头像不能为空",
		"password.required":   "密码不能为空",
	}
}

type Login struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}

type WxLogin struct {
	PhoneCode string `form:"phone_code" json:"phone_code"`
	Code      string `form:"code" json:"code" binding:"required"`
}

func (wxLogin WxLogin) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"code.required": "code不能为空",
	}
}

type CreateDoctorReq struct {
	Nickname     string    `form:"nickname" json:"nickname" binding:"required"`
	AvatarUrl    string    `form:"avatar_url" json:"avatar_url" binding:"required"`
	DepartmentID uint      `form:"department_id" json:"department_id" binding:"required"`
	Introduce    Introduce `form:"introduce" json:"introduce" binding:"required"`
}

type Introduce struct {
	Duties                string `form:"duties" json:"duties" binding:"required"`
	Speciality            string `form:"speciality" json:"speciality" binding:"required"`
	EducationalBackground string `form:"educational_background" json:"educational_background" binding:"required"`
	WorkExperience        string `form:"work_experience" json:"work_experience" binding:"required"`
	ResearchDirection     string `form:"research_direction" json:"research_direction" binding:"required"`
	AcademicPositions     string `form:"academic_positions" json:"academic_positions" binding:"required"`
}
