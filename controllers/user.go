package controllers

import (
	"encoding/json"

	"github.com/zhangCan112/stm_server_project/errcode"
	"github.com/zhangCan112/stm_server_project/errors"
	"github.com/zhangCan112/stm_server_project/services"
	"github.com/zhangCan112/stm_server_project/utils"
	"github.com/zhangCan112/stm_server_project/validation"
)

//UserController Operations about Users and UserAuth
type UserController struct {
	BaseController
}

// userReg Post请求的表单数据模型
type userReg struct {
	UserName string `json:"userName" label:"用户名" valid:"Required;AlphaDash;MaxSize(20);MinSize(3)"`
	Email    string `json:"email" label:"邮箱" valid:"Required; Email; MaxSize(100)"`
	Password string `json:"password" label:"密码" valid:"Required;MinSize(6);MaxSize(15)"`
}

// Post post
// @Title CreateUser
// @Description create users
// @Param	body		body controllers.userReg	true		"body for user content"
// @Success 200 {object} utils.Response
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {

	response := utils.NewResponse()
	defer func() {
		u.Data["json"] = response.ToMap()
		defer u.ServeJSON()
	}()

	var reg userReg
	json.Unmarshal(u.Ctx.Input.RequestBody, &reg)
	valid := validation.Validation{}
	b, err := valid.Valid(&reg)
	if err != nil {
		errors.WrapError(errcode.UserRegValidError.Desp, err).Log(errcode.UserRegValidError.Code)
		response.SetErrcode(errcode.UserRegValidError)
		return
	}
	if !b {
		response.SetScode(errcode.UserRegValidNotPass.Code)
		msg, _ := valid.FirstErrorMessage()
		response.SetMsg(msg)
		return
	}

	if isExist := services.UserIsExisted(reg.UserName, reg.Email); isExist {
		response.SetErrcode(errcode.UserRegUserHasExisted)
		return
	}

	err = services.Reg(reg.UserName, reg.Email, reg.Password)

	if err != nil {
		errors.WrapError(errcode.UserRegServiceError.Desp, err).Log(errcode.UserRegServiceError.Code)
		response.SetErrcode(errcode.UserRegServiceError)
		return
	}

	response.SetScode(errcode.Successcode.Code)
	response.SetMsg("注册成功!")
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	// users := models.GetAllUsers()
	// u.Data["json"] = users
	// u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	// uid := u.GetString(":uid")
	// if uid != "" {
	// 	user, err := models.GetUser(uid)
	// 	if err != nil {
	// 		u.Data["json"] = err.Error()
	// 	} else {
	// 		u.Data["json"] = user
	// 	}
	// }
	// u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	// uid := u.GetString(":uid")
	// if uid != "" {
	// 	var user models.User
	// 	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// 	uu, err := models.UpdateUser(uid, &user)
	// 	if err != nil {
	// 		u.Data["json"] = err.Error()
	// 	} else {
	// 		u.Data["json"] = uu
	// 	}
	// }
	// u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	// uid := u.GetString(":uid")
	// models.DeleteUser(uid)
	// u.Data["json"] = "delete success!"
	// u.ServeJSON()
}

// Login 登录接口
// @Title Login
// @Description (登录接口)Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {object} utils.Response
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {

	response := utils.NewResponse()
	defer func() {
		u.Data["json"] = response.ToMap()
		u.ServeJSON()
	}()

	if u.IsLogin == true {
		response.SetErrcode(errcode.UserRepeatLogin)
		return
	}

	username := u.GetString("username")
	password := u.GetString("password")
	if user, ok := services.Login(username, password); ok == true {
		u.SetSession("User", *user)
		response.SetScode(errcode.Successcode.Code)
		response.SetMsg("用户登录成功！")
		u.Ctx.SetCookie("isLogin", "true")
	} else {
		response.SetErrcode(errcode.UserLoginFailed)
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// @Title health
// @Description health keepAlive and refresh login status in user session(健康检查接口，主要用于前端更新用户的登陆态cookie)
// @Success 200 {object} utils.Response
// @router /health [get]
func (u *UserController) Health() {
	response := utils.NewResponse()
	defer func() {
		u.Data["json"] = response.ToMap()
		u.ServeJSON()
	}()
	response.SetMsg("I'm ok!")
}
