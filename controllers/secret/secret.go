// Author: Lutong.li
package secret

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
)

// 增加Secret
func AddSecret(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless Secret

	logs.Info(nil, "添加Secret")
	var Secret v1.Secret
	var info controllers.Info
	info.BasicInfo.Item = &Secret
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewSecret(kubeconfig, &Secret)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新Secret
func UpdateSecret(c *gin.Context) {
	logs.Info(nil, "更新Secret")
	var Secret v1.Secret
	var info controllers.Info
	info.BasicInfo.Item = &Secret
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewSecret(kubeconfig, &Secret)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除Secret
func DeleteSecret(c *gin.Context) {
	logs.Info(nil, "删除Secret")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewSecret(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeleteSecretList(c *gin.Context) {
	logs.Info(nil, "列式删除Secret")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewSecret(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeleSecretList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询Secret详情
func GetSecret(c *gin.Context) {
	logs.Info(nil, "查询Secret")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewSecret(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有Secret
func ListSecret(c *gin.Context) {
	logs.Info(nil, "列出所有Secret list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewSecret(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
