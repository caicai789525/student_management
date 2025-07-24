package config

import "time"

const (
	JWTSecretKey = "tech-123456@student_management"
	//jwt时限（应该不会有人用这么久吧）
	JWTExpirationTime = 24 * time.Hour
)
