package boot

import (
	// 先执行 packed 里面的 init 方法
	// 【【【 这里建议引入 packed 包和其他包之间加入一个空行以作区分，特别是 Goland IDE 的 import 插件不会将引入包进行自动排序 】】】
	_ "github.com/gogf/gf-demos/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

// 用于应用初始化。
func init() {
	s := g.Server()

	// 注册插件，swagger.Swagger 实现了 ghttp.Plugin
	s.Plugin(&swagger.Swagger{})
}
