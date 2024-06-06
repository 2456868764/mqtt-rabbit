package api

import (
	"bifromq_engine/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	OriginUrl string      `json:"originUrl"`
}

func ResponseAny(code int, data interface{}, msg string) *Response {
	return &Response{
		Code:    code,
		Data:    data,
		Message: msg,
	}
}

func ResponseOK(data interface{}) *Response {
	return &Response{
		Code:    0,
		Data:    data,
		Message: "",
	}
}

func ResponseError(code int, msg string) *Response {
	return &Response{
		Code:    code,
		Data:    nil,
		Message: msg,
	}
}

func ResponseErrorStatus(errCode *errcode.Status) *Response {
	return &Response{
		Code:    errCode.Code,
		Data:    nil,
		Message: errCode.Msg,
	}
}

func Success(c *gin.Context, data interface{}) {
	response := ResponseOK(data)
	c.JSON(http.StatusOK, response)
}

func Error(c *gin.Context, code int, msg string) {
	response := ResponseError(code, msg)
	c.JSON(http.StatusOK, response)
}

func ErrorStatus(c *gin.Context, errCode *errcode.Status) {
	response := ResponseErrorStatus(errCode)
	c.JSON(http.StatusOK, response)
}
