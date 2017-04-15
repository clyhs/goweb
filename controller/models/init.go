package models

import (
	"errors"

	"github.com/clyhs/helloweb/controller/setting"
	"github.com/go-xorm/xorm"
)

var (
	ErrNotExist = errors.New("not exist")
)

var orm *xorm.Engine

func Init(isProMode bool) {
	var err error
	orm, err = xorm.NewEngine(setting.DriverName, setting.DataSource)
	if err != nil {
		panic(err)
	}

	orm.SetMaxIdleConns(setting.MaxIdle)
	orm.SetMaxOpenConns(setting.MaxOpen)
	if !isProMode {
		orm.ShowSQL(true)
	}

	err = orm.Sync2(new(User))
	if err != nil {
		panic(err)
	}

}
