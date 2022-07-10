package handler

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"main/googleChat/delivery/model/request"
	"main/googleChat/delivery/usecase"
	_interface "main/googleChat/delivery/usecase/interface"
	"net/http"
)

type CreateMessageHandler struct {
	UseCase _interface.ICreateMessageUseCase
}

func NewCreateMessageHandler() *CreateMessageHandler {
	return &CreateMessageHandler{UseCase: usecase.NewCreateMessageUseCase()}
}

var val = validator.New()

// payment google chat send
// @Router /v0.1/chat/payment/send [post]
// @Summary 구글 챗 메시지 전송
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description
// @Description ■ errCode with 401
// @Description TOKEN_BAD : 토큰 인증 실패
// @Description POLICY_VIOLATION : 토큰 세션 정책 위반
// @Param json body request.ReqCreateMessage true "json body"
// @Produce json
// @Success 200 {object} domain.ResJustOk
// @Failure 400 {object} domain.ResError
// @Failure 401 {object} domain.ResError
// @Failure 500 {object} domain.ResError
// @Tags google/chat
func (m *CreateMessageHandler) CreateMessage(c echo.Context) error {
	req := request.ReqCreateMessage{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := val.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := m.UseCase.CreateMessage(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "ok")
}
