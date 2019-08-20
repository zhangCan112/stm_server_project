package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(UserAuth))
}

/// AuthType 登录验证方式
type AuthType int

const (
	USERNAME AuthType = iota
	EMAIL
)

/// UserAuth 用户登录信息
type UserAuth struct {
	Id            int       `orm:"auto"`
	UserId        int       `description:"用户(User)id"`
	IdentityType  AuthType  `description:"登录类型"`
	Identifier    string    `orm:"unique" description:"注册标识"`
	Credential    string    `description:"密码凭证"`
	LastLoginTime time.Time `orm:"null;type(datetime)" description:"最后一次登录时间"`
	Times
}

func AddUserAuth(ua UserAuth) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&ua)
	if err == nil {
		fmt.Println(id)
	}
	return id, err
}

func GetUserAuth(uaid string) (ua *UserAuth, err error) {
	o := orm.NewOrm()
	id, err := strconv.Atoi(uaid)
	if err != nil {
		return nil, err
	}
	userAuth := UserAuth{Id: id}
	err = o.Read(&userAuth)

	return &userAuth, err
}

func UpdateUserAuth(uaid string, updatefunc func(ua *UserAuth) *UserAuth) (a *UserAuth, err error) {
	o := orm.NewOrm()
	userAuth, err := GetUserAuth(uaid)
	if err != nil {
		return userAuth, err
	}

	userAuth = updatefunc(userAuth)
	if _, err := o.Update(&userAuth); err != nil {
		return nil, err
	}

	return userAuth, err
}

func Login(identityType AuthType, identifier, credential string) (success bool, userId int) {
	o := orm.NewOrm()
	var userAuth UserAuth
	err := o.QueryTable(&UserAuth{}).Filter("identityType", identityType).Filter("identifier", identifier).Filter("credential", credential).One(&userAuth, "user_id")

	if err == nil {
		return true, userAuth.UserId
	}

	return false, -1
}

func DeleteUserAuth(uaid string) error {
	o := orm.NewOrm()
	id, err := strconv.Atoi(uaid)
	if err != nil {
		return err
	}
	if _, err := o.Delete(&UserAuth{Id: id}); err != nil {
		return err
	}
	return nil
}
