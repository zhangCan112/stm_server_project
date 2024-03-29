package errors

import (
	"runtime"
	"testing"
)

func TestStackTraceModel(t *testing.T) {
	st := Caller(0)
	if st == nil {
		t.Errorf("TestStackTrace.Error(): Caller returns nil")
	}

	if st.FuncName() != "TestStackTraceModel" {
		t.Errorf("TestStackTrace.Error(): FuncName got: %s want: %s", st.FuncName(), "TestStackTraceModel")
	}

	if st.Line() != 9 {
		t.Errorf("TestStackTrace.Error(): Line got: %d want: %s", st.Line(), "9")
	}

	_, file, _, _ := runtime.Caller(0)
	if st.File() != file {
		t.Errorf("TestStackTrace.Error(): File got: %s want: %s", st.File(), file)
	}

}
