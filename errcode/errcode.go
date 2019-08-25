package errcode

type ErrCode struct {
	Code int
	Desp string
}

// 统一的成功码
var (
	Successcode = &ErrCode{Code: 0, Desp: "请求成功"}
)

// 各业务模块错误码的基础偏移码
const (
	// user
	user = 1010000
)
