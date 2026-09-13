// Author: Lutong.li
package ingress

import (
	"github.com/AlexLi-Dev/Trm/controllers/Ingress"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	ingressApiGroup := authApiGroup.Group("/ingress")
	ingressApiGroup.POST("/add", ingress.AddIngress)
	ingressApiGroup.Any("/delete", ingress.DeleteIngress)
	ingressApiGroup.Any("/deletelist", ingress.DeleteIngressList)
	ingressApiGroup.POST("/update", ingress.UpdateIngress)
	ingressApiGroup.GET("/get", ingress.GetIngress)
	ingressApiGroup.GET("/list", ingress.ListIngress)
}

//kubectl create Ingress dpv1 --image=nginx:latest --dry-run=client -o json

// kubectl port-forward -n testa svc/dpv3 8080:80
//curl -H "Host: dpv3.example.com" http://localhost:8080
