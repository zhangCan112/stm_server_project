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

// Stack 返回当前错误的堆栈数据
func (err *Error) Stack() *StackTrace {
	return err.stack
}

// StackInfo 返回当前错误的堆栈信息
func (err *Error) StackInfo() string {
	return fmt.Sprintf("[stackInfo]: %s %s:%d", err.stack.FuncName(), err.stack.File(), err.stack.Line())
}

func (err *Error) Error() string {
	return err.message
}
