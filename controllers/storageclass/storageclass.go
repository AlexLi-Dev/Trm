// Author: Lutong.li
package storageclass

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/storage/v1"
)

// 增加StorageClass
func AddStorageClass(c *gin.Context) {
	// todo 需要看一下是否需要自动创建headless StorageClass

	logs.Info(nil, "添加StorageClass")
	var StorageClass v1.StorageClass
	var info controllers.Info
	info.BasicInfo.Item = &StorageClass
	kubeconfig, err := controllers.NewInfo(c, &info, "创建成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStorageClass(kubeconfig, &StorageClass)
	kubeUtilser = instance
	info.Create(c, kubeUtilser)
}

// 更新StorageClass
func UpdateStorageClass(c *gin.Context) {
	logs.Info(nil, "更新StorageClass")
	var StorageClass v1.StorageClass
	var info controllers.Info
	info.BasicInfo.Item = &StorageClass
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStorageClass(kubeconfig, &StorageClass)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 删除StorageClass
func DeleteStorageClass(c *gin.Context) {
	logs.Info(nil, "删除StorageClass")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStorageClass(kubeconfig, nil)
	kubeUtilser = instance
	info.Delete(c, kubeUtilser)
}

func DeleteStorageClassList(c *gin.Context) {
	logs.Info(nil, "列式删除StorageClass")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "删除成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser controllers.KubeUtilserDef
	inner := kubeutils.NewStorageClass(kubeconfig, nil) // kubeutils.KubeUtilser
	instance := &controllers.DelepodList{
		inner,
	} // 包装成 DeleStorageClassList
	kubeUtilser = instance
	info.DeleteList(c, kubeUtilser)
}

// 查询StorageClass详情
func GetStorageClass(c *gin.Context) {
	logs.Info(nil, "查询StorageClass")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStorageClass(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有StorageClass
func ListStorageClass(c *gin.Context) {
	logs.Info(nil, "列出所有StorageClass list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStorageClass(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
