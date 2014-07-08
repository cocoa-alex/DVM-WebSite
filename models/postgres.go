package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DR_Postgres)

	orm.RegisterDataBase("default", postgres, dataSource, ...)
}
