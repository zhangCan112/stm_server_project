package errors

import (
	"testing"
)

func TestStackInfo(t *testing.T) {
	err := NewError("test")
	t.Error(err.StackInfo())
}
