package models

import "time"

// Base base
type Base struct {
	Id      int       `orm:"auto"`
	Created time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Updated time.Time `orm:"auto_now;type(datetime)" description:"最后一次更新时间"`
}
