//go:build k8s

// Author: Lutong.li
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1. 初始化config实例
	config, err := clientcmd.BuildConfigFromFlags("https://127.0.0.1:6443", "config/config-k8s")
	if err != nil {
		panic(err.Error())
	}

	// 2. 创建客户端工具
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pod1, err := clientset.CoreV1().Pods("kube-system").Get(context.Background(), "coredns-66bc5c9577-6b9dj", metav1.GetOptions{
		TypeMeta: metav1.TypeMeta{
			Kind: "Pod",
		},
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(pod1.Name, pod1.Spec.Containers[0].Image)

	// 3.操作集群
	podList, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind: "Pod",
		},
	})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range podList.Items {
		fmt.Println(pod.Namespace, pod.Name)
	}
	fmt.Println("共计：" + strconv.Itoa(len(podList.Items)) + " pod")

	// 4 创建一个nginx的pod
	//podNginx := &corev1.Pod{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:      "nginx-pod",
	//		Namespace: "default",
	//		Labels: map[string]string{
	//			"app": "nginx",
	//			"env": "production",
	//		},
	//	},
	//	Spec: corev1.PodSpec{
	//		Containers: []corev1.Container{
	//			{
	//				Name:  "nginx",
	//				Image: "nginx:latest",
	//				Ports: []corev1.ContainerPort{
	//					{
	//						ContainerPort: 80,
	//						Protocol:      corev1.ProtocolTCP,
	//					},
	//				},
	//			},
	//		},
	//		RestartPolicy: corev1.RestartPolicyAlways,
	//	},
	//}
	//podnginxInfo, err := clientset.CoreV1().Pods("default").Create(context.Background(), podNginx, metav1.CreateOptions{})
	//if err != nil {
	//	fmt.Println(err.Error())
	//	fmt.Println(podNginx.Name)
	//	podnginxInfo, _ := clientset.CoreV1().Pods("default").Get(context.Background(), podNginx.Name, metav1.GetOptions{})
	//	//podnginxInfo.Name = podNginx.Name + "-update"
	//
	//	fmt.Println(podnginxInfo.Status.PodIP)
	//	_, err := clientset.CoreV1().Pods("default").Update(context.Background(), podnginxInfo, metav1.UpdateOptions
	//	})
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//} else {
	//	fmt.Println(podnginxInfo.Name)
	//}

	//time.Sleep(10 * time.Second)
	//err = clientset.CoreV1().Pods("").Delete(context.TODO(), "nginx-pod", metav1.DeleteOptions{})
	//if err == nil {
	//	fmt.Println(podNginx.Name + "deleted")
	//} else {
	//	fmt.Println(err.Error())
	//}

	//deploy, _ := clientset.AppsV1().Deployments("default").Get(context.Background(), "nginx-deploy", metav1.GetOptions{})
	//fmt.Println(deploy.Spec)

	createredisjson := `{
    "kind": "Deployment",
    "apiVersion": "apps/v1",
    "metadata": {
        "name": "redis-deploy",
        "labels": {
            "app": "redis-deploy"
        }
    },
    "spec": {
        "replicas": 1,
        "selector": {
            "matchLabels": {
                "app": "redis-deploy"
            }
        },
        "template": {
            "metadata": {
                "labels": {
                    "app": "redis-deploy"
                }
            },
            "spec": {
                "containers": [
                    {
                        "name": "redis",
                        "image": "redis",
                        "resources": {}
                    }
                ]
            }
        },
        "strategy": {}
    },
    "status": {}
}`
	redis_deploy := appsv1.Deployment{}
	err = json.Unmarshal([]byte(createredisjson), &redis_deploy)
	fmt.Printf("%+v", redis_deploy)

}
