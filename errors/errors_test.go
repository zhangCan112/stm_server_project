package errors

import (
	"fmt"
	"testing"

	"github.com/zhangCan112/stm_server_project/testTool"
)

func TestNewError(t *testing.T) {
	err := NewError("test")

	if e := testTool.ShouldNil(err.Cause()); e != nil {
		t.Error(e)
	}

	if e := testTool.Expect("test", err.Error()); e != nil {
		t.Error(e)
	}
}

func TestStackTrace(t *testing.T) {
	err := NewError("test")
	if err.StackTrace() == nil {
		t.Errorf("TestStack.Error() err.Stack() should not be nil!")
	}
}

func TestErrorStackInfo(t *testing.T) {
	err := NewError("test")
	if err.StackInfo() == "" {
		t.Errorf("StackInfo.Error() err.StackInfo() should not be \"\"!")
	}
}

func TestWrapError(t *testing.T) {
	err1 := fmt.Errorf("error1")
	err2 := WrapError("error2", err1)

	if e := testTool.Expect(err1, err2.Cause()); e != nil {
		t.Error(e)
	}
}

func TestRootCause(t *testing.T) {
	rootCause := fmt.Errorf("error1")
	err2 := WrapError("error2", rootCause)
	err3 := WrapError("error3", err2)

	if e := testTool.Expect(rootCause, RootCause(err3)); e != nil {
		t.Error(e)
	}
}

func TestStackInfo(t *testing.T) {
	err := NewError("test")
	if e := testTool.Expect(err.StackInfo(), StackInfo(err)); e != nil {
		t.Error(e)
	}

	err2 := fmt.Errorf("error1")
	if e := testTool.Expect("", StackInfo(err2)); e != nil {
		t.Error(e)
	}
}

func TestFullFormat(t *testing.T) {
	rootCause := fmt.Errorf("error1")
	err2 := WrapError("error2", rootCause)
	err3 := WrapError("error3", err2)

	target := ""
	target += fmt.Sprintf("[Error]%s: %s\n", err3.Error(), err3.StackInfo())
	target += fmt.Sprintf("[Error]%s: %s\n", err2.Error(), err2.StackInfo())
	target += fmt.Sprintf("[Error]%s\n", rootCause.Error())

	if e := testTool.Expect(target, FullFormat(err3)); e != nil {
		t.Error(e)
	}
}
