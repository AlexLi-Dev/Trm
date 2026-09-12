// Author: Lutong.li
// 控制器层 实现路由的处理逻辑
package controllers

import (
	"fmt"
	"net/http"

	"github.com/AlexLi-Dev/Trm/config"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/AlexLi-Dev/Trm/utils/returndata"
	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/tools/clientcmd"
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

type Info struct {
	BasicInfo     BasicInfo `json:"basicInfo"`
	ReturnData    returndata.ReturnData
	LabelSelector string   `json:"labelSelector" form:"labelSelector"`
	FieldSelector string   `json:"fieldSelector" form:"fieldSelector"`
	ForceDeletion booltype `json:"forceDeletion" form:"forceDeletion"`
}

// ==== 改造为接口方式======

type booltype bool

func (booltype booltype) ToInt64() *int64 {
	var n int64
	if bool(booltype) {
		n = 1
	}
	return &n
}

func (b *Info) Create(c *gin.Context, kubeUtilsInstance kubeutils.KubeUtilser) {
	err := kubeUtilsInstance.Create(b.BasicInfo.NameSpace)
	if err != nil {
		b.ReturnData.Message = fmt.Sprintf("创建失败 %s", err.Error())
		b.ReturnData.Code = 400
		logs.Error(nil, b.ReturnData.Message)
		c.JSON(http.StatusOK, b.ReturnData)
		return
	}
	c.JSON(http.StatusOK, b.ReturnData)
	return
}

func (b *Info) Update(c *gin.Context, kubeUtilsInstance kubeutils.KubeUtilser) {
	err := kubeUtilsInstance.Update(b.BasicInfo.NameSpace)
	if err != nil {
		b.ReturnData.Message = fmt.Sprintf("更新失败 %s", err.Error())
		b.ReturnData.Code = 400
		logs.Error(nil, b.ReturnData.Message)
		c.JSON(http.StatusOK, b.ReturnData)
		return
	}
	c.JSON(http.StatusOK, b.ReturnData)
	return
}

func (b *Info) List(c *gin.Context, kubeUtilsInstance kubeutils.KubeUtilser) {
	items, err := kubeUtilsInstance.List(b.BasicInfo.NameSpace, b.LabelSelector, b.FieldSelector)
	if err != nil {
		b.ReturnData.Message = fmt.Sprintf("查询列表失败 %s", err.Error())
		b.ReturnData.Code = 400
		logs.Error(nil, b.ReturnData.Message)
	} else {
		data := make(map[string]interface{})
		data["items"] = items
		b.ReturnData.Data = data
	}
	c.JSON(http.StatusOK, b.ReturnData)
	return
}

func (b *Info) Get(c *gin.Context, kubeUtilsInstance kubeutils.KubeUtilser) {
	item, err := kubeUtilsInstance.Get(b.BasicInfo.NameSpace, b.BasicInfo.Name)
	if err != nil {
		b.ReturnData.Message = fmt.Sprintf("查询失败 %s", err.Error())
		b.ReturnData.Code = 400
		logs.Error(nil, b.ReturnData.Message)
	} else {
		data := make(map[string]interface{})
		data["item"] = item
		b.ReturnData.Data = data
	}
	c.JSON(http.StatusOK, b.ReturnData)
	return
}

func (b *Info) Delete(c *gin.Context, kubeUtilsInstance kubeutils.KubeUtilser) {
	err := kubeUtilsInstance.Delete(b.BasicInfo.NameSpace, b.BasicInfo.Name, b.ForceDeletion.ToInt64())
	if err != nil {
		b.ReturnData.Message = fmt.Sprintf("删除失败 %s", err.Error())
		b.ReturnData.Code = 400
		logs.Error(nil, b.ReturnData.Message)
	} else {
		b.ReturnData.Message = fmt.Sprintf("删除成功 %s", b.BasicInfo.Name)
		b.ReturnData.Code = 200
	}
	c.JSON(http.StatusOK, b.ReturnData)
	return
}

type KubeUtilserDef interface {
	kubeutils.KubeUtilser
	DeleteListwithfaild(string, []string, *int64) (failed []string)
}
type DelepodList struct {
	kubeutils.Pod
}

func (c DelepodList) DeleteListwithfaild(namespace string, nameList []string, gracePeriodSeconds *int64) (failed []string) {
	for _, name := range nameList {
		if err := c.Delete(namespace, name, gracePeriodSeconds); err != nil {
			failed = append(failed, fmt.Sprintf("%s:%s", name, err.Error()))
		}
	}
	return failed
}

func (b *Info) DeleteList(c *gin.Context, kubeUtilsInstance KubeUtilserDef) {
	fmt.Printf("name=[%s], deleteList=%v\n", b.BasicInfo.Name, b.BasicInfo.DeleteList)
	failed := kubeUtilsInstance.DeleteListwithfaild(b.BasicInfo.NameSpace, b.BasicInfo.DeleteList, b.ForceDeletion.ToInt64())
	if len(failed) > 0 {
		b.ReturnData.Message = fmt.Sprintf("删除失败 %s", failed)
		b.ReturnData.Code = 400
		data := make(map[string]interface{})
		data["items"] = failed
		logs.Error(nil, b.ReturnData.Message)
	} else {
		b.ReturnData.Message = fmt.Sprintf("删除成功 %s", b.BasicInfo.Name)
		b.ReturnData.Code = 200
	}
	c.JSON(http.StatusOK, b.ReturnData)
	return
}

func NewInfo(c *gin.Context, basicinfo *Info, returnDataMsg string) (kubeconfig string, err error) {
	logs.Info(nil, "newbasicinfo")
	basicinfo.ReturnData.Message = returnDataMsg
	basicinfo.ReturnData.Code = http.StatusOK
	switch c.Request.Method {
	case http.MethodGet:
		//debug
		//fmt.Printf("method=%s\n", c.Request.Method)
		//fmt.Printf("content-type=%s\n", c.ContentType())
		//fmt.Printf("raw query=%s\n", c.Request.URL.RawQuery)
		//fmt.Printf("query map=%v\n", c.Request.URL.Query())
		err = c.ShouldBindQuery(&basicinfo.BasicInfo)
	case http.MethodPost:
		//debug
		//fmt.Printf("method=%s\n", c.Request.Method)
		//fmt.Printf("content-type=%s\n", c.ContentType())
		//fmt.Printf("raw query=%s\n", c.Request.URL.RawQuery)
		//fmt.Printf("query map=%v\n", c.Request.URL.Query())
		err = c.ShouldBindJSON(&basicinfo.BasicInfo)
	default:
		err = fmt.Errorf("不支持的请求方法: %s", c.Request.Method)
	}
	if err != nil {
		basicinfo.ReturnData.Message = fmt.Sprintf(" newinfo 请求出错 %s", err.Error())
		basicinfo.ReturnData.Code = 400
		c.JSON(http.StatusOK, basicinfo.ReturnData)
		return
	}

	kubeconfig = config.ClusterKubeconfig[basicinfo.BasicInfo.ClusterId]
	return kubeconfig, nil
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
