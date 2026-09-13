// Author: Lutong.li
package ReplicaSet

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appv1 "k8s.io/api/apps/v1"
)

// 增加ReplicaSet
func AddReplicaSet(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless service

	logs.Info(nil, "添加ReplicaSet")
	var ReplicaSet appv1.ReplicaSet
	var info controllers.Info
	info.BasicInfo.Item = &ReplicaSet
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewReplicaSet(kubeconfig, &ReplicaSet)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新ReplicaSet
func UpdateReplicaSet(c *gin.Context) {
	logs.Info(nil, "更新ReplicaSet")
	var ReplicaSet appv1.ReplicaSet
	var info controllers.Info
	info.BasicInfo.Item = &ReplicaSet
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewReplicaSet(kubeconfig, &ReplicaSet)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除ReplicaSet
func DeleteReplicaSet(c *gin.Context) {
	logs.Info(nil, "删除ReplicaSet")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewReplicaSet(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeleteReplicaSetList(c *gin.Context) {
	logs.Info(nil, "列式删除ReplicaSet")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewReplicaSet(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeleReplicaSetList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询ReplicaSet详情
func GetReplicaSet(c *gin.Context) {
	logs.Info(nil, "查询ReplicaSet")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewReplicaSet(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有ReplicaSet
func ListReplicaSet(c *gin.Context) {
	logs.Info(nil, "列出所有ReplicaSet list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewReplicaSet(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
