package test

import (
	"awesomeProject/wangdejiang/src/models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 测试前：数据装载，配置初始化等前置工作
	code := m.Run()
	// 测试后：释放资源等收尾工作
	os.Exit(code)
}

func TestEqual(t *testing.T) {
	output := 1
	expectOutput := 1
	assert.Equal(t, expectOutput, output)
}

// TestConnect 测试数据库连接
func TestConnect(t *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:imAei@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil || db.Error != nil {
		t.Fail()
		return
	}
	var data []models.ProblemBasic
	err = db.Find(&data).Error
	if err != nil {
		t.Fail()
	}
	for _, datum := range data {
		fmt.Printf("Promblem ===> %v \n", datum)
	}
}
