package gojs_test

import (
	"context"
	"testing"

	"github.com/powerpuffpenguin/gojs"
)

func TestRuntime(t *testing.T) {
	vm := gojs.New()
	_, e := vm.RunScript("bin/main.js", nil)
	if e != nil {
		t.Fatal(e)
	}
	e = vm.RunLoop()
	if e != nil && e != context.Canceled {
		t.Fatal(e)
	}
}
