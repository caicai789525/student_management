package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户模型
// @Description 用户信息
type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"type:varchar(255);unique"`
	Password  string         `gorm:"type:varchar(255)"`
}
