// Author: Lutong.li
package persistentvolumeclaim

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
)

// 增加PersistentVolumeClaim
func AddPersistentVolumeClaim(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless PersistentVolumeClaim

	logs.Info(nil, "添加PersistentVolumeClaim")
	var PersistentVolumeClaim v1.PersistentVolumeClaim
	var info controllers.Info
	info.BasicInfo.Item = &PersistentVolumeClaim
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolumeClaim(kubeconfig, &PersistentVolumeClaim)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新PersistentVolumeClaim
func UpdatePersistentVolumeClaim(c *gin.Context) {
	logs.Info(nil, "更新PersistentVolumeClaim")
	var PersistentVolumeClaim v1.PersistentVolumeClaim
	var info controllers.Info
	info.BasicInfo.Item = &PersistentVolumeClaim
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolumeClaim(kubeconfig, &PersistentVolumeClaim)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除PersistentVolumeClaim
func DeletePersistentVolumeClaim(c *gin.Context) {
	logs.Info(nil, "删除PersistentVolumeClaim")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolumeClaim(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeletePersistentVolumeClaimList(c *gin.Context) {
	logs.Info(nil, "列式删除PersistentVolumeClaim")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewPersistentVolumeClaim(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DelePersistentVolumeClaimList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询PersistentVolumeClaim详情
func GetPersistentVolumeClaim(c *gin.Context) {
	logs.Info(nil, "查询PersistentVolumeClaim")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolumeClaim(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有PersistentVolumeClaim
func ListPersistentVolumeClaim(c *gin.Context) {
	logs.Info(nil, "列出所有PersistentVolumeClaim list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolumeClaim(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
