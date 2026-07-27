// Author: Lutong.li
package initcontroller

import (
	"github.com/AlexLi-Dev/Trm/utils/logs"
)

func init() {
	logs.Debug(nil, "初始化数据")
	// 1. 通过kubeconfig文件去创建client-go的客户端
	// 2. 检查元数据命名空间是否已经创建
	/*
		 	a. 已经创建： 元数据命令空间已经创建了
			b. 未创建： 创建命名空间
	*/
	metadataInit()

}
