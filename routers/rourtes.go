// Author: Lutong.li
package routers

import (
	"github.com/AlexLi-Dev/Trm/routers/cluster"
	"github.com/AlexLi-Dev/Trm/routers/cronjob"
	"github.com/AlexLi-Dev/Trm/routers/daemonset"
	"github.com/AlexLi-Dev/Trm/routers/deployment"
	"github.com/AlexLi-Dev/Trm/routers/ingress"
	"github.com/AlexLi-Dev/Trm/routers/namespace"
	"github.com/AlexLi-Dev/Trm/routers/node"
	"github.com/AlexLi-Dev/Trm/routers/pod"
	"github.com/AlexLi-Dev/Trm/routers/replicaset"
	"github.com/AlexLi-Dev/Trm/routers/service"
	statefulset "github.com/AlexLi-Dev/Trm/routers/statefulset"
	"github.com/gin-gonic/gin"
)
import "github.com/AlexLi-Dev/Trm/routers/auth"

func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api/")
	auth.RegisterSubRouter(apiGroup)
	cluster.RegisterSubRouter(apiGroup)
	namespace.RegisterSubRouter(apiGroup)
	pod.RegisterSubRouter(apiGroup)
	deployment.RegisterSubRouter(apiGroup)
	statefulset.RegisterSubRouter(apiGroup)
	daemonset.RegisterSubRouter(apiGroup)
	cronjob.RegisterSubRouter(apiGroup)
	replicaset.RegisterSubRouter(apiGroup)
	node.RegisterSubRouter(apiGroup)
	service.RegisterSubRouter(apiGroup)
	ingress.RegisterSubRouter(apiGroup)
}

//kubectl get replicasets.apps -n testa
//
