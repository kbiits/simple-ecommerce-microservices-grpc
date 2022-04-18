package utils

import (
	"net/http"

	"github.com/kbiits/microservices-grpc-simple-ecommerce-gateway/pkg/commons"
	"google.golang.org/grpc/status"
)

func ErrorResponseFromGrpc(err error) *commons.ErrorResponse {
	st, ok := status.FromError(err)
	// error is not grpc status
	if !ok {
		return &commons.ErrorResponse{
			Error:  err.Error(),
			Status: commons.StatusInternalServerError,
			Code:   http.StatusInternalServerError,
		}
	}

	return &commons.ErrorResponse{
		Error:  st.Message(),
		Status: ToSnakeCaseUpper(st.Code().String()),
		Code:   GrpcCodeToHttpStatus(st.Code()),
	}
}

func UnauthorizedResponse(msg string) *commons.ErrorResponse {
	return &commons.ErrorResponse{
		Error:  msg,
		Status: commons.StatusUnauthorized,
		Code:   http.StatusUnauthorized,
	}
}

func ErrorResponseFromMessage(msg string, status uint32) *commons.ErrorResponse {
	return &commons.ErrorResponse{
		Error:  msg,
		Status: commons.HttpCodeToString(int(status)),
		Code:   int(status),
	}
}

func ToOkResponse(code uint32, data interface{}) *commons.OkResponse {
	return &commons.OkResponse{
		Data:   data,
		Status: commons.HttpCodeToString(int(code)),
		Code:   int(code),
	}
}
