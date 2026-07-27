// Author: Lutong.li
package main

import (
	"github.com/AlexLi-Dev/Trm/config"
	_ "github.com/AlexLi-Dev/Trm/controllers/initcontroller"
	"github.com/AlexLi-Dev/Trm/routers"
	"github.com/AlexLi-Dev/Trm/utils/logs"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载程序的配置

	//配置gin
	r := gin.Default()
	routers.RegisterRouters(r)
	conf := config.NewConfig()
	conf.LoadDefault()
	logs.Debug(nil, "程序启动成功")
	r.Run(conf.Port)

}
