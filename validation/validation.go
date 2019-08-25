package validation

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/validation"
)

func init() {
	validation.SetDefaultMessage(messageTmpls)
}

const (
	structLabelTag = "label"
)

// Validation 自定义的Validation
type Validation struct {
	validation.Validation
	labelMap map[string]string
}

// Valid Validate a struct.
// the obj parameter must be a struct or a struct pointer
func (v *Validation) Valid(obj interface{}) (b bool, err error) {
	b, err = v.Validation.Valid(obj)

	//验证过程没有抛错，且验证结果没有pass
	if err == nil && !b {
		objT := reflect.TypeOf(obj).Elem()
		for i := 0; i < objT.NumField(); i++ {
			field := objT.Field(i)
			name := field.Name
			label := field.Tag.Get(structLabelTag)
			if len(label) > 0 {
				v.setLabelMap(name, label)
			} else {
				v.setLabelMap(name, name)
			}
		}
	}

	return b, err
}

func (v *Validation) setLabelMap(field, label string) {
	if v.labelMap == nil {
		v.labelMap = make(map[string]string)
	}
	if len(field) > 0 && len(label) > 0 {
		v.labelMap[field] = label
	}
}

// GetLabel return StructTag:"label"'s value with validation.Error's Field
func (v *Validation) GetLabel(field string) string {
	return v.labelMap[field]
}

// FormatErrorMessage 格式化后的错误信息
func (v *Validation) FormatErrorMessage(err *validation.Error) string {
	return fmt.Sprintf("%s%s", v.GetLabel(err.Field), err.Message)
}

// FirstErrorMessage 返回验证所有错误数组的第一条错误的格式化错误信息
func (v *Validation) FirstErrorMessage() (msg string, hasErr bool) {
	hasErr = v.HasErrors()
	msg = ""
	if hasErr && len(v.Errors) > 0 {
		msg = v.FormatErrorMessage(v.Errors[0])
	}

	return msg, hasErr
}

var messageTmpls = map[string]string{
	"Required":     "不能为空",
	"Min":          "最小为 %d",
	"Max":          "最大为 %d",
	"Range":        "范围在 %d 至 %d",
	"MinSize":      "最小长度为 %d",
	"MaxSize":      "最大长度为 %d",
	"Length":       "长度必须是 %d",
	"Alpha":        "必须是有效的字母字符",
	"Numeric":      "必须是有效的数字字符",
	"AlphaNumeric": "必须是有效的字母或数字字符",
	"Match":        "必须匹配格式 %s",
	"NoMatch":      "必须不匹配格式 %s",
	"AlphaDash":    "必须是有效的字母或数字或破折号(-_)字符",
	"Email":        "必须是有效的邮件地址",
	"IP":           "必须是有效的IP地址",
	"Base64":       "必须是有效的base64字符",
	"Mobile":       "必须是有效手机号码",
	"Tel":          "必须是有效电话号码",
	"Phone":        "必须是有效的电话号码或者手机号码",
	"ZipCode":      "必须是有效的邮政编码",
}
