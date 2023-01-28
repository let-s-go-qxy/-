package middleware

import (
	"awesomeProject/wangdejiang/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthAdminCheck 鉴权中间件
func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		var userClaim = new(service.UserClaims)
		err := userClaim.ParseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "error: " + err.Error(),
			})
			return
		}
		if userClaim.IsAdmin != 1 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "非管理员",
			})
			return
		}
		c.Next()
	}
}
