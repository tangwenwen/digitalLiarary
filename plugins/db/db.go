package db

import (
	"DigitalLibrary/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strconv"
)

func GetEngine() (*xorm.Engine, error) {
	dburl := common.MysqlUserName + ":" + common.MysqlPassword + "@tcp(" + common.MysqlHost + ":" + strconv.Itoa(common.MysqlPort) + ")/" + common.MysqlDBName + "?charset=utf8"
	orm, err := xorm.NewEngine("mysql", dburl)
	if err != nil {
		return nil, err
	}
	orm.ShowSQL(true)
	return orm, nil
}
