// Author: Lutong.li
package temp

import (
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/gin-gonic/gin"
)

func AddNameSpace(c *gin.Context) {
	logs.Info(nil, "添加集群")
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
