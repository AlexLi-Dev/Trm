// Author: Lutong.li
package auth

import (
	"github.com/AlexLi-Dev/Trm/controllers/auth"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	//用户登陆登出相关操作的api
	authApiGroup.POST("/login", auth.Login)
	authApiGroup.GET("/logout", auth.Logout)
}
