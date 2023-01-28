package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity         string `gorm:"type:varchar(36)" json:"identity"` // 标识
	Name             string `gorm:"type:varchar(100)" json:"name"`
	Password         string `gorm:"type:varchar(32)" json:"password,omitempty"`
	Phone            string `gorm:"type:varchar(100)" json:"phone"`
	Mail             string `gorm:"type:varchar(100)" json:"mail"`
	FinishProblemSum int64  `json:"finish_problem_sum"`
	SubmitNum        int64  `json:"submit_num"`
	IsAdmin          int    `json:"is_admin"`
}

func (receiver *UserBasic) TableName() string {
	return "user_basic"
}
