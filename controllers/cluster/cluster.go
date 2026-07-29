// Author: Lutong.li
package cluster

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/AlexLi-Dev/Trm/config"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/AlexLi-Dev/Trm/utils/returndata"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//func AddCluster(c *gin.Context) {
//	logs.Info(nil, "添加集群")
//	var clusterConfig ClusterConfig
//	if err := c.ShouldBindJSON(&clusterConfig); err != nil {
//		logs.Error(map[string]interface{}{"addconluster function": "集群信息解析失败 " + err.Error()}, "添加集群信息不完整")
//		msg := "添加集群失败" + err.Error()
//		response := returndata.NewReturnData()
//		response.Code = http.StatusBadRequest
//		response.Message = msg
//		c.JSON(http.StatusOK, response)
//		return
//	}
//	fmt.Println("clusterConfig:", clusterConfig.Kubeconfig)
//	clusterStatus, err := clusterConfig.GetClusterStatusV2()
//	if err != nil {
//		response := returndata.NewReturnData()
//		response.Code = http.StatusInternalServerError
//		response.Message = "无法获取该集群的信息" + err.Error()
//		c.JSON(http.StatusOK, response)
//		return
//	}
//	clusterStatusMap, _ := changedatastyle.StructToMap(clusterStatus)
//
//	logs.Info(map[string]interface{}{"clusterName": clusterConfig.DisplayName, "cluster_Id": clusterConfig.ID}, "集群信息解析成功，开始添加集群")
//	// 创建集群信息的secret
//	var clusterConfigSecret corev1.Secret
//	clusterConfigSecret.Name = clusterConfig.DisplayName
//	clusterConfigSecret.Namespace = config.MetadataNamespace
//	//添加标签
//	clusterConfigSecret.Labels = make(map[string]string)
//	clusterConfigSecret.Labels["kubeasy.com/cluster.metadata"] = "true"
//	clusterConfigSecret.Labels["name"] = clusterConfig.DisplayName
//
//	// 添加注解，保存集群的配置信息
//	//Todo：改造为通过json转为map，然后直接给annotaion进行赋值
//	clusterConfigSecret.Annotations = make(map[string]string)
//	clusterConfigSecret.Annotations = clusterStatusMap
//	//clusterConfigSecret.Annotations["displayname"] = clusterConfig.DisplayName
//	//clusterConfigSecret.Annotations["city"] = clusterConfig.City
//	//clusterConfigSecret.Annotations["district"] = clusterConfig.District
//	//clusterConfigSecret.Annotations["status"] = clusterStatus.Status
//	//clusterConfigSecret.Annotations["Version"] = clusterStatus.Version
//
//	// 保存kubeconfig
//	clusterConfigSecret.StringData = make(map[string]string)
//	clusterConfigSecret.StringData["kubeconfig"] = clusterConfig.Kubeconfig
//
//	//创建secret
//	_, err = config.InclusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Create(context.TODO(), &clusterConfigSecret, metav1.CreateOptions{})
//	if err != nil {
//		logs.Error(map[string]interface{}{"namespace": config.MetadataNamespace, "type": "secret"}, clusterConfigSecret.Name+" 集群信息创建失败")
//		response := returndata.NewReturnData()
//		response.Code = http.StatusInternalServerError
//		response.Message = clusterConfigSecret.Name + " 集群信息保存失败 " + err.Error()
//		c.JSON(http.StatusOK, response)
//		return
//	}
//
//	//创建成功
//	response := returndata.NewReturnData()
//	response.Code = http.StatusOK
//	response.Message = clusterConfigSecret.Name + "集群信息添加成功"
//	response.Data = clusterConfig
//	c.JSON(http.StatusOK, response)
//}

func AddCluster(c *gin.Context) {
	logs.Info(nil, "添加集群")
	addOrUpdate(c, "create")
}

func UpdateCluster(c *gin.Context) {
	logs.Info(nil, "更新集群")
	addOrUpdate(c, "update")
}

func DeleteCluster(c *gin.Context) {
	logs.Info(nil, "删除集群")

	// 1. 接受删除集群的id
	clusterName := c.Query("clusterName")

	//1.1 增加对前段param参数会多“”的情况等一种容错
	// 去除首尾的双引号
	//clusterName="localk8s-2"
	clusterName = strings.Trim(clusterName, `"`)
	// 或更通用的去除所有引号: clusterName=local"k8"s-2
	clusterName = strings.ReplaceAll(clusterName, `"`, "")

	if clusterName == "" {
		logs.Info(map[string]interface{}{"function": "DeleteCluster"}, "clustername传入了空值")

		response := returndata.NewReturnData()
		response.Code = http.StatusBadRequest
		response.Message = "clusterName is required"
		c.JSON(http.StatusOK, response)
		return
	}

	//clusterId := c.Query("clusterId")

	//2. 判断是否存在，存在就删除
	err := config.InclusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Delete(context.TODO(), clusterName, metav1.DeleteOptions{})
	response := returndata.NewReturnData()
	if err != nil {
		logs.Error(map[string]interface{}{"namespace": config.MetadataNamespace, "cluster": clusterName, "type": "secret"}, clusterName)

		response.Code = http.StatusInternalServerError
		response.Message = "删除失败" + err.Error()
	} else {
		delete(config.ClusterKubeconfig, clusterName)
		response.Code = http.StatusOK
		response.Message = clusterName + "删除成功"
	}

	c.JSON(200, response)
}

func GetCluster(c *gin.Context) {
	//todo 20260727 继续开始从这里写

	logs.Info(nil, "获取集群的详情")
	clusterName := c.Query("clusterName")
	clusterName = strings.Trim(clusterName, `"`)
	// 或更通用的去除所有引号: clusterName=local"k8"s-2
	clusterName = strings.ReplaceAll(clusterName, `"`, "")

	if clusterName == "" {
		logs.Info(map[string]interface{}{"function": "GetCluster"}, "clustername传入了空值")
		response := returndata.NewReturnData()
		response.Code = http.StatusBadRequest
		response.Message = "clusterName is required"
		c.JSON(http.StatusOK, response)
		return
	}
	clusterItem, err := config.InclusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Get(context.TODO(), clusterName, metav1.GetOptions{})
	response := returndata.NewReturnData()
	if err != nil {
		logs.Error(map[string]interface{}{"namespace": config.MetadataNamespace, "cluster": clusterName, "type": "secret"}, clusterName)

		response.Code = http.StatusInternalServerError
		response.Message = fmt.Sprintf("查询到%s集群不存在！", clusterName)
	} else {
		response.Code = http.StatusOK
		response.Message = fmt.Sprintf("%s查询成功", clusterName)
		info := clusterItem.Annotations
		info["kubeconfig"] = string(clusterItem.Data["kubeconfig"])
		response.Data = map[string]interface{}{
			"item": info,
		}
	}

	c.JSON(200, response)
}

func ListCluster(c *gin.Context) {
	logs.Info(nil, "列出所有集群信息")
	listOptions := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", config.ClusterConfigSecretLabelKey, config.ClusterConfigSecretLabelValue),
	}

	clusterList, err := config.InclusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(context.TODO(), listOptions)
	response := struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Data    []interface{} `json:"data"`
	}{}
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "查询所有集群的信息失败" + err.Error()
		c.JSON(http.StatusOK, response)
	}
	response.Code = http.StatusOK
	response.Message = "查询成功，所有集群信息如下"
	for _, cluster := range clusterList.Items {
		var info = make(map[string]interface{})
		info[cluster.Name] = map[string]interface{}{
			"clusterName": cluster.Name,
			"city":        cluster.Annotations["city"],
			"district":    cluster.Annotations["district"],
			"status":      cluster.Annotations["status"],
			"version":     cluster.Annotations["version"],
			"labels":      cluster.Labels,
		}
		response.Data = append(response.Data, info)
	}
	c.JSON(200, response)
}
