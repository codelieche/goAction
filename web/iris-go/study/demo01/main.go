package main

import (
	"log"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		log.Println(ctx.Request().URL)
		ctx.Write([]byte("<h1>Hello Iris Demo 01!</h1>"))

		ctx.WriteString("<br/>Hello World!<br/>")
	})

	addr := ":8080"
	app.Run(iris.Addr(addr))
}
