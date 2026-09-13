// Author: Lutong.li
package cronjob

import (
	"github.com/AlexLi-Dev/Trm/controllers/cronjob"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	CronJobApiGroup := authApiGroup.Group("/cronjob")
	CronJobApiGroup.POST("/add", cronjob.AddCronJob)
	CronJobApiGroup.Any("/delete", cronjob.DeleteCronJob)
	CronJobApiGroup.Any("/deletelist", cronjob.DeleteCronJobList)
	CronJobApiGroup.POST("/update", cronjob.UpdateCronJob)
	CronJobApiGroup.GET("/get", cronjob.GetCronJob)
	CronJobApiGroup.GET("/list", cronjob.ListCronJob)
}

//kubectl create statefulset dpv1 --image=nginx:latest --dry-run=client -o json
