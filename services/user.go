package services

import (	
	"github.com/astaxie/beego/orm"
	md "github.com/zhangCan112/stm_server_project/models"
)

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
			ok = true
		}
	}
	return &user, ok
}

//Reg 用户注册
func Reg(userName, email, password string) error {
	_, err := md.AddUser(md.User{
		UserName: userName,
		Email:    email,
		Password: password,
	})
	if err != nil {		
		return err
	}
	return nil
}
