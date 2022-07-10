package domain

type GoogleChat struct {
	Text string `json:"text"`
}

type ResError struct {
	ErrCode    string `json:"errCode"`
	Msg        string `json:"msg,omitempty"`
	Detail     string `json:"detail,omitempty"`
	StackTrace string `json:"stackTrace,omitempty"`
}

type ResJustOk struct {
	Result string `json:"result" example:"ok"` // "ok"로만 리턴됨
}

var ResJustOkData = ResJustOk{
	Result: "ok",
}
