package models

import "gorm.io/gorm"

type TestCase struct {
	gorm.Model
	Identity        string `json:"identity"`
	ProblemIdentity string `json:"problem_identity"`
	Input           string `json:"input"`
	Output          string `json:"output"`
}

func (t *TestCase) TableName() string {
	return "test_case"
}
