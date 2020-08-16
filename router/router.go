package router

import (
	"github.com/gogf/gf-demos/app/api/chat"
	"github.com/gogf/gf-demos/app/api/curd"
	"github.com/gogf/gf-demos/app/api/user"
	"github.com/gogf/gf-demos/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到 router 目录下。
func init() {
	// 返回一个 ghttp.Server 实例
	s := g.Server()

	// 某些浏览器直接请求 favicon.ico 文件，特别是产生 404 时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")
	// SetRewrite 可以为一些静态资源重写路由

	// 分组路由注册方式，可以给分组路由指定一个 prefix 前缀
	// https://goframe.org/net/ghttp/group/index
	s.Group("/", func(group *ghttp.RouterGroup) {
		// func new(Type) *Type 返回一个执行该类型的指针
		ctlChat := new(chat.Controller)
		ctlUser := new(user.Controller)

		// 给路由分组绑定一个或多个中间件
		group.Middleware(middleware.CORS)

		// ALL 代表允许所有的请求方式，如 GET、POST 等都可以
		group.ALL("/chat", ctlChat)
		group.ALL("/user", ctlUser)
		group.ALL("/curd/:table", new(curd.Controller))

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth)
			group.ALL("/user/profile", ctlUser, "Profile")
		})
	})
}
