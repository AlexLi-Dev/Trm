// Author: Lutong.li
package ingress

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/networking/v1"
)

// 增加Ingress
func AddIngress(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless Ingress

	logs.Info(nil, "添加Ingress")
	var Ingress v1.Ingress
	var info controllers.Info
	info.BasicInfo.Item = &Ingress
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewIngress(kubeconfig, &Ingress)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新Ingress
func UpdateIngress(c *gin.Context) {
	logs.Info(nil, "更新Ingress")
	var Ingress v1.Ingress
	var info controllers.Info
	info.BasicInfo.Item = &Ingress
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewIngress(kubeconfig, &Ingress)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除Ingress
func DeleteIngress(c *gin.Context) {
	logs.Info(nil, "删除Ingress")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewIngress(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeleteIngressList(c *gin.Context) {
	logs.Info(nil, "列式删除Ingress")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewIngress(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeleIngressList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询Ingress详情
func GetIngress(c *gin.Context) {
	logs.Info(nil, "查询Ingress")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewIngress(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有Ingress
func ListIngress(c *gin.Context) {
	logs.Info(nil, "列出所有Ingress list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewIngress(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
