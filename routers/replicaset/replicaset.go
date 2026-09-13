// Author: Lutong.li
package replicaset

import (
	"github.com/AlexLi-Dev/Trm/controllers/ReplicaSet"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	ReplicaSetApiGroup := authApiGroup.Group("/replicaset")
	ReplicaSetApiGroup.POST("/add", ReplicaSet.AddReplicaSet)
	ReplicaSetApiGroup.Any("/delete", ReplicaSet.DeleteReplicaSet)
	ReplicaSetApiGroup.Any("/deletelist", ReplicaSet.DeleteReplicaSetList)
	ReplicaSetApiGroup.POST("/update", ReplicaSet.UpdateReplicaSet)
	ReplicaSetApiGroup.GET("/get", ReplicaSet.GetReplicaSet)
	ReplicaSetApiGroup.GET("/list", ReplicaSet.ListReplicaSet)
}

//kubectl create ReplicaSet dpv1 --image=nginx:latest --dry-run=client -o json
