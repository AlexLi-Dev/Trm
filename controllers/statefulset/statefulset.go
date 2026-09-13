// Author: Lutong.li
package statefulset

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appv1 "k8s.io/api/apps/v1"
)

// 增加StatefulSet
func AddStatefulSet(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless service

	logs.Info(nil, "添加StatefulSet")
	var StatefulSet appv1.StatefulSet
	var info controllers.Info
	info.BasicInfo.Item = &StatefulSet
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, &StatefulSet)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新StatefulSet
func UpdateStatefulSet(c *gin.Context) {
	logs.Info(nil, "更新StatefulSet")
	var StatefulSet appv1.StatefulSet
	var info controllers.Info
	info.BasicInfo.Item = &StatefulSet
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, &StatefulSet)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除StatefulSet
func DeleteStatefulSet(c *gin.Context) {
	logs.Info(nil, "删除StatefulSet")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeleteStatefulSetList(c *gin.Context) {
	logs.Info(nil, "列式删除StatefulSet")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewStatefulSet(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeleStatefulSetList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询StatefulSet详情
func GetStatefulSet(c *gin.Context) {
	logs.Info(nil, "查询StatefulSet")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有StatefulSet
func ListStatefulSet(c *gin.Context) {
	logs.Info(nil, "列出所有StatefulSet list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
