package controllers

import "net/http"

func init() {

}

type UserController struct {
}

// 注册
func (user_controller UserController) Register(r http.ResponseWriter, w *http.Request) {
	r.Write([]byte("注册"))
}

// 登录
func (user_controller UserController) Login(r http.ResponseWriter, w *http.Request) {

}

// 退出
func (user_controller UserController) Logout(r http.ResponseWriter, w *http.Request) {

}

// 上传头像
func (user_controller UserController) UploadAvatar(r http.ResponseWriter, w *http.Request) {

}

// 修改资料
func (user_controller UserController) Mod(r http.ResponseWriter, w *http.Request) {

}
