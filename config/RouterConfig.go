package config

import (
	"DigitalLibrary/handle/bookHandle"
	"DigitalLibrary/handle/importHandle"
	"DigitalLibrary/handle/queryHandle"
	"DigitalLibrary/handle/userHandle"
	"DigitalLibrary/plugins/auth"
	"github.com/emicklei/go-restful"
	"net/http"
)

var user = new(userHandle.User)
var importB = new(importHandle.Import)
var query = new(queryHandle.Query)
var book = new(bookHandle.Book)

const (
	PATH = "/api"
)

// restful容器初始化
var wc = restful.NewContainer()

// restful 服务初始化
var ws = new(restful.WebService)

func routerConf() {
	usersRouterConf()
	importRouterConf()
	queryRouterConf()
	bookRouterConf()
}

func GetRouterContainer() *restful.Container {

	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path(PATH)
	routerConf()
	wc.Add(ws)
	return wc
}

func usersRouterConf() {
	rootPath := "users/"
	methodPath := "login"
	ws.Route(ws.POST(rootPath + methodPath).To(user.Login))

	methodPath = "logout"
	ws.Route(ws.GET(rootPath + methodPath).To(user.Logout).Filter(auth.TokenFilter))
}

func importRouterConf() {
	rootPath := "import/"
	methodPath := "singleImport"
	ws.Route(ws.POST(rootPath + methodPath).To(importB.SingleImport).Filter(auth.TokenFilter))

}

func queryRouterConf() {
	rootPath := "query/"
	methodPath := "getAllInfo"
	ws.Route(ws.POST(rootPath + methodPath).To(query.GetAllInfo))
}

func bookRouterConf() {
	rootPath := "book/"
	methodPath := "delBook"
	ws.Route(ws.GET(rootPath + methodPath).To(book.DelBook).Filter(auth.TokenFilter))
	methodPath = "updateBook"
	ws.Route(ws.POST(rootPath + methodPath).To(book.UpdateBook).Filter(auth.TokenFilter))
}

func FileImport() func(http.ResponseWriter, *http.Request) {
	return importB.FileImport
}
