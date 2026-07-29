// Author: Lutong.li
package namespace

import (
	"fmt"

	"github.com/AlexLi-Dev/Trm/controllers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/AlexLi-Dev/Trm/utils/returndata"
	"github.com/gin-gonic/gin"
)

func AddNameSpace(c *gin.Context) {
	logs.Info(nil, "添加NameSpace")
	basicInfo := controllers.BasicInfo{}
	respdata := returndata.NewReturnData()
	if err := c.ShouldBindJSON(&basicInfo); err != nil {
		respdata.Code = 400
		respdata.Message = fmt.Sprintf("请求参数传递错误 %s", err.Error())
		c.JSON(respdata.Code, respdata)
		return
	}

	// 判断集群
	//todo 2260729
	//kubeconfig := config.ClusterKubeconfig[basicInfo.ClusterId]

	respdata.Code = 200
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func UpdateNameSpace(c *gin.Context) {
	logs.Info(nil, "更新集群")
}

func DeleteNameSpace(c *gin.Context) {
	logs.Info(nil, "删除集群")
}

func GetNameSpace(c *gin.Context) {

}

func ListNameSpace(c *gin.Context) {
	logs.Info(nil, "列出所有集群信息")

}
