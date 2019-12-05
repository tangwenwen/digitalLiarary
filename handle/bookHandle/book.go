package bookHandle

import (
	"DigitalLibrary/models/bookModel"
	"DigitalLibrary/types"
	bookTypes "DigitalLibrary/types/book"
	"encoding/json"
	"errors"
	"github.com/emicklei/go-restful"
	"io/ioutil"
	"strconv"
)

type Book struct {
}

func (e *Book) DelBook(req *restful.Request, rsp *restful.Response) {
	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	bookId := req.Request.URL.Query().Get("book_id")
	bookIdInt, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		panic(err)
	}
	if bookIdInt < 0 {
		panic(errors.New("illegal bookId"))
	}
	err = bookModel.DelBook(bookIdInt)
	if err != nil {
		panic(err)
	}
	types.RspSucRestData(rsp, "", "ok")

}

func (e *Book) UpdateBook(req *restful.Request, rsp *restful.Response) {
	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	dataReq := new(bookTypes.UpdateBookInfoReq)
	rawData, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(rawData, dataReq)
	if err != nil {
		panic(err)
	}
	err = bookModel.UpdateBook(dataReq)
	if err != nil {
		panic(err)
	}
	types.RspSucRestData(rsp, "", "ok")
}
