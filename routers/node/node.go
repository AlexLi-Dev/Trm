// Author: Lutong.li
package node

import (
	"github.com/AlexLi-Dev/Trm/controllers/node"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	nodeApiGroup := authApiGroup.Group("/node")
	nodeApiGroup.POST("/update", node.UpdateNode)
	nodeApiGroup.GET("/get", node.GetNode)
	nodeApiGroup.GET("/list", node.ListNode)
}

//kubectl create ReplicaSet dpv1 --image=nginx:latest --dry-run=client -o json
