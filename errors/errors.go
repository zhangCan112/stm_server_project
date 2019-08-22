package errors

import "fmt"

// Error 自定义的Error
type Error struct {
	message string
	cause   error
	stack   *StackTrace
}

// NewError 新建一个error, stack会记录当前执行当前方法时的堆栈信息
func NewError(msg string) *Error {
	return &Error{
		message: msg,
		stack:   Caller(1),
	}
}

// WrapError 包含指定的error生成一个新的Error
func WrapError(msg string, err error) *Error {
	return &Error{
		message: msg,
		cause:   err,
		stack:   Caller(1),
	}
}

// StackTrace 返回当前错误的堆栈数据对象
func (err *Error) StackTrace() *StackTrace {
	return err.stack
}

// StackInfo 返回当前错误的堆栈信息
func (err *Error) StackInfo() string {
	return fmt.Sprintf("[stackInfo]: %s %s:%d", err.stack.FuncName(), err.stack.File(), err.stack.Line())
}

// Cause 返回当前错误的上一级错误
func (err *Error) Cause() error {
	return err.cause
}

func (err *Error) Error() string {
	return err.message
}

// RootCause 返回error的最终根错误
// 如果error没有Cause()方法，则返回自身
func RootCause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

// StackInfo 返回error的堆栈信息
// 如果error没有实现StackInfo方法则返回为""空字符串
func StackInfo(err error) string {
	type stacker interface {
		StackInfo() string
	}

	stack, ok := err.(stacker)
	if !ok {
		return ""
	}
	return stack.StackInfo()
}

// FullFormat 返回包含每一级错误信息+堆栈信息的完整错误信息
// 如果error没有同时实现StackInfo 和 Cause 方法则结束遍历，返回当前Error信息
func FullFormat(err error) string {
	type myerror interface {
		StackInfo() string
		Cause() error
		Error() string
	}
	format := ""
	for err != nil {
		my, ok := err.(myerror)
		if !ok {
			format += fmt.Sprintf("[Error]%s\n", err.Error())
			break
		}
		format += fmt.Sprintf("[Error]%s: %s\n", my.Error(), my.StackInfo())
		err = my.Cause()
	}
	return format

}
