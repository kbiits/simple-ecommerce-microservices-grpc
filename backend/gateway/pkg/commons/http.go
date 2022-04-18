package commons

var httpCodesMap map[int]string = map[int]string{
	200: "OK",
	201: "CREATED",
	202: "ACCEPTED",
	204: "NO_CONTENT",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	402: "PAYMENT_REQUIRED",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	405: "METHOD_NOT_ALLOWED",
	408: "REQUEST_TIMEOUT",
	409: "CONFLICT",
	410: "GONE",
	412: "PRECONDITION_FAILED",
	413: "PAYLOAD_TO_LARGE",
	415: "UNSUPPORTED_MEDIA_TYPE",
	422: "UNPROCESSABLE_ENTITY",
	423: "LOCKED",
	429: "TOO_MANY_REQUEST",
	500: "INTERNAL_SERVER_ERROR",
	501: "NOT_IMPLEMENTED",
	502: "BAD_GATEWAY",
	503: "SERVICE_UNAVAILABLE",
	504: "GATEWAY_TIMEOUT",
}

func HttpCodeToString(code int) string {
	if v, ok := httpCodesMap[code]; !ok {
		return httpCodesMap[500]
	} else {
		return v
	}
}
