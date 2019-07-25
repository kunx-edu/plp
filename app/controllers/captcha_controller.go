package controllers

import (
	"encoding/json"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

func init() {
}

type CaptchaController struct {
}

// 渲染验证码
func (captcha_controller CaptchaController) Captcha(w http.ResponseWriter, r *http.Request) {
	data := createCaptcha()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	body := map[string]interface{}{
		"code":      1,
		"data":      data["img"],
		"captchaId": data["uuid"],
		"msg":       "success",
	}
	json.NewEncoder(w).Encode(body)
}

// 校验验证码
func (captcha_controller CaptchaController) check(w http.ResponseWriter, r *http.Request) {

}

// 生成数字验证码
func createCaptcha() (data map[string]string) {
	// 数字验证码
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}

	// 第一个参数为空字符串，包会自动在服务器随机种子生成一个随机uiid
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)

	// 以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	data = map[string]string{
		"uuid": idKeyD,
		"img":  base64stringD,
	}
	return
}
