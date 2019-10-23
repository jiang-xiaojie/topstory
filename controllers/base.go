package controllers

import (
	"fmt"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Extra   interface{} `json:"extra,omitempty"`
}

func formatJSON(data interface{}, err error) interface{} {
	if err != nil {
		return &Response{
			Code:    -1,
			Message: fmt.Sprintf("%v", err),
			Data:    data,
			Extra:   nil,
		}
	}
	return &Response{
		Code:    0,
		Message: "",
		Data:    data,
		Extra:   nil,
	}
}
