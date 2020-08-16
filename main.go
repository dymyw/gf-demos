package main

import (
	// 先执行 boot 里面的 init 方法
	_ "github.com/gogf/gf-demos/boot"

	_ "github.com/gogf/gf-demos/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
