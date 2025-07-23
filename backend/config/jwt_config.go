package config

import "time"

const (
	JWTSecretKey      = "your_secret_key" // 替换为实际的密钥
	JWTExpirationTime = 24 * time.Hour    // Token 过期时间
)
