package testTool

import (
	"fmt"
	"reflect"
)

/* 测试辅助函数 */
func Expect(a, b interface{}) error {
	if a != b {
		return fmt.Errorf(fmt.Sprintf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a)))
	}
	return nil
}

func Refute(a, b interface{}) error {
	if a == b {
		return fmt.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
	return nil
}

func NotNil(a interface{}) error {
	if a == nil {
		return fmt.Errorf("Should not be nil")
	}
	return nil
}

func ShouldNil(a interface{}) error {
	if a != nil {
		return fmt.Errorf("Should be nil but Got %v (type %v)", a, reflect.TypeOf(a))
	}
	return nil
}
