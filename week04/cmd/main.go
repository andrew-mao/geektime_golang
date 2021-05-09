package main

import (
	"fmt"
	"github.com/pkg/errors"
	"mblog/cmd/conf"
)

func main() {

	//读取配置文件 "E:/learn/Mblog_go/cmd/conf.toml"
	err := conf.Init("")
	if err != nil {
		fmt.Printf("origin: %T \n %v", errors.Cause(err), errors.Cause(err))
		return
	}
	//初始化日志
	conf.Conf.Log.LogInit()
	//初始化 数据库
	err = conf.Conf.MySQL.Connect()
	if err != nil {
		conf.Conf.Log.WriteRunOutPut(err)
		return
	}
	//初始化 gin
	conf.Conf.Gin.GinInit()
	//启动
	conf.Conf.Gin.GinRun()
}
