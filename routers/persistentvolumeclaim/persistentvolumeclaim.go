// Author: Lutong.li
package persistentvolumeclaim

import (
	"github.com/AlexLi-Dev/Trm/controllers/persistentvolumeclaim"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	persistentvolumeclaimApiGroup := authApiGroup.Group("/persistentvolumeclaim")
	persistentvolumeclaimApiGroup.POST("/add", persistentvolumeclaim.AddPersistentVolumeClaim)
	persistentvolumeclaimApiGroup.Any("/delete", persistentvolumeclaim.DeletePersistentVolumeClaim)
	persistentvolumeclaimApiGroup.Any("/deletelist", persistentvolumeclaim.DeletePersistentVolumeClaimList)
	persistentvolumeclaimApiGroup.POST("/update", persistentvolumeclaim.UpdatePersistentVolumeClaim)
	persistentvolumeclaimApiGroup.GET("/get", persistentvolumeclaim.GetPersistentVolumeClaim)
	persistentvolumeclaimApiGroup.GET("/list", persistentvolumeclaim.ListPersistentVolumeClaim)
}

//kubectl create PersistentVolume dpv1 --image=nginx:latest --dry-run=client -o json
