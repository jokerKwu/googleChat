package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	swaggerDocs "main/docs"
	paymentChatHandler "main/googleChat/payment/handler"
)

func main() {
	e := echo.New()
	swaggerDocs.SwaggerInfo.Host = "localhost:3000"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	gApiV01 := e.Group("/v0.1")
	paymentChatHandler.RegisterGoogleChatHandler(gApiV01)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":3000"))
	return
}
