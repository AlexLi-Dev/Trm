// Author: Lutong.li
package pod

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

// 增加pod
func Addpod(c *gin.Context) {
	logs.Info(nil, "添加pod")
	var pod corev1.Pod
	respdata := returndata.NewReturnData()
	clientset, basicInfo, err := controllers.BasicInit(c, &pod)
	if err != nil {
		respdata.Code = 400
		respdata.Message = err.Error()
		c.JSON(http.StatusOK, respdata)
		return
	}
	namespace := basicInfo.NameSpace
	pod.ObjectMeta.Namespace = namespace // 这里需要补齐yaml中的namespace
	_, err = clientset.CoreV1().Pods(namespace).Create(context.TODO(), &pod, metav1.CreateOptions{})
	if err != nil {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("Addpod块 namespace创建失败 %s", err.Error())
		c.JSON(http.StatusOK, respdata)
		return
	}
	respdata.Code = 200
	respdata.Message = "Addpod块 namespace创建成功"
	c.JSON(http.StatusOK, respdata)
}

// 更新pod
func Updatepod(c *gin.Context) {
	logs.Info(nil, "更新pod")

}

// 删除pod
func Deletepod(c *gin.Context) {
	logs.Info(nil, "删除pod")
	respdata := returndata.NewReturnData()
	clientset, basicInfo, err := controllers.BasicInit(c, nil)
	if err != nil {
		respdata.Code = 400
		respdata.Message = err.Error()
		c.JSON(http.StatusOK, respdata)
		return
	}
	ctx := context.TODO()

	if strings.HasPrefix(strings.ToLower(basicInfo.NameSpace), "kube") {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("kube 开头是k8s系统的，不能删除，当前为: %s", basicInfo.Name)
		c.JSON(http.StatusOK, respdata)
		return
	}
	if basicInfo.Name != "" {
		err = clientset.CoreV1().Pods(basicInfo.NameSpace).Delete(ctx, basicInfo.Name, metav1.DeleteOptions{})
		if err != nil {
			respdata.Code = 400
			respdata.Message = fmt.Sprintf("Addpod块 pod删除失败 %s", err.Error())
			c.JSON(http.StatusOK, respdata)
			return
		}
	}

	//批量删除
	if len(basicInfo.DeleteList) > 0 {
		var failed []string
		for _, podname := range basicInfo.DeleteList {
			if err := clientset.CoreV1().Pods(basicInfo.NameSpace).Delete(ctx, podname, metav1.DeleteOptions{}); err != nil {
				failed = append(failed, fmt.Sprintf("%s: %v", podname, err))
			}
		}
		if len(failed) > 0 {
			respdata.Code = 400
			respdata.Message = fmt.Sprintf("部分 pod 删除失败: %v", failed)
			c.JSON(http.StatusOK, respdata)
			return
		}
	}

	respdata.Code = 200
	respdata.Message = "Addpod块 pod删除成功"
	c.JSON(http.StatusOK, respdata)
}

// 查询pod详情
func Getpod(c *gin.Context) {
	logs.Info(nil, "查询pod")

}

// 查询所有pod
func Listpod(c *gin.Context) {
	logs.Info(nil, "列出所有pod list信息")
	respdata := returndata.NewReturnData()
	clientset, basicInfo, err := controllers.BasicInit(c, nil)
	if err != nil {
		respdata.Code = 400
		respdata.Message = err.Error()
		c.JSON(http.StatusOK, respdata)
		return
	}

	namespaceList, err := clientset.CoreV1().Pods(basicInfo.NameSpace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("ListNamespace块 列出所有 pod 失败: %s", err.Error())
		c.JSON(http.StatusOK, respdata)
		return
	}
	data := make(map[string]interface{})
	data["items"] = namespaceList.Items

	respdata.Code = 200
	respdata.Message = "listpod块 查询成功"
	respdata.Data = data

	c.JSON(http.StatusOK, respdata)
}
