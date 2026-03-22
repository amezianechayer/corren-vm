package compiler

import (
	"fmt"
	"testing"
)

func TestEndCharacter(t *testing.T) {
	src := `
transfer [DZD.2 200] from @a
send 500% to @b
send 50% to @c
`
	_, err := Compile(src)
	if err == nil {
		t.Fatal("expected error and got none")
	}
	if _, ok := err.(*CompileErrorList); !ok {
		t.Fatal("error had wrong type")
	}
	lerr := err.(*CompileErrorList).errors[0]
	if lerr.startl != 3 {
		t.Fatal(fmt.Sprintf("start line was %v", lerr.startl))
	}
	if lerr.startc != 5 {
		t.Fatal(fmt.Sprintf("start character was %v", lerr.startc))
	}
}
