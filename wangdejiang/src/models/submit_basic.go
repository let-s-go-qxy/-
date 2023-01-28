package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string       `gorm:"type:varchar(36)" json:"identity"`
	ProblemIdentity string       `gorm:"type:varchar(36)" json:"problem_identity"`
	ProblemBasic    ProblemBasic `gorm:"foreignKey:identity;references:problem_identity"` // problem关联
	UserIdentity    string       `gorm:"type:varchar(36)" json:"user_identity"`
	UserBasic       UserBasic    `gorm:"foreignKey:identity;references:user_identity"` // user关联
	Path            string       `gorm:"type:varchar(255)" json:"path"`
	Status          int          `gorm:"type:tinyint(1)" json:"status"` // 代码状态 -1 待判断 1 答案正确 2 答案错误 3 运行超时 4 运行超内存
}

func (receiver SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(pid, uid string, status int) (tx *gorm.DB) {
	// 自定义Preload
	tx = Db.Model(&SubmitBasic{}).Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
		return Db.Omit("content")
	}).Preload("UserBasic", func(db *gorm.DB) *gorm.DB {
		return Db.Omit("password")
	})
	if pid != "" {
		tx.Where("problem_identity = ?", pid)
	}
	if uid != "" {
		tx.Where("user_identity = ?", uid)
	}
	if status != 0 {
		tx.Where("status = ?", status)
	}
	return
}
