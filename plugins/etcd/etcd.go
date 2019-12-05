package etcd

import (
	"DigitalLibrary/plugins/logs"
	"context"
	"errors"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
)

const (
	ETCDROOT = "/tww/tokens"
	Timout   = 5 * time.Second
)

type enum_KeyStatus uint8

const (
	KeyInDb    enum_KeyStatus = 0
	KeyNotInDb enum_KeyStatus = 1
)

var (
	config *clientv3.Config
)

func EtcdConfSet(ipports string) {
	config = new(clientv3.Config)
	end := strings.SplitN(ipports, ",", -1)
	config.Endpoints = end
	config.DialTimeout = Timout
}

func EtcdPut(key, value string) error {
	cli, err := clientv3.New(*config)
	if err != nil {
		logs.Error("", err)
		return err
	}
	if cli == nil {
		err := errors.New("cli == nil")
		logs.Error("", err)
		return err
	}

	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), Timout)
	_, err = cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		logs.Error("", err)
		return err
	} else {
		return nil
	}
}

func EtcdGet(key string) (string, enum_KeyStatus, error) {
	cli, err := clientv3.New(*config)
	if err != nil {
		return "", KeyNotInDb, err
	}

	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), Timout)
	rest, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		return "", KeyNotInDb, err
	}
	if len(rest.Kvs) != 1 {
		return "", KeyNotInDb, nil
	}

	jsb := rest.Kvs[0].Value
	jss := string(jsb)
	return jss, KeyInDb, nil
}

func EtcdPutLease(key, value string, s int64) error {
	cli, err := clientv3.New(*config)
	if err != nil {
		logs.Error("", err)
		return err
	}
	resp, err := cli.Grant(context.TODO(), s)
	if err != nil {
		logs.Error("", err)
		return err
	}

	defer cli.Close()
	_, err = cli.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err != nil {
		logs.Error("", err)
		return err
	}
	return err
}

func EtcdDel(key string) (enum_KeyStatus, error) {
	_, keystatus, err := EtcdGet(key)
	if err != nil {
		return KeyNotInDb, err
	}
	if keystatus == KeyNotInDb {
		return KeyNotInDb, nil
	}
	cli, err := clientv3.New(*config)
	if err != nil {
		logs.Error("", err)
		return KeyNotInDb, err
	}

	_, err = cli.Delete(context.TODO(), key, clientv3.WithPrefix())
	if err != nil {
		return KeyNotInDb, err
	}
	return KeyInDb, nil
}
