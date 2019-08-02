package lib

import (
	"crypto/md5"
	"fmt"
	"io"
)

const (
	password_salt string = "plp_password"
)

// 密码加盐加密
func EncryptPassword(password string) string {
	w := md5.New()
	io.WriteString(w, password+password_salt)
	return fmt.Sprintf("%x", w.Sum(nil))
}
