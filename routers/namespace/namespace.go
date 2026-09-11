// Author: Lutong.li
package namespace

import (
	"github.com/AlexLi-Dev/Trm/controllers/namespace"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	namespaceApiGroup := authApiGroup.Group("/namespace")
	namespaceApiGroup.POST("/add", namespace.AddNameSpace)
	namespaceApiGroup.Any("/delete", namespace.DeleteNameSpace)
	namespaceApiGroup.POST("/update", namespace.UpdateNameSpace)
	namespaceApiGroup.GET("/get", namespace.GetNameSpace)
	namespaceApiGroup.GET("/list", namespace.ListNameSpace)
}
