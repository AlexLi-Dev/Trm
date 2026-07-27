// Author: Lutong.li
package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登陆
func Login(c *gin.Context) {
	/*
		测试 API
		curl -X POST http://localhost:8080/api/login \
		  -H "Content-Type: application/json" \
		  -d '{"username":"a","password":"admin123"}'

	*/

	type User struct {
		Username string `json:"username"` // ✅ 需要匹配请求中的 "username"
		Password string `json:"password"` // ✅ 需要匹配请求中的 "password"
	}
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误: " + err.Error(),
		})

		return
	}
	if user.Username == "admin" && user.Password == "admin123" {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "login success",
		})
		return
	}
	// 3. 登录失败
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    401,
		"message": "用户名或密码错误",
	})
}

// 退出
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "logout success",
	})
}
