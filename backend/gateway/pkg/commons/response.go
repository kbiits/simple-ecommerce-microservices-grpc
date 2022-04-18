package commons

import (
	"net/http"
)

type OkResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
	Code   int         `json:"code"`
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func NewBadRequestResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Error:  err.Error(),
		Status: StatusBadRequest,
		Code:   http.StatusBadRequest,
	}
}
