package main

import (
	"net/http"
	"plp/app/controllers"
	"plp/lib"
)

type myHandler struct {
}

// 路由列表
var routers map[string]func(w http.ResponseWriter, r *http.Request)

func init() {
	var user_controller controllers.UserController
	var captcha_controller controllers.CaptchaController
	routers = map[string]func(r http.ResponseWriter, w *http.Request){
		"/": func(r http.ResponseWriter, w *http.Request) {
			r.Write([]byte("hello 漂流瓶"))
		},
		"/plp/register":      user_controller.Register,
		"/plp/login":         user_controller.Login,
		"/plp/logout":        user_controller.Logout,
		"/plp/upload-avatar": user_controller.UploadAvatar,
		"/plp/mod":           user_controller.Mod,
		"/plp/captcha":       captcha_controller.Captcha,
	}
}

func (handler myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if hand_func, ok := routers[path]; ok {
		hand_func(w, r)
	} else {
		http.NotFound(w, r)
	}
	return;
}

func main() {
	// 注册配置文件
	config := new(lib.Config)
	config.InitConfig("./app/conf/config.ini")
	port := config.Read("base", "port")

	// 开启服务
	http.ListenAndServe(":"+port, myHandler{})

}
