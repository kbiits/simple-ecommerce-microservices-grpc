package utils

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

var codeMaps map[codes.Code]int = map[codes.Code]int{
	codes.OK: http.StatusOK,

	codes.AlreadyExists:   http.StatusBadRequest,
	codes.Canceled:        http.StatusBadRequest,
	codes.InvalidArgument: http.StatusBadRequest,

	codes.Unauthenticated:  http.StatusUnauthorized,
	codes.PermissionDenied: http.StatusForbidden,
	codes.NotFound:         http.StatusNotFound,

	codes.Unavailable: http.StatusServiceUnavailable,
	codes.Internal:    http.StatusInternalServerError,
}

func GrpcCodeToHttpStatus(code codes.Code) int {
	if v, ok := codeMaps[code]; !ok {
		return codeMaps[codes.Internal]
	} else {
		return v
	}
}
