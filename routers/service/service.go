// Author: Lutong.li
package service

import (
	"github.com/AlexLi-Dev/Trm/controllers/service"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	serviceApiGroup := authApiGroup.Group("/service")
	serviceApiGroup.POST("/add", service.AddService)
	serviceApiGroup.Any("/delete", service.DeleteService)
	serviceApiGroup.Any("/deletelist", service.DeleteServiceList)
	serviceApiGroup.POST("/update", service.UpdateService)
	serviceApiGroup.GET("/get", service.GetService)
	serviceApiGroup.GET("/list", service.ListService)
}

//kubectl create service dpv1 --image=nginx:latest --dry-run=client -o json
