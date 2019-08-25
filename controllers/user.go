package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/zhangCan112/stm_server_project/models"
	"github.com/zhangCan112/stm_server_project/services"
	"github.com/zhangCan112/stm_server_project/utils"
	"github.com/zhangCan112/stm_server_project/validation"
)

//UserController Operations about Users and UserAuth
type UserController struct {
	beego.Controller
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
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var reg userReg
	json.Unmarshal(u.Ctx.Input.RequestBody, &reg)
	valid := validation.Validation{}
	b, err := valid.Valid(&reg)
	if err != nil {
		// handle error
	}
	if !b {
		println(valid.FormatErrorMessage(valid.Errors[0]))
	}

	user := models.User{
		UserName: reg.UserName,
		Email:    reg.Email,
		Password: reg.Password,
	}
	uid, err := models.AddUser(user)
	if err != nil {

	}
	u.Data["json"] = map[string]interface{}{"uid": uid}
	u.ServeJSON()
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

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {object} utils.Response
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	response := utils.NewResponse()
	if _, ok := services.Login(username, password); ok == true {
		response.SetScode(0)
		response.SetMsg("user login success")
	} else {
		response.SetScode(1)
		response.SetMsg("user not exist")
	}
	u.Data["json"] = response.ToMap()
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
