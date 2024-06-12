package assert

import (
	"fmt"
	"testing"
)

func EqualSliceErrorf[T comparable](t *testing.T, expect, actual []T, format string, args ...any) {
	if len(expect) != len(actual) {
		t.Errorf("\nexpect:%v\nactual:%v\nmsg:%s\n", expect, actual,
			fmt.Sprintf(format, args...))
		return
	}
	for i := range expect {
		if expect[i] != actual[i] {
			t.Errorf("\nexpect:%v\nactual:%v\nmsg:%s\n", expect, actual,
				fmt.Sprintf(format, args...))
			return
		}
	}
}
func EqualSliceFatalf[T comparable](t *testing.T, expect, actual []T, format string, args ...any) {
	if len(expect) != len(actual) {
		t.Fatalf("\nexpect:%v\nactual:%v\nmsg:%s\n", expect, actual,
			fmt.Sprintf(format, args...))
		return
	}
	for i := range expect {
		if expect[i] != actual[i] {
			t.Fatalf("\nexpect:%v\nactual:%v\nmsg:%s\n", expect, actual,
				fmt.Sprintf(format, args...))
			return
		}
	}
}
func EqualErrorf[T comparable](t *testing.T, expect, actual T, format string, args ...any) {
	if expect != actual {
		t.Errorf("\nexpect:%v\nactual:%v\nmsg:%s\n", expect, actual,
			fmt.Sprintf(format, args...))
	}
}

func EqualFatalf[T comparable](t *testing.T, expect, actual T, format string, args ...any) {
	if expect != actual {
		t.Fatalf("\nexpect:%v\nactual:%v\nmsg:%s\n", expect, actual,
			fmt.Sprintf(format, args...))
	}
}
