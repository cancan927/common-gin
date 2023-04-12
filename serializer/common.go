package serializer

import (
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func Success(msg string, data any) Response {
	return Response{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	}
}

func Fail(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}
}
