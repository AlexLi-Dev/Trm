// Author: Lutong.li
package deployment

import (
	"github.com/AlexLi-Dev/Trm/controllers/deployment"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	statefulsetApiGroup := authApiGroup.Group("/deployment")
	statefulsetApiGroup.POST("/add", deployment.Adddeployment)
	statefulsetApiGroup.Any("/delete", deployment.Deletedeployment)
	statefulsetApiGroup.Any("/deletelist", deployment.DeletedeploymentList)
	statefulsetApiGroup.POST("/update", deployment.Updatedeployment)
	statefulsetApiGroup.GET("/get", deployment.Getdeployment)
	statefulsetApiGroup.GET("/list", deployment.Listdeployment)
}

//kubectl create deployment dpv1 --image=nginx:latest --dry-run=client -o json
