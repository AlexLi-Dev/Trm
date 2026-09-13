// Author: Lutong.li
package storageclass

import (
	"github.com/AlexLi-Dev/Trm/controllers/storageclass"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	storageclassApiGroup := authApiGroup.Group("/storageclass")
	storageclassApiGroup.POST("/add", storageclass.AddStorageClass)
	storageclassApiGroup.Any("/delete", storageclass.DeleteStorageClass)
	storageclassApiGroup.Any("/deletelist", storageclass.DeleteStorageClassList)
	storageclassApiGroup.POST("/update", storageclass.UpdateStorageClass)
	storageclassApiGroup.GET("/get", storageclass.GetStorageClass)
	storageclassApiGroup.GET("/list", storageclass.ListStorageClass)
}

//kubectl create PersistentVolume dpv1 --image=nginx:latest --dry-run=client -o json
