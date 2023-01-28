package controller

import (
	"awesomeProject/wangdejiang/src/models"
	"awesomeProject/wangdejiang/src/service"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Summary 问题列表
// @Tags 公共方法
// @Param page query int false "请输入当前页，默认1"
// @Param size query int false "请输入size，默认10"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "分类唯一标识"
// @Success 200 {string} json"{"code": 200, "msg":"ok", "data":[]}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	cid := c.Query("category_identity")
	keyword := c.Query("keyword")
	var total int64
	tx := models.GetProblemList(keyword, cid)
	var problems []models.ProblemBasic
	err := tx.Count(&total).Offset((page - 1) * size).Limit(size).Find(&problems).Error
	if err != nil {
		log.Panicln("GetProblemList error", err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "not ok",
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": map[string]interface{}{
			"data":  problems,
			"count": total,
		},
	})
}

// GetProblemDetail
// @Summary 问题详情
// @Tags 公共方法
// @Param identity query string true "唯一标识"
// @Success 200 {string} json"{"code": 200, "msg":"ok", "data":{}}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "唯一标识不能为空",
		})
		return
	}
	p := new(models.ProblemBasic)
	err := models.Db.Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		First(&p, "identity = ?", identity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "记录不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "其他错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  "ok",
		"data": p,
	})
}

// ProblemCreate
// @Summary 创建问题
// @Tags 管理员私有
// @Param token header string true "token"
// @Param title formData string true "title"
// @Param content formData string true "content"
// @Param max_runtime formData string true "max_runtime"
// @Param max_mem formData string true "max_memory"
// @Param category_ids formData array false "问题分类, 注意传递id列表 1"
// @Param test_cases formData array false "测试样例 {"input":"1 2\n", "output":"3\n"}
// @Success 200 {string} json"{"code": 200, "msg":"ok", "data":{}}"
// @Router /problem-create [post]
func ProblemCreate(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	maxRuntime, _ := strconv.Atoi(c.PostForm("max_runtime"))
	maxMem, _ := strconv.Atoi(c.PostForm("max_mem"))
	categoryIds := c.PostFormArray("category_ids")
	testCases := c.PostFormArray("test_cases")
	// TODO 判断为空

	data := models.ProblemBasic{
		Identity:   service.CreateUuid(),
		Title:      title,
		Content:    content,
		MaxMem:     maxMem,
		MaxRuntime: maxRuntime,
		PassNumb:   0,
		TotalNumb:  0,
	}

	// 处理分类
	problemCategories := make([]models.ProblemCategory, 0)
	for _, id := range categoryIds {
		unitId, _ := strconv.Atoi(id)
		problemCategories = append(problemCategories, models.ProblemCategory{
			ProblemId:  data.ID,
			CategoryId: uint(unitId),
		})
	}
	// 处理测试样例
	testCacies := make([]models.TestCase, 0)
	for _, pCase := range testCases {
		caseMap := make(map[string]string)
		json.Unmarshal([]byte(pCase), &caseMap)
		// TODO 加判断是map否有key
		testCacies = append(testCacies, models.TestCase{
			Identity:        service.CreateUuid(),
			ProblemIdentity: data.Identity,
			Input:           caseMap["input"],
			Output:          caseMap["output"],
		})
	}
	data.ProblemCategories = problemCategories
	data.TestCase = testCacies
	err := models.Db.Create(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建失败:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": data,
	})
}
