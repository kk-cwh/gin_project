package this

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode
	InvalidRequestErrorCode ResponseCode = 401

)

type Result struct {
	ErrorCode ResponseCode `json:"code"`
	ErrorMsg  string       `json:"msg"`
	Data      interface{}  `json:"data"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	resp := &Result{ErrorCode: code, ErrorMsg: err.Error(), Data: ""}
	c.JSON(200, resp)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Result{ErrorCode: SuccessCode, ErrorMsg: "", Data: data}
	c.JSON(200, resp)
}

func Response(c *gin.Context, err error, data interface{}) {
	if err!=nil {
		ResponseError(c,http.StatusInternalServerError,err)
	}else {
		ResponseSuccess(c, data)
	}
}
