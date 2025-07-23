package model

import (
	"gorm.io/gorm"
	"time"
)

// Student 学生模型
// @Description 学生信息
type Student struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	Age       int            `gorm:"type:int" json:"age"`
	Grade     string         `gorm:"type:varchar(255)" json:"grade"`
	Major     string         `gorm:"type:varchar(255)" json:"major"`
	// 学号，由年份加学院序号加学生序号组成的 11 位数字
	StudentID string `gorm:"type:char(11);unique" json:"student_id"`
	Class     string `gorm:"type:varchar(255)" json:"class"`
	Gender    string `gorm:"type:enum('男','女','未知')" json:"gender"`
}
