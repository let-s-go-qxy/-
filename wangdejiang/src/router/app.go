package router

import (
	"awesomeProject/wangdejiang/src/controller"
	_ "awesomeProject/wangdejiang/src/docs"
	"awesomeProject/wangdejiang/src/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	// swagger配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 定义路由规则
	// 问题路由
	r.GET("/problem-list", controller.GetProblemList)
	r.GET("/problem-detail", controller.GetProblemDetail)
	// 用户路由
	r.GET("/user-detail", controller.GetUserDetail)
	r.POST("/login", controller.Login)
	r.POST("/send-code", controller.SendCode)
	r.POST("/register", controller.Register)
	// 提交记录路由
	r.GET("/submit-list", controller.GetSubmitList)
	// 排行榜
	r.GET("/rank-list", controller.GetRankList)

	// 私有方法
	// 问题创建
	r.POST("/problem-create", middleware.AuthAdminCheck(), controller.ProblemCreate)
	return r
}
