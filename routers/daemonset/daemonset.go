// Author: Lutong.li
package daemonset

import (
	DaemonSet "github.com/AlexLi-Dev/Trm/controllers/daemonset"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	DaemonSetApiGroup := authApiGroup.Group("/daemonset")
	DaemonSetApiGroup.POST("/add", DaemonSet.AddDaemonSet)
	DaemonSetApiGroup.Any("/delete", DaemonSet.DeleteDaemonSet)
	DaemonSetApiGroup.Any("/deletelist", DaemonSet.DeleteDaemonSetList)
	DaemonSetApiGroup.POST("/update", DaemonSet.UpdateDaemonSet)
	DaemonSetApiGroup.GET("/get", DaemonSet.GetDaemonSet)
	DaemonSetApiGroup.GET("/list", DaemonSet.ListDaemonSet)
}

//kubectl create statefulset dpv1 --image=nginx:latest --dry-run=client -o json
