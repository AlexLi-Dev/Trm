// Author: Lutong.li
package statefulset

import (
	"github.com/AlexLi-Dev/Trm/controllers/statefulset"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	statefulsetApiGroup := authApiGroup.Group("/statefulset")
	statefulsetApiGroup.POST("/add", statefulset.AddStatefulSet)
	statefulsetApiGroup.Any("/delete", statefulset.DeleteStatefulSet)
	statefulsetApiGroup.Any("/deletelist", statefulset.DeleteStatefulSetList)
	statefulsetApiGroup.POST("/update", statefulset.UpdateStatefulSet)
	statefulsetApiGroup.GET("/get", statefulset.GetStatefulSet)
	statefulsetApiGroup.GET("/list", statefulset.ListStatefulSet)
}

//kubectl create statefulset dpv1 --image=nginx:latest --dry-run=client -o json
