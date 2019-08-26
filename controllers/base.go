package controllers

import (
	"github.com/astaxie/beego"
	md "github.com/zhangCan112/stm_server_project/models"
)

type BaseController struct {
	beego.Controller
	User    md.User
	IsLogin bool
}

// Prepare implemented Prepare method for baseRouter.
func (b *BaseController) Prepare() {
	user := b.GetSession("User")

	isLogin := ""
	if user != nil {
		b.User = user.(md.User)
		b.IsLogin = true
		isLogin = "true"
	} else {
		b.IsLogin = false
		isLogin = "false"
	}

	//通过isLogin的状态告诉前端当前是否为登录状态
	if b.Ctx.GetCookie("isLogin") != isLogin {
		b.Ctx.SetCookie("isLogin", isLogin)
	}

}
