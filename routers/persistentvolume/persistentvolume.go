// Author: Lutong.li
package persistentvolume

import (
	PersistentVolume "github.com/AlexLi-Dev/Trm/controllers/persistentvolume"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	persistentvolumeApiGroup := authApiGroup.Group("/persistentvolume")
	persistentvolumeApiGroup.POST("/add", PersistentVolume.AddPersistentVolume)
	persistentvolumeApiGroup.Any("/delete", PersistentVolume.DeletePersistentVolume)
	persistentvolumeApiGroup.Any("/deletelist", PersistentVolume.DeletePersistentVolumeList)
	persistentvolumeApiGroup.POST("/update", PersistentVolume.UpdatePersistentVolume)
	persistentvolumeApiGroup.GET("/get", PersistentVolume.GetPersistentVolume)
	persistentvolumeApiGroup.GET("/list", PersistentVolume.ListPersistentVolume)
}

//kubectl create PersistentVolume dpv1 --image=nginx:latest --dry-run=client -o json
