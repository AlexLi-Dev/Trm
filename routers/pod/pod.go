// Author: Lutong.li
package pod

import (
	"github.com/AlexLi-Dev/Trm/controllers/pod"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	podApiGroup := authApiGroup.Group("/pod")
	podApiGroup.POST("/add", pod.Addpod)
	podApiGroup.Any("/delete", pod.Deletepod)
	podApiGroup.POST("/update", pod.Updatepod)
	podApiGroup.GET("/get", pod.Getpod)
	podApiGroup.GET("/list", pod.Listpod)
}
