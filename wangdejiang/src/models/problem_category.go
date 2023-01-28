package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId     uint          `json:"problem_id"`
	CategoryId    uint          `json:"category_id"`
	CategoryBasic CategoryBasic `gorm:"foreignKey:id;references:category_id"` //关联分类的基础信息表
}

func (c ProblemCategory) TableName() string {
	return "problem_category"
}
