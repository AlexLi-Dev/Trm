// Author: Lutong.li
package service

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
)

// 增加Service
func AddService(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless service

	logs.Info(nil, "添加Service")
	var Service v1.Service
	var info controllers.Info
	info.BasicInfo.Item = &Service
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewService(kubeconfig, &Service)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新Service
func UpdateService(c *gin.Context) {
	logs.Info(nil, "更新Service")
	var Service v1.Service
	var info controllers.Info
	info.BasicInfo.Item = &Service
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewService(kubeconfig, &Service)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除Service
func DeleteService(c *gin.Context) {
	logs.Info(nil, "删除Service")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewService(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeleteServiceList(c *gin.Context) {
	logs.Info(nil, "列式删除Service")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewService(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeleServiceList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询Service详情
func GetService(c *gin.Context) {
	logs.Info(nil, "查询Service")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewService(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有Service
func ListService(c *gin.Context) {
	logs.Info(nil, "列出所有Service list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewService(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
