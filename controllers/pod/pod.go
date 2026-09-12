// Author: Lutong.li
package pod

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

// 增加pod
func Addpod(c *gin.Context) {
	logs.Info(nil, "添加pod")
	var pod corev1.Pod
	var info controllers.Info
	info.BasicInfo.Item = &pod
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, &pod)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新pod
func Updatepod(c *gin.Context) {
	logs.Info(nil, "更新pod")

}

// 删除pod
func Deletepod(c *gin.Context) {
	logs.Info(nil, "删除pod")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeletepodList(c *gin.Context) {
	logs.Info(nil, "列式删除pod")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewPod(kubeconfig, nil)   // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{*inner} // 包装成 DelepodList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询pod详情
func Getpod(c *gin.Context) {
	logs.Info(nil, "查询pod")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有pod
func Listpod(c *gin.Context) {
	logs.Info(nil, "列出所有pod list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
