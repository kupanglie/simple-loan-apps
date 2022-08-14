package helper

import (
	"strconv"

	"github.com/kupanglie/simple-loan-apps/services/loan/internal/constant"
)

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type commonResponse struct {
	Data  interface{}    `json:"data"`
	Error *errorResponse `json:"error"`
}

func errorGenerator(code int) *errorResponse {
	if code == 0 {
		return nil
	}
	return &errorResponse{
		Code:    strconv.Itoa(code),
		Message: constant.ERROR_MAPPING[code],
	}
}

func ResponseGenerator(data interface{}, code int) *commonResponse {
	return &commonResponse{
		Data:  data,
		Error: errorGenerator(code),
	}
}
