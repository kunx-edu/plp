package controllers

import (
	"fmt"
	"net/http"
	"plp/app/services"
	"regexp"
)

func init() {

}

type UserController struct {
}

// 注册
func (user_controller UserController) Register(w http.ResponseWriter, r *http.Request) {
	// 检验数据是否完整
	login_name := r.FormValue("login_name")
	password := r.FormValue("password")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	gender := r.FormValue("gender")

	login_name = "a"
	if len(login_name) < 1 {
		fmt.Println("用户名不能为空")
	} else if valid_login_name_flag, _ := regexp.MatchString(`^[a-z\d_]{5-16}$`, login_name); !valid_login_name_flag {
		fmt.Println("用户名由5-16位小写字母、数字和下划线组成")
	}

	//使用正则零宽断言（正则环视）
	//最短8位，最长16位 {8,16}
	//必须包含1个数字
	//必须包含2个小写字母
	//必须包含2个大写字母
	//必须包含1个特殊字符
	if len(password) < 1 {
		fmt.Println("密码不能为空")
	} else if valid_password_flag, _ := regexp.MatchString(`^.*(?=.{8,16})(?=.*\d)(?=.*[A-Z]{2,})(?=.*[a-z]{2,})(?=.*[!@#$%^&*?\(\)]).*$`, password); !valid_password_flag {
		fmt.Println("密码必须为8-16位，必须包含1个数字，必须包含2个小写字母，必须包含2个大写字母，必须包含1个特殊字符")
	}

	if len(email) < 1 {
		fmt.Println("邮箱不能为空")
	} else if valid_email_flag, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email); !valid_email_flag {
		fmt.Println("邮箱不合法")
	}

	// 正则验证
	if len(phone) < 1 {
		fmt.Println("手机号码不能为空")
	} else if valid_phone_flag, _ := regexp.MatchString(`^\d{11}$`, phone); !valid_phone_flag {
		fmt.Println("手机号码不合法")
	}

	if len(gender) < 1 {
		fmt.Println("性别不能为空")
	} else if gender != "男" && gender != "女" && gender != "保密" {
		fmt.Println("性别不合法")
	}

	// 校验验证码
	uuid := r.FormValue("uuid")
	code := r.FormValue("code")
	if !services.ValidCaptcha(uuid, code) {
		w.WriteHeader(http.StatusForbidden)
		//http.Error(w, "禁止访问", http.StatusForbidden)
	} else {
		w.Write([]byte("注册"))
	}

	// todo 密码加密，保存到数据库

}

// 登录
func (user_controller UserController) Login(w http.ResponseWriter, r *http.Request) {

}

// 退出
func (user_controller UserController) Logout(w http.ResponseWriter, r *http.Request) {

}

// 上传头像
func (user_controller UserController) UploadAvatar(w http.ResponseWriter, r *http.Request) {

}

// 修改资料
func (user_controller UserController) Mod(w http.ResponseWriter, r *http.Request) {

}
