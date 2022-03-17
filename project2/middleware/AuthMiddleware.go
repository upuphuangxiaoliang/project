package middleware

import (
	"net/http"
	"project/project2/common"
	datebase "project/project2/database"
	"strings"

	"github.com/gin-gonic/gin"
)

// 这一部分是中间件   这个中间件是为了验证token的正确性
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取authmiddle header
		tokenString := c.GetHeader("Authorization")

		// 验证格式 如果这个字符欻为空 或者 不是一bearer开头  表明为错误的token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 如果格式对    因为头部占了7个字符
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			// 如果不正确
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 这里是中间件用法 表示不在向下运行了
			c.Abort()
			return
		}

		// 验证同uo 获取的 claim中的userid
		userID := claims.UserId
		user := datebase.ID2User(int(userID))
		if datebase.ID(user) == -1 {
			// -1代表没有找到 代表不存在
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 如果存在 j减user信息写入上下文  给后面试用
		c.Set("user", user)

		// 跳到下一个fun
		c.Next()
	}
}
