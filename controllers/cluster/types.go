// Author: Lutong.li
package cluster

import (
	"context"
	"fmt"

	"github.com/AlexLi-Dev/Trm/config"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 定义一个人结构体，用于描述创建集群所用的配置信息

type ClusterInfo struct {
	ID int `json:"id"`
	//Name string `json:"name"`
	DisplayName string `json:"displayname"` //别名
	City        string `json:"city"`        // 城市
	District    string `json:"district"`    //区域
}

type ClusterConfig struct {
	ClusterInfo
	Kubeconfig string `json:"kubeconfig"`
}

type ClusterStatus struct {
	ClusterInfo
	Version string `json:"version"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (c *ClusterInfo) GetClusterStatusV1() (ClusterStatus, error) {
	clusterStatus := ClusterStatus{
		ClusterInfo: *c,
		Status:      "Unknown",
	}

	// 1. 获取版本信息
	version, err := config.InclusterClientSet.Discovery().ServerVersion()
	if err != nil {
		clusterStatus.Status = "Error"
		clusterStatus.Message = "无法获取版本信息: " + err.Error()
		return clusterStatus, fmt.Errorf("获取版本失败: %w", err)
	}
	clusterStatus.Version = version.GitVersion

	// 2. 获取节点列表（只获取一个节点检查）
	nodes, err := config.InclusterClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{
		Limit: 1,
	})
	if err != nil {
		clusterStatus.Status = "Error"
		clusterStatus.Message = "无法获取节点信息: " + err.Error()
		return clusterStatus, fmt.Errorf("获取节点失败: %w", err)
	}

	// 3. 检查是否有节点
	if len(nodes.Items) == 0 {
		clusterStatus.Status = "Warning"
		clusterStatus.Message = "集群没有节点"
		return clusterStatus, fmt.Errorf("集群没有节点")
	}

	// 4. 检查节点状态
	node := nodes.Items[0]
	isReady := false
	for _, condition := range node.Status.Conditions {
		if condition.Type == corev1.NodeReady && condition.Status == corev1.ConditionTrue {
			isReady = true
			break
		}
	}

	// 5. 设置最终状态
	if isReady {
		clusterStatus.Status = "Ready"
		clusterStatus.Message = "集群运行正常"
	} else {
		clusterStatus.Status = "Warning"
		clusterStatus.Message = "节点未就绪"
	}

	return clusterStatus, nil
}

func (c *ClusterConfig) GetClusterStatusV2() (ClusterStatus, error) {
	//判断集群是否正常
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = c.ClusterInfo
	//创建一个clientset
	fmt.Println("----", c.Kubeconfig)
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		return clusterStatus, err
	}
	// 2. 创建 clientset
	clientset, _ := kubernetes.NewForConfig(restConfig)

	// 3. 获取版本信息
	version, err := clientset.Discovery().ServerVersion()

	clusterStatus.Version = version.String()

	clusterStatus.Status = "Ready"

	return clusterStatus, nil
}
