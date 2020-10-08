package main

import (
	"goFlow/model"
	"goFlow/routers"
)

func main() {
	//引用数据库
	model.InitDb()
	//注册路由
	routers.InitRouter()
}
