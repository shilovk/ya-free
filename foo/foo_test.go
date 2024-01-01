// файл foo_test.go
package foo_test

import (
	"testing"
	"ya-free/foo"
)

func TestFooFunc(t *testing.T) {
	expectedFooResult := "bar"
	if actualFooResult := foo.Foo(); actualFooResult != expectedFooResult {
		t.Errorf("expected %s; got: %s", expectedFooResult, actualFooResult)
	}
}
