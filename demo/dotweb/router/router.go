package router

import (
	"github.com/devfeel/dotweb"
)

func NewRouter() *dotweb.Router{
	app := dotweb.New()
	//set route
	app.HttpServer.GET("/index", func(ctx dotweb.Context) error{
		return ctx.WriteString("welcome to my first web!")
	})
	return nil;
}

