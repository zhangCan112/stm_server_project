package errcode

var (
	UserRegValidError     = &ErrCode{Code: user + 1, Desp: "注册表单数据验证出错!"}
	UserRegValidNotPass   = &ErrCode{Code: user + 2, Desp: "注册表单数据验证未通过!"}
	UserRegServiceError   = &ErrCode{Code: user + 3, Desp: "注册失败！"}
	UserRegUserHasExisted = &ErrCode{Code: user + 4, Desp: "注册用户已存在！"}
	UserLoginFailed       = &ErrCode{Code: user + 5, Desp: "用户名或密码错误！"}
	UserRepeatLogin       = &ErrCode{Code: user + 6, Desp: "用户已登录！"}
)
