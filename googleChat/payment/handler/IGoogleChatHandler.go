package handler

import (
	_interface "main/googleChat/payment/usecase/interface"
)

type ICreateMessageHandler interface {
	CreateMessage() error
}

type IGoogleChatHandler interface {
	NewCreateMessageHandler(UseCase _interface.ICreateMessageUseCase) *CreateMessageHandler
}

type GoogleChatHandler struct {
	CreateMessageHandler CreateMessageHandler
}

func NewGoogleChatHandler() *GoogleChatHandler {
	return &GoogleChatHandler{
		CreateMessageHandler: *NewCreateMessageHandler(),
	}
}
