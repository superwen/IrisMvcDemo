package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"fmt"
	"github.com/kataras/iris/mvc"
	"github.com/superwen/IrisMvcDemo/controllers"
	"github.com/superwen/IrisMvcDemo/services"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.RegisterView(iris.HTML("./templates", ".html"))
	//app.RegisterView(iris.HTML("./templates", ".html").Layout("layout.html"))
	// 指定静态目录
	app.StaticWeb("/public", "./web/public")
	// 指定公用错误返回页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})
	//初始化service
	commonService := services.NewCommonService()

	homeMvc := mvc.New(app.Party("/"))
	//users.Router.Use(middleware.BasicAuth)
	homeMvc.Register(commonService)
	homeMvc.Handle(new(controllers.HomeController))


	app.Run(
		iris.Addr(fmt.Sprintf("%s:%s", "", "8099")),
		iris.WithConfiguration(iris.YAML("./iris.yml")),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
