package controllers

import (
	"net/http"
)

func init() {
}

type CaptchaController struct {
}

// 渲染验证码
func (captcha_controller CaptchaController) Captcha(w http.ResponseWriter, r *http.Request) {
}

// 校验验证码
func (captcha_controller CaptchaController) check(w http.ResponseWriter, r *http.Request) {

}
