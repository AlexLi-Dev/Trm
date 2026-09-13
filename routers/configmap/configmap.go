// Author: Lutong.li
package configmap

import (
	"github.com/AlexLi-Dev/Trm/controllers/ConfigMap"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	configmapApiGroup := authApiGroup.Group("/configmap")
	configmapApiGroup.POST("/add", configmap.AddConfigMap)
	configmapApiGroup.Any("/delete", configmap.DeleteConfigMap)
	configmapApiGroup.Any("/deletelist", configmap.DeleteConfigMapList)
	configmapApiGroup.POST("/update", configmap.UpdateConfigMap)
	configmapApiGroup.GET("/get", configmap.GetConfigMap)
	configmapApiGroup.GET("/list", configmap.ListConfigMap)
}

//kubectl create ConfigMap dpv1 --image=nginx:latest --dry-run=client -o json
