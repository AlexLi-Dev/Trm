// Author: Lutong.li
package cluster

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AlexLi-Dev/Trm/config"
	"github.com/AlexLi-Dev/Trm/utils/changedatastyle"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/AlexLi-Dev/Trm/utils/returndata"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func addOrUpdate(c *gin.Context, method string) {
	var arg string
	if method == "create" {
		arg = "添加"
	} else {
		arg = "更新"
	}
	logs.Info(nil, arg+"集群")
	var clusterConfig ClusterConfig
	if err := c.ShouldBindJSON(&clusterConfig); err != nil {
		logs.Error(map[string]interface{}{"addconluster function": "集群信息解析失败 " + err.Error()}, arg+"集群信息不完整")
		msg := arg + "集群失败" + err.Error()
		response := returndata.NewReturnData()
		response.Code = http.StatusBadRequest
		response.Message = msg
		c.JSON(http.StatusOK, response)
		return
	}
	fmt.Println("clusterConfig:", clusterConfig.Kubeconfig)
	clusterStatus, err := clusterConfig.GetClusterStatusV2()
	if err != nil {
		response := returndata.NewReturnData()
		response.Code = http.StatusInternalServerError
		response.Message = "无法获取该集群的信息" + err.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	clusterStatusMap, _ := changedatastyle.StructToMap(clusterStatus)
	fmt.Println(clusterStatusMap)

	logs.Info(map[string]interface{}{"clusterName": clusterConfig.DisplayName, "cluster_Id": clusterConfig.ID}, "集群信息解析成功，开始添加集群")
	// 创建集群信息的secret
	var clusterConfigSecret corev1.Secret
	clusterConfigSecret.Name = clusterConfig.DisplayName
	clusterConfigSecret.Namespace = config.MetadataNamespace
	//添加标签
	clusterConfigSecret.Labels = make(map[string]string)
	clusterConfigSecret.Labels["kubeasy.com/cluster.metadata"] = "true"
	clusterConfigSecret.Labels["name"] = clusterConfig.DisplayName

	// 添加注解，保存集群的配置信息
	//Todo：改造为通过json转为map，然后直接给annotaion进行赋值
	clusterConfigSecret.Annotations = make(map[string]string)
	clusterConfigSecret.Annotations = clusterStatusMap
	// 保存kubeconfig
	clusterConfigSecret.StringData = make(map[string]string)
	clusterConfigSecret.StringData["kubeconfig"] = clusterConfig.Kubeconfig
	if method == "create" {
		//创建secret
		_, err = config.InclusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Create(context.TODO(), &clusterConfigSecret, metav1.CreateOptions{})
	} else {
		//创建secret
		_, err = config.InclusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Update(context.TODO(), &clusterConfigSecret, metav1.UpdateOptions{})
	}
	if err != nil {
		msg := fmt.Sprintf("%s 集群信息%s失败", clusterConfigSecret.Name, arg)
		logs.Error(map[string]interface{}{"namespace": config.MetadataNamespace, "type": "secret"}, msg)
		response := returndata.NewReturnData()
		response.Code = http.StatusInternalServerError
		response.Message = msg + err.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	//创建成功
	response := returndata.NewReturnData()
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("%s 集群信息添加成功", clusterConfigSecret.Name)
	response.Data = clusterConfig
	c.JSON(http.StatusOK, response)
}
