package utils

import (
	"github.com/gofrs/uuid/v5"
	"time"
)

// GenerateSimpleRandomNumber 生成一个6位的伪随机数字
func GenerateSimpleRandomNumber() int {
	// 获取当前时间的纳秒部分
	nano := time.Now().UnixNano()

	// 对纳秒取模以确保结果在6位数范围内
	randomNumber := int(nano%900000) + 100000

	return randomNumber
}

// GenerateUuid 生成一个uuid
func GenerateUuid() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}
