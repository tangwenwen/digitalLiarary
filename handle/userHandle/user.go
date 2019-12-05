package userHandle

import (
	"DigitalLibrary/models/userModel"
	"DigitalLibrary/types"
	"encoding/json"
	"github.com/emicklei/go-restful"
	"io/ioutil"
)

type User struct {
}

func (e *User) Login(req *restful.Request, rsp *restful.Response) {
	defer func() {
		if e := recover(); e != nil {
			rsp.Header().Add("Content-Type", "application/json")
			types.ResponseFailHttpData(rsp, e.(error).Error())
		}
	}()

	type UserCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	reqData, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		panic("Login")
	}
	user := new(UserCredentials)
	if err := json.Unmarshal(reqData, user); err != nil {
		panic("Login")
	}
	data, err := userModel.Login(user.Username, user.Password)
	if err != nil {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseFailHttpData(rsp, err.Error())
	} else {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseSuccessHttpData(rsp, data)
	}
}

func (e *User) Logout(req *restful.Request, rsp *restful.Response) {
	userToken := req.Request.Header.Get("Authorization")

	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	err := userModel.Logout(userToken)
	if err != nil {
		panic(err)
	} else {
		types.RspSucRestData(rsp, "", "ok")
	}
}
