package importHandle

import (
	"DigitalLibrary/models/importModel"
	"DigitalLibrary/types"
	_import "DigitalLibrary/types/import"
	"encoding/json"
	"github.com/emicklei/go-restful"
	"io/ioutil"
	"net/http"
)

type Import struct {
}

func (e *Import) SingleImport(req *restful.Request, rsp *restful.Response) {
	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	importData := new(_import.SingleImportReq)
	reqData, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqData, importData)
	if err != nil {
		panic(err)
	}
	err = importModel.SingleImport(importData)
	if err != nil {
		panic(err)
	}
	types.RspSucRestData(rsp, "", "ok")

}

func (e *Import) FileImport(rsp http.ResponseWriter, req *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			types.ResponseFailHttpData(rsp, e.(error).Error())
		}
	}()
	rsp.Header().Set("Content-Type", "application/json")
	rsp.Header().Add("Access-Control-Allow-Origin", "*")
	rsp.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	rsp.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type,Auth,Authorization")
	rsp.Header().Add("Access-Control-Allow-Credentials", "true")
	file, _, err := req.FormFile("file")
	if err != nil {
		panic(err)
	}
	err = importModel.FileImport(file)
	if err != nil {
		panic(err)
	}
	types.ResponseSuccessHttpData(rsp, "ok")

}
