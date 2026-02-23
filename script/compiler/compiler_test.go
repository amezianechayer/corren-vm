package compiler

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/amezianechayer/aurex-vm/core"
	"github.com/amezianechayer/aurex-vm/vm/program"
)

type CaseResult struct {
	Instructions []byte
	Constants    []core.Value
	Variables    []string
	Error        string
}

type TestCase struct {
	Case     string
	Expected CaseResult
}

func test(t *testing.T, c TestCase) {
	p, err := Compile(c.Case)
	if c.Expected.Error != "" {
		if err == nil {
			t.Error(errors.New("expected error and got none"))
			return
		} else if !strings.Contains(err.Error(), c.Expected.Error) {
			t.Error(fmt.Errorf("error is not the one expected: %v", err))
			return
		}
	} else {
		if err != nil {
			t.Error(fmt.Errorf("did not expect error: %v", err))
			return
		} else if p == nil {
			t.Error(errors.New("did not receive any output"))
			return
		} else if !bytes.Equal(p.Instructions, c.Expected.Instructions) {
			t.Error(fmt.Errorf("generated program is incorrect: %v", p.Instructions))
			return
		} else if len(p.Constants) != len(c.Expected.Constants) {
			t.Error(fmt.Errorf("unexpected program Constants: %v", p.Constants))
			return
		} else {
			for i := range c.Expected.Constants {
				if p.Constants[i] != c.Expected.Constants[i] {
					t.Error(fmt.Errorf("unexpected program Constants: %v", p.Constants))
					return
				}
			}
		}
	}
}

func TestSimplePrint(t *testing.T) {
	test(t, TestCase{
		Case: "print 1",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_PRINT,
			},
			Constants: []core.Value{},
			Error:     "",
		},
	})
}

func TestCompositeExpr(t *testing.T) {
	test(t, TestCase{
		Case: "print 29 + 15 - 2",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_IPUSH, 29, 00, 00, 00, 00, 00, 00, 00,
				program.OP_IPUSH, 15, 00, 00, 00, 00, 00, 00, 00,
				program.OP_IADD,
				program.OP_IPUSH, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_ISUB,
				program.OP_PRINT,
			},
			Constants: []core.Value{},
			Error:     "",
		},
	})
}

func TestFail(t *testing.T) {
	test(t, TestCase{
		Case: "fail",
		Expected: CaseResult{
			Instructions: []byte{program.OP_FAIL},
			Constants:    []core.Value{},
			Error:        "",
		},
	})
}

func TestTransfer(t *testing.T) {
	test(t, TestCase{
		Case: "transfer [DZD.2 99] from @yanis to @ilyes",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00, // yanis (source)
				program.OP_APUSH, 01, 00, // ilyes (dest)
				program.OP_APUSH, 02, 00, // monetary
				program.OP_SEND,
			},
			Constants: []core.Value{
				core.Account("yanis"),
				core.Account("ilyes"),
				core.Monetary{Asset: "DZD.2", Amount: 99},
			},
			Error: "",
		},
	})
}

func TestTransferWithVariables(t *testing.T) {
	test(t, TestCase{
		Case: `vars {
	account $rider
	account $driver
}
transfer [DZD.2 999] from $rider to $driver`,
		Expected: CaseResult{
			Instructions: []byte{
				// $rider — variable addr 0 → 1<<15 + 0 = 0x8000
				program.OP_APUSH, 0x00, 0x80,
				// $driver — variable addr 1 → 1<<15 + 1 = 0x8001
				program.OP_APUSH, 0x01, 0x80,
				// monetary [DZD.2 999] — constante 0
				program.OP_APUSH, 0x00, 0x00,
				program.OP_SEND,
			},
			Constants: []core.Value{
				core.Monetary{Asset: "DZD.2", Amount: 999},
			},
			Variables: []string{"rider", "driver"},
			Error:     "",
		},
	})
}

func TestSyntaxError(t *testing.T) {
	test(t, TestCase{
		Case: "print fail",
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "mismatched input",
		},
	})
}

func TestLogicError(t *testing.T) {
	test(t, TestCase{
		Case: "transfer [DZD.2 100] from 100 to @ilyes",
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "wrong argument type",
		},
	})
}

func TestDuplicateVariable(t *testing.T) {
	test(t, TestCase{
		Case: `vars {
	account $rider
	account $rider
}
fail`,
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "duplicate variable",
		},
	})
}

func TestUndeclaredVariable(t *testing.T) {
	test(t, TestCase{
		Case: "transfer [DZD.2 100] from $unknown to @bank",
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "variable not declared",
		},
	})
}

// func TestTooManyConstants(t *testing.T) {
// 	script := "print 1"
// 	for i := 0; i < 20000; i++ {
// 		script += fmt.Sprintf("\ntransfer [A%d.2 0] from @a%d to @b%d", i, i, i)
// 	}
// 	test(t, TestCase{
// 		Case: script,
// 		Expected: CaseResult{
// 			Instructions: nil,
// 			Constants:    nil,
// 			Error:        "exceeded",
// 		},
// 	})
// }
