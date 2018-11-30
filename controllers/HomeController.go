package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
	"github.com/superwen/IrisMvcDemo/services"
	"time"
)

type HomeController struct{
	Ctx iris.Context
	Service services.CommonService
}

type PageInfo struct {
	Title string
	Keywords string
}

/*
 * 直接返回html
 */
func (c *HomeController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

/**
 * 直接返回文字
 */
func (c *HomeController) GetPing() string {
	return "pong"
}

/**
 * 测试redis
 */
func (c *HomeController) GetRedis() string {
	redisClient := c.Service.RedisPool.Get()
	defer redisClient.Close()

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	//redisClient.Do("GET", "test", &nowTime)
	return nowTime
}

/**
 * 返回json对象
 */
func (c *HomeController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

/**
 * 渲染视图后返回
 * 指定不同的layout
 */
func (c *HomeController) GetUserBy(username string) mvc.Result {
	return mvc.View{
		Name: "user/username.html",
		Data: PageInfo{"测试页面", "测试,Iris"},
		Layout: iris.NoLayout,
		//Layout: "layou2.html",
	}
}
