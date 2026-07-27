// Author: Lutong.li
package cluster

import (
	"github.com/AlexLi-Dev/Trm/controllers/cluster"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(authApiGroup *gin.RouterGroup) {
	//cluster相关操作的api
	//clusterApiGroup := authApiGroup.Group("/cluster").Use(func(c *gin.Context) {
	//	fmt.Println("======================经过中间价===================================")
	//	response := returndata.NewReturnData()
	//	// 获取节点列表
	//	nodes, err := config.InclusterClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{
	//		Limit: 1, // 只获取一个节点，减少开销
	//	})
	//	if err != nil {
	//		return
	//	}
	//
	//	// 检查是否有节点
	//	if len(nodes.Items) == 0 {
	//		response.Code = http.StatusBadGateway
	//		response.Message = "网关不可用，incluster集群报错！"
	//		logs.Error(nil, "incluster 集群报错，请检查集群情况！")
	//		return
	//	}
	//	c.Next()
	//})
	clusterApiGroup := authApiGroup.Group("/cluster")
	clusterApiGroup.POST("/add", cluster.AddCluster)
	clusterApiGroup.GET("/delete", cluster.DeleteCluster)
	clusterApiGroup.POST("/update", cluster.UpdateCluster)
	clusterApiGroup.GET("/get", cluster.GetCluster)
	clusterApiGroup.GET("/list", cluster.ListCluster)
}
