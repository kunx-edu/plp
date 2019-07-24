package controllers

import "net/http"

func init() {

}

type UserController struct {
}

// 注册
func (user_controller UserController) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("注册"))
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
