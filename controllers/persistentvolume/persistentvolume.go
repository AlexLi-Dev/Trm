// Author: Lutong.li
package persistentvolume

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
)

// 增加PersistentVolume
func AddPersistentVolume(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless PersistentVolume

	logs.Info(nil, "添加PersistentVolume")
	var PersistentVolume v1.PersistentVolume
	var info controllers.Info
	info.BasicInfo.Item = &PersistentVolume
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolume(kubeconfig, &PersistentVolume)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新PersistentVolume
func UpdatePersistentVolume(c *gin.Context) {
	logs.Info(nil, "更新PersistentVolume")
	var PersistentVolume v1.PersistentVolume
	var info controllers.Info
	info.BasicInfo.Item = &PersistentVolume
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolume(kubeconfig, &PersistentVolume)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除PersistentVolume
func DeletePersistentVolume(c *gin.Context) {
	logs.Info(nil, "删除PersistentVolume")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolume(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeletePersistentVolumeList(c *gin.Context) {
	logs.Info(nil, "列式删除PersistentVolume")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewPersistentVolume(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DelePersistentVolumeList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询PersistentVolume详情
func GetPersistentVolume(c *gin.Context) {
	logs.Info(nil, "查询PersistentVolume")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolume(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有PersistentVolume
func ListPersistentVolume(c *gin.Context) {
	logs.Info(nil, "列出所有PersistentVolume list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPersistentVolume(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
