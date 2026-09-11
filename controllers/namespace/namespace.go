// Author: Lutong.li
package namespace

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/AlexLi-Dev/Trm/utils/returndata"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 增加namespace
func AddNameSpace(c *gin.Context) {
	logs.Info(nil, "添加NameSpace")
	//basicInfo := controllers.BasicInfo{}
	//respdata := returndata.NewReturnData()
	//if err := c.ShouldBindJSON(&basicInfo); err != nil {
	//	respdata.Code = 400
	//	respdata.Message = fmt.Sprintf("请求参数传递错误 %s", err.Error())
	//	c.JSON(http.StatusOK, respdata)
	//	return
	//}
	respdata := returndata.NewReturnData()
	clientset, basicInfo, err := controllers.BasicInit(c)
	if err != nil {
		respdata.Code = 400
		respdata.Message = err.Error()
		c.JSON(http.StatusOK, respdata)
		return
	}

	// 判断集群
	//todo 2260729
	//kubeconfig := config.ClusterKubeconfig[basicInfo.ClusterId]
	//restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	//if err != nil {
	//	respdata.Code = 400
	//	respdata.Message = fmt.Sprintf("AddNamespace块 kubeconfig解析错误 %s", err.Error())
	//	c.JSON(http.StatusOK, respdata)
	//	return
	//}

	//clientset, err := kubernetes.NewForConfig(restConfig)
	//if err != nil {
	//	respdata.Code = 400
	//	respdata.Message = fmt.Sprintf("AddNamespace块 客户端工具创建失败  %s", err.Error())
	//	c.JSON(http.StatusOK, respdata)
	//	return
	//}
	var namespace corev1.Namespace
	namespace.Name = basicInfo.Name
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
	if err != nil {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("AddNamespace块 namespace创建失败 %s", err.Error())
		c.JSON(http.StatusOK, respdata)
		return
	}
	respdata.Code = 200
	respdata.Message = "AddNamespace块 namespace创建成功"
	c.JSON(http.StatusOK, respdata)
}

// 更新namespace
func UpdateNameSpace(c *gin.Context) {
	logs.Info(nil, "更新namespace")
	respdata := returndata.NewReturnData()
	clientset, basicInfo, err := controllers.BasicInit(c)
	if err != nil {
		respdata.Code = 400
		respdata.Message = err.Error()
		c.JSON(http.StatusOK, respdata)
		return
	}

	ns, _ := basicInfo.Item.(*corev1.Namespace)
	updatens, err := clientset.CoreV1().Namespaces().Update(context.TODO(), ns, metav1.UpdateOptions{})
	if err != nil {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("updateNamespace块 namespace删除失败 %s", err.Error())
		c.JSON(http.StatusOK, respdata)
		return
	}
	data := make(map[string]interface{})
	data["item"] = updatens

	respdata.Code = 200
	respdata.Message = "updateNamespace块 namespace更新成功"
	respdata.Data = data
	c.JSON(http.StatusOK, respdata)

}

// 删除namespace
func DeleteNameSpace(c *gin.Context) {
	logs.Info(nil, "删除namespace")
	respdata := returndata.NewReturnData()
	clientset, basicInfo, err := controllers.BasicInit(c)
	if err != nil {
		respdata.Code = 400
		respdata.Message = err.Error()
		c.JSON(http.StatusOK, respdata)
		return
	}
	// 正常过滤方式比strings方式会慢一些
	//var kubePrefixRe = regexp.MustCompile(`(?i)^kube`)
	//if kubePrefixRe.MatchString(basicInfo.Name) {
	//	respdata.Code = 400
	//	respdata.Message = fmt.Sprintf("kube 开头是k8s系统的，不能删除，当前为: %s", basicInfo.Name)
	//	c.JSON(http.StatusOK, respdata)
	//	return
	//}

	if strings.HasPrefix(strings.ToLower(basicInfo.Name), "kube") {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("kube 开头是k8s系统的，不能删除，当前为: %s", basicInfo.Name)
		c.JSON(http.StatusOK, respdata)
		return
	}

	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), basicInfo.Name, metav1.DeleteOptions{})
	if err != nil {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("AddNamespace块 namespace删除失败 %s", err.Error())
		c.JSON(http.StatusOK, respdata)
		return
	}
	respdata.Code = 200
	respdata.Message = "AddNamespace块 namespace删除成功"
	c.JSON(http.StatusOK, respdata)

}

// 查询namespace详情
func GetNameSpace(c *gin.Context) {
	logs.Info(nil, "查询namespace")
	respdata := returndata.NewReturnData()
	clientset, basicInfo, err := controllers.BasicInit(c)
	if err != nil {
		respdata.Code = 400
		respdata.Message = err.Error()
		c.JSON(http.StatusOK, respdata)
		return
	}

	namespace, err := clientset.CoreV1().Namespaces().Get(context.TODO(), basicInfo.Name, metav1.GetOptions{})
	if err != nil {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("ListNamespace块 查询 namespace失败: %s", err.Error())
		c.JSON(http.StatusOK, respdata)
		return
	}

	data := make(map[string]interface{})
	data["item"] = namespace

	respdata.Code = 200
	respdata.Message = "listNamespace块 查询成功"
	respdata.Data = data
	c.JSON(http.StatusOK, respdata)
}

// 查询所有namespace
func ListNameSpace(c *gin.Context) {
	logs.Info(nil, "列出所有集群信息")
	respdata := returndata.NewReturnData()
	clientset, _, err := controllers.BasicInit(c)
	if err != nil {
		respdata.Code = 400
		respdata.Message = err.Error()
		c.JSON(http.StatusOK, respdata)
		return
	}

	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("ListNamespace块 列出所有 namespace失败: %s", err.Error())
		c.JSON(http.StatusOK, respdata)
		return
	}
	data := make(map[string]interface{})
	data["items"] = namespaceList.Items

	respdata.Code = 200
	respdata.Message = "listNamespace块 查询成功"
	respdata.Data = data

	c.JSON(http.StatusOK, respdata)
}
