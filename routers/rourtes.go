// Author: Lutong.li
package routers

import (
	"github.com/AlexLi-Dev/Trm/routers/cluster"
	"github.com/AlexLi-Dev/Trm/routers/namespace"
	"github.com/gin-gonic/gin"
)
import "github.com/AlexLi-Dev/Trm/routers/auth"

func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api/")
	auth.RegisterSubRouter(apiGroup)
	cluster.RegisterSubRouter(apiGroup)
	namespace.RegisterSubRouter(apiGroup)
}
