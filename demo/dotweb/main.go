package main

import (
	"github.com/devfeel/dotweb"
)

func main() {
	StartServer()
}


func StartServer() error {
	//init DotApp
	app := dotweb.New()
	//set route
	app.HttpServer.GET("/index", func(ctx dotweb.Context) error{
		return ctx.WriteString("welcome to my first web!")
	})
	//begin server
	err := app.StartServer(3000)
	return err
}