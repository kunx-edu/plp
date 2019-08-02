package controllers

import (
	"fmt"
	"github.com/thinkeridea/go-extend/exnet"
	"net/http"
	"plp/app/services"
	"plp/lib"
	"regexp"
	"time"
)

func init() {

}

type UserController struct {
	*Controller
}

// 注册
func (u UserController) Register(w http.ResponseWriter, r *http.Request) {
	// 检验数据是否完整
	login_name := r.FormValue("login_name")
	password := r.FormValue("password")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	gender := r.FormValue("gender")

	error_message := ""
	if len(login_name) < 1 {
		error_message = "用户名不能为空"
	} else if valid_login_name_flag, _ := regexp.MatchString(`^[a-z\d_]{5,16}$`, login_name); !valid_login_name_flag {
		error_message = "用户名由5-16位小写字母、数字和下划线组成"
	} else if len(password) < 1 {
		error_message = "用户名由5-16位小写字母、数字和下划线组成"
	} else if valid_password_flag, _ := regexp.MatchString(`^.{8,16}$`, password); !valid_password_flag {
		error_message = "用户名由5-16位小写字母、数字和下划线组成"
	} else if len(email) < 1 {
		error_message = "用户名由5-16位小写字母、数字和下划线组成"
	} else if valid_email_flag, _ := regexp.MatchString(`^[a-zA-Z0-9_\-\.]+@[a-zA-Z0-9_\-]+(\.[a-zA-Z0-9_\-]+)+$`, email); !valid_email_flag {
		error_message = "用户名由5-16位小写字母、数字和下划线组成"
	} else if len(phone) < 1 {
		error_message = "用户名由5-16位小写字母、数字和下划线组成"
	} else if valid_phone_flag, _ := regexp.MatchString(`^\d{11}$`, phone); !valid_phone_flag {
		error_message = "用户名由5-16位小写字母、数字和下划线组成"
	} else if len(gender) < 1 {
		error_message = "用户名由5-16位小写字母、数字和下划线组成"
	} else if gender != "男" && gender != "女" && gender != "保密" {
		error_message = "用户名由5-16位小写字母、数字和下划线组成"
	}

	if len(error_message) > 0 {
		fmt.Println(error_message)
		data := map[string]interface{}{
			"code": 0,
			"msg":  error_message,
		}
		u.json(w, r, data)
		return
	}

	// 校验验证码
	//uuid := r.FormValue("uuid")
	//code := r.FormValue("code")
	//if !services.ValidCaptcha(uuid, code) {
	//	w.WriteHeader(http.StatusForbidden)
	//	fmt.Println("禁止访问")
	//	//http.Error(w, "禁止访问", http.StatusForbidden)
	//} else {
	// 查看用户名是否已经存在
	select_sql := "SELECT count(1) FROM users WHERE login_name=?"
	db := services.DB_CONN
	var count int
	err := db.QueryRow(select_sql, login_name).Scan(&count)
	if err != nil {
		error_message = err.Error()
	} else if count > 0 {
		error_message = "用户已存在"
	}

	if len(error_message) > 0 {
		fmt.Println(error_message)
		data := map[string]interface{}{
			"code": 0,
			"msg":  error_message,
		}
		u.json(w, r, data)
		return
	}

	// 密码加密，保存到数据库
	encrypt_password := lib.EncryptPassword(password)
	insert_sql := "INSERT INTO users (login_name,password,gender,phone,email,register_time,register_ip) VALUES (?,?,?,?,?,?,?)"
	register_time := time.Now().Format("2006-01-02 15:04:05")
	register_ip := exnet.RemoteIP(r)
	stmt, _ := db.Prepare(insert_sql)

	_, err = stmt.Exec(login_name, encrypt_password, gender, phone, email, register_time, register_ip) // 中文出现了乱码

	if err != nil {
		error_message = err.Error()
	}

	if len(error_message) > 0 {
		fmt.Println(error_message)
		data := map[string]interface{}{
			"code": 0,
			"msg":  error_message,
		}
		u.json(w, r, data)
		return
	}

	success_message := "注册成功,用户名为：" + login_name
	fmt.Println(success_message)
	data := map[string]interface{}{
		"code": 0,
		"msg":  success_message,
	}
	u.json(w, r, data)
	return

}

// 登录
func (u UserController) Login(w http.ResponseWriter, r *http.Request) {

}

// 退出
func (u UserController) Logout(w http.ResponseWriter, r *http.Request) {

}

// 上传头像
func (u UserController) UploadAvatar(w http.ResponseWriter, r *http.Request) {

}

// 修改资料
func (u UserController) Mod(w http.ResponseWriter, r *http.Request) {

}
