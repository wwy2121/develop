package helper

import (
	"crypto/md5"
	"fmt"
)
//加密
func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
