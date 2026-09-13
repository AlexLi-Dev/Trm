// Author: Lutong.li
package configmap

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
)

// 增加ConfigMap
func AddConfigMap(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless ConfigMap

	logs.Info(nil, "添加ConfigMap")
	var ConfigMap v1.ConfigMap
	var info controllers.Info
	info.BasicInfo.Item = &ConfigMap
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewConfigMap(kubeconfig, &ConfigMap)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新ConfigMap
func UpdateConfigMap(c *gin.Context) {
	logs.Info(nil, "更新ConfigMap")
	var ConfigMap v1.ConfigMap
	var info controllers.Info
	info.BasicInfo.Item = &ConfigMap
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewConfigMap(kubeconfig, &ConfigMap)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除ConfigMap
func DeleteConfigMap(c *gin.Context) {
	logs.Info(nil, "删除ConfigMap")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewConfigMap(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeleteConfigMapList(c *gin.Context) {
	logs.Info(nil, "列式删除ConfigMap")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewConfigMap(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeleConfigMapList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询ConfigMap详情
func GetConfigMap(c *gin.Context) {
	logs.Info(nil, "查询ConfigMap")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewConfigMap(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有ConfigMap
func ListConfigMap(c *gin.Context) {
	logs.Info(nil, "列出所有ConfigMap list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewConfigMap(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
