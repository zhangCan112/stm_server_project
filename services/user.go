package services

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/zhangCan112/stm_server_project/errors"
	md "github.com/zhangCan112/stm_server_project/models"
)

// UserIsExisted 用户是否已存在
func UserIsExisted(userName, email string) bool {
	o := orm.NewOrm()
	var (
		user md.User
	)
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("user_name", userName).Or("email", email)
	qs := o.QueryTable(&user).SetCond(cond)
	if err := qs.One(&user); err != nil {
		return false
	}
	return true
}

// Login 登录账户
func Login(identifier, password string) (u *md.User, ok bool) {
	o := orm.NewOrm()
	var (
		user md.User
		err  error
	)
	ok = false
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("user_name", identifier).Or("email", identifier)
	qs := o.QueryTable(&user).SetCond(cond)
	if err = qs.One(&user); err == nil {
		if user.Password == password {
			user.LastLoginTime = time.Now()
			md.UpdateUser(&user)
			ok = true
		}
	}
	return &user, ok
}

//Reg 用户注册
func Reg(userName, email, password string) (err error) {
	_, err = md.AddUser(md.User{
		UserName: userName,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return errors.WrapError("注册用户失败!", err)
	}
	return nil
}
