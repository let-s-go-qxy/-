package controller

import (
	"awesomeProject/wangdejiang/src/models"
	"awesomeProject/wangdejiang/src/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// GetUserDetail
// @Summary 用户详情
// @Tags 公共方法
// @Param identity query string false "唯一标识"
// @Success 200 {string} json"{"code": 200, "msg":"ok", "data":{}}"
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "唯一标识不能为空",
		})
		return
	}
	user := new(models.UserBasic)
	err := models.Db.Omit("password").
		First(&user, "identity = ?", identity).Error
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
			"msg":  "其他错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  "ok",
		"data": user,
	})
}

// Login
// @Summary 用户登录,得到token
// @Tags 公共方法
// @Param username formData string false "user"
// @Param password formData string false "password"
// @Success 200 {string} json"{"code": 200, "msg":"ok", "data":{}}"
// @Router /login [post]
func Login(c *gin.Context) {
	name := c.PostForm("username")
	pw := service.GetMd5(c.PostForm("password")) // md5
	if name == "" || pw == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
		return
	}
	data := new(models.UserBasic)
	err := models.Db.First(data, "name = ? AND password = ?", name, pw).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码出错",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "其他错误: " + err.Error(),
		})
		return
	}
	// 生成token
	claim := service.UserClaims{
		Identity: data.Identity,
		Name:     data.Name,
		IsAdmin:  data.IsAdmin,
	}
	token, _ := claim.GenerateToken()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": map[string]string{"token": token},
	})
}

// SendCode
// @Summary 发送验证码
// @Tags 公共方法
// @Param email formData string false "发送账号"
// @Success 200 {string} json"{"code": 200, "msg":"ok"}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "无邮箱",
		})
	}
	code := service.GenarateCode()
	service.Rdb.Set(c, email, code, time.Second*60*5)
	s := service.SendCode(email, code)
	if s != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "err: " + s.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  "邮箱发送成功",
	})
}

// Register
// @Summary 用户注册
// @Tags 公共方法
// @Param mail formData string false "邮箱"
// @Param code formData string false "验证码"
// @Param name formData string false "名称"
// @Param password formData string false "密码"
// @Param phone formData string false "手机号"
// @Success 200 {string} json"{"code": 200, "msg":"ok"}"
// @Router /register [post]
func Register(c *gin.Context) {
	mail := c.PostForm("mail")
	code := c.PostForm("code")
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	// TODO 这里应该分别判断和返回
	if mail == "" || code == "" || name == "" || password == "" || phone == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数缺失",
		})
		return
	}
	// 验证码验证
	code2, err := service.Rdb.Get(c, mail).Result()
	if code2 != code || err != nil {
		println("error: ", err)
		println("code: ", code2, "exceptCode: ", code)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱验证失败",
		})
		return
	}
	// 判断邮箱是否已存在
	var userCount int64
	err = models.Db.Model(models.UserBasic{}).Where("mail = ?", mail).Count(&userCount).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Count User Error" + err.Error(),
		})
		return
	}
	if userCount > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该邮箱已经有用户绑定",
		})
		return
	}
	// 数据插入
	identity := service.CreateUuid()
	var data = &models.UserBasic{
		Identity: identity,
		Name:     name,
		Password: service.GetMd5(password), // 密码md5
		Mail:     mail,
		Phone:    phone,
	}
	err = models.Db.Create(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create User Error" + err.Error(),
		})
		return
	}
	// 生成token
	claims := service.UserClaims{
		Identity: identity,
		Name:     name,
		IsAdmin:  data.IsAdmin,
	}
	token, _ := claims.GenerateToken()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// GetRankList
// @Summary 用户榜单
// @Tags 公共方法
// @Param page query int false "请输入当前页，默认1"
// @Param size query int false "请输入size，默认10"
// @Success 200 {string} json"{"code": 200, "msg":"ok"}"
// @Router /rank-list [get]
func GetRankList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	var count int64
	var list []models.UserBasic
	err := models.Db.Model(models.UserBasic{}).Count(&count).Order("finish_problem_sum desc, submit_num asc").
		Offset((page - 1) * size).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
