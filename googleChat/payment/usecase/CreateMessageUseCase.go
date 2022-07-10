package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/googleChat/payment/model/domain"
	"main/googleChat/payment/model/request"
	_interface "main/googleChat/payment/usecase/interface"
	"net/http"
)

type CreateMessageUseCase struct {
}

func NewCreateMessageUseCase() _interface.ICreateMessageUseCase {
	return &CreateMessageUseCase{}
}

func (c *CreateMessageUseCase) CreateMessage(reqData *request.ReqCreateMessage) error {

	webhookURL := "https://chat.googleapis.com/v1/spaces/AAAAfp69B9E/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=5EIAjyug0v-plihlE-wZdcGrpNyyhJ8nMLRdmK84g_w%3D"
	data := domain.GoogleChat{Text: reqData.Message}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, webhookURL, &buf)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		fmt.Println(req)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
	return nil
}
