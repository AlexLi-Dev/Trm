// Author: Lutong.li
package deployment

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appv1 "k8s.io/api/apps/v1"
)

// 增加deployment
func Adddeployment(c *gin.Context) {
	logs.Info(nil, "添加deployment")
	var deployment appv1.Deployment
	var info controllers.Info
	info.BasicInfo.Item = &deployment
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDeployment(kubeconfig, &deployment)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新deployment
func Updatedeployment(c *gin.Context) {
	logs.Info(nil, "更新deployment")
	var deployment appv1.Deployment
	var info controllers.Info
	info.BasicInfo.Item = &deployment
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDeployment(kubeconfig, &deployment)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除deployment
func Deletedeployment(c *gin.Context) {
	logs.Info(nil, "删除deployment")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDeployment(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeletedeploymentList(c *gin.Context) {
	logs.Info(nil, "列式删除deployment")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewDeployment(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeledeploymentList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询deployment详情
func Getdeployment(c *gin.Context) {
	logs.Info(nil, "查询deployment")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDeployment(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有deployment
func Listdeployment(c *gin.Context) {
	logs.Info(nil, "列出所有deployment list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDeployment(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
