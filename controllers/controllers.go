// Author: Lutong.li
// 控制器层 实现路由的处理逻辑
package controllers

// 定义全局的数据结构
type BasicInfo struct {
	ClusterId string `json:"clusterId"`
	NameSpace string `json:"namespace"`
	Name      string `json:"name"`
}
