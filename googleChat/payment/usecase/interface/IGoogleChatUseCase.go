package _interface

import "main/googleChat/payment/model/request"

type ICreateMessageUseCase interface {
	CreateMessage(reqData *request.ReqCreateMessage) error
}
