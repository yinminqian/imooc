package main

import "github.com/kataras/iris"

func main() {
	//创建iris实例
	app := iris.New()
	//设置错误模式,在MVC下提示错误
	app.Logger().SetLevel("debug")
	//注册模板
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//设置模板静态文件
	app.StaticWeb("/assets", "./backend/web/assets")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(context iris.Context) {
		context.ViewData("messages", context.Values().GetStringDefault("messages", "访问的页面出错"))
		context.ViewLayout("")
		context.View("share/error.html")
	})

	//注册控制器

	//启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
