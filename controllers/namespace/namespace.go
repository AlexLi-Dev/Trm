// Author: Lutong.li
package namespace

import (
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/gin-gonic/gin"
)

func AddNameSpace(c *gin.Context) {
	logs.Info(nil, "添加集群")
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
