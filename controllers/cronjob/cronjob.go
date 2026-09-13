// Author: Lutong.li
package cronjob

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	batchv1 "k8s.io/api/batch/v1"
)

// 增加CronJob
func AddCronJob(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless service

	logs.Info(nil, "添加CronJob")
	var CronJob batchv1.CronJob
	var info controllers.Info
	info.BasicInfo.Item = &CronJob
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewCronJob(kubeconfig, &CronJob)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新CronJob
func UpdateCronJob(c *gin.Context) {
	logs.Info(nil, "更新CronJob")
	var CronJob batchv1.CronJob
	var info controllers.Info
	info.BasicInfo.Item = &CronJob
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewCronJob(kubeconfig, &CronJob)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除CronJob
func DeleteCronJob(c *gin.Context) {
	logs.Info(nil, "删除CronJob")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewCronJob(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeleteCronJobList(c *gin.Context) {
	logs.Info(nil, "列式删除CronJob")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewCronJob(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeleCronJobList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询CronJob详情
func GetCronJob(c *gin.Context) {
	logs.Info(nil, "查询CronJob")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewCronJob(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有CronJob
func ListCronJob(c *gin.Context) {
	logs.Info(nil, "列出所有CronJob list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewCronJob(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
