package validation

import (
	"testing"

	"github.com/zhangCan112/stm_server_project/testTool"
)

type testUserReg struct {
	UserName string `json:"userName" label:"用户名" valid:"Required;AlphaDash;MaxSize(20);MinSize(3)"`
	Email    string `json:"email" label:"邮箱" valid:"Required; Email; MaxSize(100)"`
	Password string `json:"password" label:"密码" valid:"Required;MinSize(6);MaxSize(15)"`
	Test     string
}

func TestValid(t *testing.T) {
	testData := &testUserReg{
		UserName: "",
		Email:    "1212",
		Password: "",
	}

	valid := &Validation{}
	b, err := valid.Valid(testData)

	if e := testTool.ShouldNil(err); e != nil {
		t.Error(e)
	}

	if e := testTool.Expect(false, b); e != nil {
		t.Error(e)
	}

	userName := valid.labelMap["UserName"]
	email := valid.labelMap["Email"]
	password := valid.labelMap["Password"]
	test := valid.labelMap["Test"]

	if e := testTool.Expect("用户名", userName); e != nil {
		t.Error(e)
	}

	if e := testTool.Expect("邮箱", email); e != nil {
		t.Error(e)
	}

	if e := testTool.Expect("密码", password); e != nil {
		t.Error(e)
	}

	if e := testTool.Expect("Test", test); e != nil {
		t.Error(e)
	}

}

func TestGetLabel(t *testing.T) {
	testData := &testUserReg{
		UserName: "",
		Email:    "1212",
		Password: "",
	}

	valid := &Validation{}
	valid.Valid(testData)

	if e := testTool.Expect("邮箱", valid.GetLabel("Email")); e != nil {
		t.Error(e)
	}

	if e := testTool.Expect("密码", valid.GetLabel("Password")); e != nil {
		t.Error(e)
	}

	if e := testTool.Expect("Test", valid.GetLabel("Test")); e != nil {
		t.Error(e)
	}

}

func TestFirstErrorMessage(t *testing.T) {
	testData := &testUserReg{
		UserName: "",
		Email:    "1212",
		Password: "",
	}

	valid := &Validation{}
	valid.Valid(testData)

	msg, hasErr := valid.FirstErrorMessage()

	if e := testTool.Expect(true, hasErr); e != nil {
		t.Error(e)
	}

	if e := testTool.Expect("用户名不能为空", msg); e != nil {
		t.Error(e)
	}

}
