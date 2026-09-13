// Author: Lutong.li
package node

import (
	"net/http"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
)

// 更新Node
func UpdateNode(c *gin.Context) {
	logs.Info(nil, "更新Node")
	var Node v1.Node
	var info controllers.Info
	info.BasicInfo.Item = &Node
	kubeconfig, err := controllers.NewInfo(c, &info, "更新成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewNode(kubeconfig, &Node)
	kubeUtilser = instance
	info.Update(c, kubeUtilser)

}

// 查询Node详情
func GetNode(c *gin.Context) {
	logs.Info(nil, "查询Node")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "获取成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewNode(kubeconfig, nil)
	kubeUtilser = instance
	info.Get(c, kubeUtilser)

}

// 查询所有Node
func ListNode(c *gin.Context) {
	logs.Info(nil, "列出所有Node list信息")
	var info controllers.Info
	kubeconfig, err := controllers.NewInfo(c, &info, "查询成功")
	if err != nil {
		c.JSON(http.StatusOK, info.ReturnData)
	}
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewNode(kubeconfig, nil)
	kubeUtilser = instance
	info.List(c, kubeUtilser)

}
