package types

import (
	"encoding/json"
	"github.com/emicklei/go-restful"
	"net/http"
)

type Response struct {
	Code       int         `json:"code"` // 0 成功， 1失败
	Data       interface{} `json:"data"`
	Error      string      `json:"error"`
}


//NewResponse 普通返回
func NewResponse(code int, errorstr string, data ...interface{}) *Response {
	res := new(Response)
	if len(data) > 0 {
		res.Data = data[0]
	}
	res.Code = code
	res.Error = errorstr
	return res
}

//RspSucRestData 成功返回
func RspSucRestData(rsp *restful.Response, errMsg string, data interface{}) {
	response := NewResponse(0, errMsg, data)
	rsp.WriteAsJson(response)
}

// RspFailRestData 失败返回 数据
func RspFailRestData(rsp *restful.Response, errMsg string) {
	response := NewResponse(1, errMsg)
	rsp.WriteAsJson(response)
}


func ResponseSuccessHttpData(rsp http.ResponseWriter, data interface{}) {
	response, _ := NewResponse(0, "", data).Pack()
	rsp.Write(response)
}
func ResponseFailHttpData(rsp http.ResponseWriter, errMsg string) {
	response, _ := NewResponse(1, errMsg).Pack()
	rsp.Write(response)
}
func (res *Response) Pack() (data []byte, err error) {
	data, err = json.Marshal(res)
	return
}