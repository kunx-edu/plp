package controllers

import (
	"net/http"
	"plp/app/services"
)

func init() {
}

type CaptchaController struct {
	*Controller
}

// 渲染验证码
func (c CaptchaController) Captcha(w http.ResponseWriter, r *http.Request) {
	data := services.GenerateCaptcha()

	body := map[string]interface{}{
		"code":      1,
		"data":      data["img"],
		"captchaId": data["uuid"],
		"msg":       "success",
	}
	c.json(w, r, body)
}
