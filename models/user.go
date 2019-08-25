package models

import (
	"strconv"
	"time"

	"github.com/zhangCan112/stm_server_project/errors"

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

	if err != nil {
		return -1, errors.WrapError("数据库用户数据插入操作失败！", err)
	}

	return id, nil
}

func GetUser(uid string) (u *User, err error) {
	o := orm.NewOrm()
	id, err := strconv.Atoi(uid)
	if err != nil {
		return nil, errors.WrapError("用户Id转为数字类型失败！", err)
	}
	user := User{Id: id}
	err = o.Read(&user)

	return &user, errors.WrapError("数据库读取用户数据失败！", err)
}

func UpdateUser(uid string, updatefunc func(u *User) *User) (a *User, err error) {
	o := orm.NewOrm()
	user, err := GetUser(uid)
	if err != nil {
		return user, errors.WrapError("更新用户数据时读取用户数据失败！", err)
	}

	user = updatefunc(user)
	if _, err := o.Update(&user); err != nil {
		return nil, errors.WrapError("数据库更新用户数据操作失败！", err)
	}

	return user, nil
}

func DeleteUser(uid string) error {
	o := orm.NewOrm()
	id, err := strconv.Atoi(uid)
	if err != nil {
		return errors.WrapError("用户Id转为数字类型失败！", err)
	}
	if _, err := o.Delete(&User{Id: id}); err != nil {
		return errors.WrapError("数据库删除用户数据操作失败！", err)
	}
	return nil
}
