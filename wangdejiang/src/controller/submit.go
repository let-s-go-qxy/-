package controller

import (
	"awesomeProject/wangdejiang/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetSubmitList
// @Summary 提交列表
// @Tags 公共方法
// @Param page query int false "请输入当前页，默认1"
// @Param size query int false "请输入size，默认10"
// @Param problem_identity query string false "问题唯一标识"
// @Param user_identity query string false "用户唯一标识"
// @Param status query string false "问题状态"
// @Success 200 {string} json"{"code": 200, "msg":"ok", "data":[]}"
// @Router /submit-list [get]
func GetSubmitList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	var count int64
	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	status, _ := strconv.Atoi(c.DefaultQuery("status", "0"))
	tx := models.GetSubmitList(problemIdentity, userIdentity, status)

	var data []models.SubmitBasic
	err := tx.Count(&count).Offset((page - 1) * size).Limit(size).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "error: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": map[string]interface{}{
				"list":  data,
				"count": count,
			},
		})
	}
}
