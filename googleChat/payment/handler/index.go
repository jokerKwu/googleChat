package handler

import "github.com/labstack/echo/v4"

func RegisterGoogleChatHandler(e *echo.Group) {
	handler := NewGoogleChatHandler()
	api := e.Group("/chat/payment")
	api.POST("/send", handler.CreateMessageHandler.CreateMessage)
}
