package main

import (
	"fmt"
	"net/http"
	"plp/app/controllers"
	"plp/app/services"
	"plp/lib"
)

type myHandler struct {
}

// 路由列表
var routers map[string]func(w http.ResponseWriter, r *http.Request)

func register_request_handler() {
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
		fmt.Println("接到请求 " + path)
		hand_func(w, r)
	} else {
		http.NotFound(w, r)
	}
	return;
}

func main() {
	// 注册配置文件
	fmt.Println("读取配置文件")
	lib.InitConfig("./app/conf/config.ini")

	// 连接数据库
	fmt.Println("连接数据库")
	services.InitDBConn()

	// 注册路由
	fmt.Println("注册路由")
	register_request_handler();
	port := lib.Read("base", "port")

	// 开启服务
	fmt.Println("开启TCP服务 监听端口 " + port)
	http.ListenAndServe(":"+port, myHandler{})

}
