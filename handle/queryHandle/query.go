package queryHandle

import (
	"DigitalLibrary/models/queryModel"
	"DigitalLibrary/types"
	queryTypes "DigitalLibrary/types/query"
	"encoding/json"
	"github.com/emicklei/go-restful"
	"io/ioutil"
)

type Query struct {
}

func (e *Query) GetAllInfo(req *restful.Request, rsp *restful.Response) {
	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	reqQuery := new(queryTypes.ReqQuery)
	reqData, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqData, reqQuery)
	if err != nil {
		panic(err)
	}
	data, err := queryModel.QueryAllInfo(reqQuery.QueryType, reqQuery.QueryOpt)
	if err != nil {
		panic(err)
	}
	types.RspSucRestData(rsp, "", data)
}
