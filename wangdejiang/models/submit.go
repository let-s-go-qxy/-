package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `gorm:"type:varchar(36)" json:"identity"`
	ProblemIdentity string `gorm:"type:varchar(36)" json:"problem_identity"`
	UserIdentity    string `gorm:"type:varchar(36)" json:"user_identity"`
	Path            string `gorm:"type:varchar(255)" json:"path"`
	Status          int    `gorm:"type:tinyint(1)" json:"status"`
}

func (receiver Submit) TableName() string {
	return "submit"
}
