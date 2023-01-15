package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity string `gorm:"type:varchar(36)" json:"identity"` // 标识
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Password string `gorm:"type:varchar(32)" json:"password"`
	Phone    string `gorm:"type:varchar(100)" json:"phone"`
	Mail     string `gorm:"type:varchar(100)" json:"mail"`
}

func (receiver *User) TableName() string {
	return "user"
}
