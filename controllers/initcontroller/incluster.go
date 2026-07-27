// Author: Lutong.li
package initcontroller

import (
	"context"
	"fmt"

	"github.com/AlexLi-Dev/Trm/config"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func metadataInit() {
	//初始化元数据命名空间
	logs.Debug(nil, "初始化元数据命名空间")
	restconfig, err := clientcmd.BuildConfigFromFlags("https://127.0.0.1:6443", "config/config-k8s")
	if err != nil {
		logs.Error(map[string]interface{}{"msg": err.Error()}, "incluster kubeconfig 加载失败")
		panic(err.Error())
	}

	// 2. 创建客户端工具
	clientset, err := kubernetes.NewForConfig(restconfig)
	if err != nil {
		logs.Error(map[string]interface{}{"msg": err.Error()}, "incluster clientset 创建失败")
		panic(err.Error())
	}
	// 初始化时，将创建的clientset保存给config里面的全局，避免重复创建
	// todo 通过单例模式改造
	config.InclusterClientSet = clientset
	inclusterVersion, err := clientset.Discovery().ServerVersion()
	fmt.Println(inclusterVersion)
	if err != nil {
		panic(err.Error())
	}

	//3. 检查元数据命名空间是否存在
	_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), config.MetadataNamespace, metav1.GetOptions{})
	if err != nil {
		logs.Info(map[string]interface{}{"msg": err.Error()}, "元数据命名空间不存在，即将创建")
		var MetadataNamespace corev1.Namespace
		MetadataNamespace.Name = config.MetadataNamespace
		_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &MetadataNamespace, metav1.CreateOptions{})
		if err != nil {
			logs.Error(map[string]interface{}{"msg": err.Error()}, "元数据命名空间创建错误")
			panic(err.Error())
		}
		logs.Info(map[string]interface{}{"namespace": config.MetadataNamespace}, "元数据命名空间创建成功，版本号为 "+inclusterVersion.String())
	} else {
		logs.Info(map[string]interface{}{"namespace": config.MetadataNamespace}, "元数据命名空间已存在,版本号为 "+inclusterVersion.String())
	}
}
