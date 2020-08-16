package main

import (
	// 先执行 boot 里面的 init 方法
	_ "github.com/gogf/gf-demos/boot"

	// 执行 router 里面的 init 方法
	_ "github.com/gogf/gf-demos/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	// 以阻塞方式启动服务器侦听，通常用于单服务器情况
	g.Server().Run()
}
