package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity string `gorm:"type:varchar(36)" json:"identity"` // 标识
	//CategoryId string `gorm:"column:cid;type:varchar(255)" json:"category_id"` // 分类id逗号分隔
	ProblemCategories []ProblemCategory `gorm:"foreignKey:problem_id;references:id"` // 关联问题分类表
	Title             string            `gorm:"type:varchar(255)" json:"title"`      // 文章title
	Content           string            `gorm:"type:text" json:"content"`            // 文章内容
	MaxMem            int               `gorm:"type:int(11)" json:"max_mem"`         // 最大内存
	MaxRuntime        int               `gorm:"type:int(11)" json:"max_runtime"`     // 最大运行时间
	PassNumb          int               `gorm:"type:int(11)" json:"pass_numb"`       // 问题通过数
	TotalNumb         int               `gorm:"type:int(11)" json:"total_numb"`      // 问题总数
	TestCase          []TestCase        `gorm:"foreignKey:problem_identity;references:identity"`
}

// TableName 获取表格名称
func (receiver *ProblemBasic) TableName() string {
	return "problem_basic"
}

// GetProblemList 获取Problem表所有记录
func GetProblemList(keyword, cid string) (tx *gorm.DB) {
	tx = Db.Model(new(ProblemBasic)).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")
	// 进行匹配关联查找操作
	if cid != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ?)", cid)
	}
	return
}
