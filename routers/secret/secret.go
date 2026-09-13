// Author: Lutong.li
package secret

import (
	"github.com/AlexLi-Dev/Trm/controllers/secret"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	secretApiGroup := authApiGroup.Group("/secret")
	secretApiGroup.POST("/add", secret.AddSecret)
	secretApiGroup.Any("/delete", secret.DeleteSecret)
	secretApiGroup.Any("/deletelist", secret.DeleteSecretList)
	secretApiGroup.POST("/update", secret.UpdateSecret)
	secretApiGroup.GET("/get", secret.GetSecret)
	secretApiGroup.GET("/list", secret.ListSecret)
}

//kubectl create secret dpv1 --image=nginx:latest --dry-run=client -o json
