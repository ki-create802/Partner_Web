package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomCode 生成指定长度的随机验证码
func GenerateRandomCode(length int) string {
	var letters = []rune("0123456789")
	b := make([]rune, length)
	// 使用当前时间作为随机数生成器的种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
