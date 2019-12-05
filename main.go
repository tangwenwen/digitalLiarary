package main

import (
	"DigitalLibrary/common"
	"DigitalLibrary/config"
	"DigitalLibrary/plugins/etcd"
	"DigitalLibrary/plugins/logs"
	"flag"
	"fmt"
	"github.com/micro/go-web"
)

func main() {
	port := flag.String("httpport", "7777", "http listen port")
	flag.Parse()

	service := web.NewService(
		web.Address("0.0.0.0:" + *port),
	)

	service.Init()
	service.Handle("/", config.GetRouterContainer())
	service.HandleFunc("/api/import/fileImport", config.FileImport())

	logs.Info("http has listen port on 7777")
	etcdAddr := fmt.Sprintf("%s:%d", common.EtcdHost, common.EccdPort)
	etcd.EtcdConfSet(etcdAddr)
	if err := service.Run(); err != nil {
		logs.Error(err)
	}

}
