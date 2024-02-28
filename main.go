package main

import (
	"github.com/labstack/echo/v4"
	"lessonGo/handler"
)

type DB struct {
}

func main() {
	app := echo.New()
	//var db DB
	userHandler := handler.UserHandler{}
	app.GET("/", userHandler.HandlerUserShow)
	app.Start(":8082")
}
