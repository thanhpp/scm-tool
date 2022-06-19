package ginutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespTemplate struct {
	RespTemplateError
	Data interface{} `json:"data"`
}

func (resp *RespTemplate) SetData(in interface{}) {
	resp.Data = in
}

type RespTemplateError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (r *RespTemplateError) SetCode(code int) {
	r.Error.Code = code
}

func (r *RespTemplateError) SetMessage(msg string) {
	r.Error.Message = msg
}

func (r *RespTemplateError) Set200OK() {
	r.Error.Code = http.StatusOK
	r.Error.Message = "OK"
}

func RespErr(c *gin.Context, httpCode int, err error, opts ...RespErrorOpt) {
	resp := new(RespTemplate)
	resp.SetCode(httpCode)

	for i := range opts {
		opts[i](resp)
	}

	if len(resp.Error.Message) == 0 {
		resp.Error.Message = err.Error()
	} else {
		if resp.Data == nil {
			resp.Data = err.Error()
		}
	}

	c.JSON(
		httpCode,
		resp,
	)
}

type respErrorOptsFn func(*RespTemplate)

type RespErrorOpt respErrorOptsFn

func WithCode(code int) RespErrorOpt {
	return func(rt *RespTemplate) {
		rt.SetCode(code)
	}
}

func WithMessage(msg string) RespErrorOpt {
	return func(rt *RespTemplate) {
		rt.SetMessage(msg)
	}
}

func WithData(in interface{}) RespErrorOpt {
	return func(rt *RespTemplate) {
		rt.SetData(in)
	}
}
