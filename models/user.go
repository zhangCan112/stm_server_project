package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

// User 用户信息
type User struct {
	Id            int       `orm:"auto"`
	UserName      string    `orm:"unique" description:"用户名"`
	Email         string    `orm:"null;unique" description:"电子邮箱"`
	Password      string    `description:"密码"`
	LastLoginTime time.Time `orm:"null;type(datetime)" description:"最后一次登录时间"`
	Times
}

func AddUser(u User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&u)
	if err == nil {
		fmt.Println(id)
	}
	return id, err
}

func GetUser(uid string) (u *User, err error) {
	o := orm.NewOrm()
	id, err := strconv.Atoi(uid)
	if err != nil {
		return nil, err
	}
	user := User{Id: id}
	err = o.Read(&user)

	return &user, err
}

func UpdateUser(uid string, updatefunc func(u *User) *User) (a *User, err error) {
	o := orm.NewOrm()
	user, err := GetUser(uid)
	if err != nil {
		return user, err
	}

	user = updatefunc(user)
	if _, err := o.Update(&user); err != nil {
		return nil, err
	}

	return user, err
}

func DeleteUser(uid string) error {
	o := orm.NewOrm()
	id, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	if _, err := o.Delete(&User{Id: id}); err != nil {
		return err
	}
	return nil
}
