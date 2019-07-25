package services

import "github.com/mojocn/base64Captcha"

// 生成数字验证码
func GenerateCaptcha() (data map[string]string) {
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

// 校验验证码
func ValidCaptcha(uuid string, code string) bool {
	return base64Captcha.VerifyCaptcha(uuid, code)
}
