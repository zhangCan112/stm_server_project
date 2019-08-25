package utils

import (
	"github.com/zhangCan112/stm_server_project/errcode"
)

//Response 约束的Response的基本结构
type Response struct {
	scode int
	body  map[string]interface{}
	msg   string
}

//NewResponse creator
func NewResponse() *Response {
	return &Response{}
}

//SetErrcode 用errcode.ErrCode设置Scode和msg
func (rp *Response) SetErrcode(ec *errcode.ErrCode) {
	rp.scode = ec.Code
	rp.msg = ec.Desp
}

//SetScode 设置Scode值
func (rp *Response) SetScode(s int) {
	rp.scode = s
}

//SetMsg 设置Msg值
func (rp *Response) SetMsg(s string) {
	rp.msg = s
}

//SetBody 设置Body字段
func (rp *Response) SetBody(body map[string]interface{}) {
	rp.body = body
}

//UpdateInBody 更新body中的某个字段
func (rp *Response) UpdateInBody(key string, val interface{}) {
	if rp.body == nil {
		rp.body = make(map[string]interface{})
	}
	rp.body[key] = val
}

//ToMap 转换为Map数据
func (rp *Response) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["scode"] = rp.scode
	result["body"] = rp.body
	result["msg"] = rp.msg
	return result
}
