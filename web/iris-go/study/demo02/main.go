package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	//viewEngine := iris.HTML("./", ".html")
	viewEngine := iris.HTML("/Users/alex.zhou/go/src/goAction/web/iris-go/study/demo02", ".html")
	app.RegisterView(viewEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello Index")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "模板页面")
		ctx.ViewData("Content", "模板页面的内容")
		ctx.ViewData("Test", "Test22")
		ctx.View("hello.html")
	})

	addr := ":8080"
	app.Run(iris.Addr(addr), iris.WithCharset("UTF-8"))
}
