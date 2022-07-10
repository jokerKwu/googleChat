package _interface

import "main/googleChat/delivery/model/request"

type ICreateMessageUseCase interface {
	CreateMessage(reqData *request.ReqCreateMessage) error
}
