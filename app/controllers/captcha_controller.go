package controllers

import (
	"encoding/json"
	"net/http"
	"plp/app/services"
)

func init() {
}

type CaptchaController struct {
}

// 渲染验证码
func (captcha_controller CaptchaController) Captcha(w http.ResponseWriter, r *http.Request) {
	data := services.GenerateCaptcha()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	body := map[string]interface{}{
		"code":      1,
		"data":      data["img"],
		"captchaId": data["uuid"],
		"msg":       "success",
	}
	json.NewEncoder(w).Encode(body)
}
