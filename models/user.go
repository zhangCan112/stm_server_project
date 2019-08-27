package models

import (
	"time"

	"github.com/zhangCan112/stm_server_project/errors"

	"github.com/astaxie/beego/orm"
	"github.com/rs/xid"
)

func init() {
	orm.RegisterModel(new(User))
}

// User 用户信息
type User struct {
	Id            int       `orm:"auto"`
	Guid          string    `orm:"unique" description:"用户唯一Id"`
	UserName      string    `orm:"unique" description:"用户名"`
	Email         string    `orm:"null;unique" description:"电子邮箱"`
	Password      string    `description:"密码"`
	LastLoginTime time.Time `orm:"null;type(datetime)" description:"最后一次登录时间"`
	Times
}

func AddUser(u User) (string, error) {
	o := orm.NewOrm()
	u.Guid = xid.New().String()
	_, err := o.Insert(&u)

	if err != nil {
		return "", errors.WrapError("数据库用户数据插入操作失败！", err)
	}

	return u.Guid, nil
}

func GetUser(guid string) (u *User, err error) {
	o := orm.NewOrm()

	user := User{Guid: guid}
	err = o.Read(&user)

	return &user, errors.WrapError("数据库读取用户数据失败！", err)
}

func UpdateUser(user *User) (a *User, err error) {
	o := orm.NewOrm()
	if _, err := o.Update(user); err != nil {
		return nil, errors.WrapError("数据库更新用户数据操作失败！", err)
	}
	return user, nil
}

func GetAndUpdateUser(guid string, updatefunc func(u *User) *User) (a *User, err error) {
	o := orm.NewOrm()
	user, err := GetUser(guid)
	if err != nil {
		return user, errors.WrapError("获取并更新用户数据时读取用户数据失败！", err)
	}

	user = updatefunc(user)
	if _, err := o.Update(&user); err != nil {
		return nil, errors.WrapError("获取并更新用户数据时更新操作失败！", err)
	}

	return user, nil
}

func DeleteUser(guid string) error {
	o := orm.NewOrm()
	if _, err := o.Delete(&User{Guid: guid}); err != nil {
		return errors.WrapError("数据库删除用户数据操作失败！", err)
	}
	return nil
}
