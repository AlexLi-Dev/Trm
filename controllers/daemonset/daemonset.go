// Author: Lutong.li
package DaemonSet

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appv1 "k8s.io/api/apps/v1"
)

// 增加DaemonSet
func AddDaemonSet(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless service

	logs.Info(nil, "添加DaemonSet")
	var DaemonSet appv1.DaemonSet
	var info controllers.Info
	info.BasicInfo.Item = &DaemonSet
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, &DaemonSet)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新DaemonSet
func UpdateDaemonSet(c *gin.Context) {
	logs.Info(nil, "更新DaemonSet")
	var DaemonSet appv1.DaemonSet
	var info controllers.Info
	info.BasicInfo.Item = &DaemonSet
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, &DaemonSet)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除DaemonSet
func DeleteDaemonSet(c *gin.Context) {
	logs.Info(nil, "删除DaemonSet")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeleteDaemonSetList(c *gin.Context) {
	logs.Info(nil, "列式删除DaemonSet")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewDaemonSet(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeleDaemonSetList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询DaemonSet详情
func GetDaemonSet(c *gin.Context) {
	logs.Info(nil, "查询DaemonSet")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有DaemonSet
func ListDaemonSet(c *gin.Context) {
	logs.Info(nil, "列出所有DaemonSet list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
