package response

import (
	"github.com/gin-gonic/gin"
	"m4-im/pkg/e"
	"net/http"
)

type Rsp struct {
	C    *gin.Context
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *Rsp) SetCode(code int) *Rsp {
	r.Code = code
	return r
}
func (r *Rsp) SetMsg(msg string) *Rsp {
	r.Msg = msg
	return r
}
func (r *Rsp) SetData(data interface{}) *Rsp {
	r.Data = data
	return r
}

func (r *Rsp) OK() {
	r.Response(http.StatusOK)
	return
}

// Rsp setting gin.JSON
func (r *Rsp) Response(httpCode int) {
	responseMap := make(map[string]interface{})
	responseMap["code"] = r.Code
	// get default error message
	if r.Msg == "" {
		responseMap["msg"] = e.GetMsg(r.Code)
	}
	responseMap["data"] = r.Data
	r.C.JSON(httpCode, responseMap)
	return
}

func NewResponseBuilder(ctx *gin.Context) *Rsp {
	return &Rsp{
		C:    ctx,
		Data: nil,
		Code: e.SUCCESS,
	}
}

func InValidParamResponse() {

}
