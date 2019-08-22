package errors

import (
	"runtime"
	"strings"
)

//StackTrace 堆栈估计
type StackTrace struct {
	pc   uintptr
	fn   string
	file string
	line int
}

//Caller return a StackTrace point
func Caller(skip int) *StackTrace {
	if pc, file, line, ok := runtime.Caller(skip + 1); ok == true {
		return &StackTrace{
			pc:   pc,
			file: file,
			line: line,
		}
	}
	return nil
}

//File return file path
func (st *StackTrace) File() string {
	return st.file
}

//Line return line No.
func (st *StackTrace) Line() int {
	return st.line
}

//FuncFullName return function full path Name
func (st *StackTrace) FuncFullName() string {
	f := runtime.FuncForPC(st.pc)
	if f == nil {
		return "unknown"
	}
	return f.Name()
}

//FuncName return function short Name
func (st *StackTrace) FuncName() string {
	name := st.FuncFullName()
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
