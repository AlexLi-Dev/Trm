// Author: Lutong.li
// 控制器层 实现路由的处理逻辑
package controllers

import (
	"fmt"
	"net/http"

	"github.com/AlexLi-Dev/Trm/config"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 定义全局的数据结构
type BasicInfo struct {
	ClusterId string `json:"clusterId"  form:"clusterId"`
	NameSpace string `json:"namespace" form:"namespace"`
	Name      string `json:"name" form:"name"`

	//实现更新功能，新增字段
	Item interface{} `json:"item"`

	// 删除列式pod
	DeleteList []string `json:"deleteList"`
}

func BasicInit(c *gin.Context, item any) (clientset *kubernetes.Clientset, basicinfo BasicInfo, err error) {
	logs.Info(nil, "初始化clientset")
	//var ns corev1.Namespace
	basicinfo.Item = item

	switch c.Request.Method {
	case http.MethodGet:
		//debug
		//fmt.Printf("method=%s\n", c.Request.Method)
		//fmt.Printf("content-type=%s\n", c.ContentType())
		//fmt.Printf("raw query=%s\n", c.Request.URL.RawQuery)
		//fmt.Printf("query map=%v\n", c.Request.URL.Query())
		err = c.ShouldBindQuery(&basicinfo)
	case http.MethodPost:
		//debug
		//fmt.Printf("method=%s\n", c.Request.Method)
		//fmt.Printf("content-type=%s\n", c.ContentType())
		//fmt.Printf("raw query=%s\n", c.Request.URL.RawQuery)
		//fmt.Printf("query map=%v\n", c.Request.URL.Query())
		err = c.ShouldBindJSON(&basicinfo)
	default:
		err = fmt.Errorf("不支持的请求方法: %s", c.Request.Method)
	}
	fmt.Println(basicinfo)
	if err != nil {
		return nil, basicinfo, fmt.Errorf("请求参数绑定失败: %w", err)
	}

	kubeconfig := config.ClusterKubeconfig[basicinfo.ClusterId]
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		return nil, basicinfo, fmt.Errorf("basicinit 块 kubeconfig解析错误 %s", err.Error())
	}

	clientset, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, basicinfo, fmt.Errorf("basicinit 块 客户端工具创建失败  %s", err.Error())
	}
	return clientset, basicinfo, nil
}
