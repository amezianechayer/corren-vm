package compiler

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
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
			t.Error(fmt.Errorf("generated program is incorrect: %v", *p))
			fmt.Println(p.Instructions, "vs", c.Expected.Instructions)
			return
		} else if len(p.Constants) != len(c.Expected.Constants) {
			t.Error(fmt.Errorf("unexpected program constants: %v", *p))
			return
		} else {
			for i := range c.Expected.Constants {
				if !core.ValueEquals(p.Constants[i], c.Expected.Constants[i]) {
					t.Error(fmt.Errorf("unexpected program constants: %v", *p))
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

func TestConstant(t *testing.T) {
	user := core.Account("@user:001")
	test(t, TestCase{
		Case: "print @user:001",
		Expected: CaseResult{
			Instructions: []byte{program.OP_APUSH, 00, 00, program.OP_PRINT},
			Constants:    []core.Value{user},
			Error:        "",
		},
	})
}

func TestAllocationFractions(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 43] from @foo
send 1/8 to @bar
send 7/8 to @baz
`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_ALLOC,
				program.OP_APUSH, 03, 00,
				program.OP_SEND,
				program.OP_APUSH, 04, 00,
				program.OP_SEND,
			},
			Constants: []core.Value{
				core.Monetary{Asset: "DZD.2", Amount: core.NewAmountSpecific(43)},
				core.Account("@foo"),
				core.Allotment{*big.NewRat(1, 8), *big.NewRat(7, 8)},
				core.Account("@bar"),
				core.Account("@baz"),
			},
			Error: "",
		},
	})
}

func TestAllocationPercentages(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 43] from @foo
send 12.5% to @bar
send 37.5% to @baz
send 50% to @qux
`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_ALLOC,
				program.OP_APUSH, 03, 00,
				program.OP_SEND,
				program.OP_APUSH, 04, 00,
				program.OP_SEND,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
			},
			Constants: []core.Value{
				core.Monetary{Asset: "DZD.2", Amount: core.NewAmountSpecific(43)},
				core.Account("@foo"),
				core.Allotment{*big.NewRat(125, 1000), *big.NewRat(375, 1000), *big.NewRat(1, 2)},
				core.Account("@bar"),
				core.Account("@baz"),
				core.Account("@qux"),
			},
			Error: "",
		},
	})
}

func TestTransfer(t *testing.T) {
	alice := core.Account("@alice")
	bob := core.Account("@bob")
	test(t, TestCase{
		Case: "transfer [DZD.2 99] from @alice to @bob",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			},
			Constants: []core.Value{
				core.Monetary{Asset: "DZD.2", Amount: core.NewAmountSpecific(99)},
				alice,
				bob,
			},
			Error: "",
		},
	})
}

func TestTransferAll(t *testing.T) {
	alice := core.Account("@alice")
	bob := core.Account("@bob")
	test(t, TestCase{
		Case: "transfer [DZD.2 *] from @alice to @bob",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			},
			Constants: []core.Value{
				core.Monetary{Asset: "DZD.2", Amount: core.NewAmountAll()},
				alice,
				bob,
			},
			Error: "",
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
		Case: "transfer [DZD.2 200] from 200 to @bob",
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "expected",
		},
	})
}

func TestPreventTakeAllFromWorld(t *testing.T) {
	test(t, TestCase{
		Case: "transfer [DZD.2 *] from @world to @foo",
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "cannot",
		},
	})
}

func TestOverflowingAllocation(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 15] from @world
send 2/3 to @a
send 2/3 to @b
`,
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "100%",
		},
	})

	test(t, TestCase{
		Case: `transfer [DZD.2 15] from @world
send 1/2 to @a
send 1/2 to @b
keep remaining
`,
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "100%",
		},
	})

	test(t, TestCase{
		Case: `transfer [DZD.2 15] from @world
send 2/3 to @a
send 1/2 to @b
keep remaining
`,
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "100%",
		},
	})

	test(t, TestCase{
		Case: `var $prop: portion
transfer [DZD.2 15] from @world
send 2/3 to @a
send $prop to @b
`,
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "remaining",
		},
	})
}
