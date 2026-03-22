package compiler

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/amezianechayer/aurex-vm/core"
	"github.com/amezianechayer/aurex-vm/vm/program"
)

type CaseResult struct {
	Instructions []byte
	Resources    []program.Resource
	Error        string
}

type TestCase struct {
	Case     string
	Expected CaseResult
}

func checkResourceEquals(t *testing.T, res program.Resource, expected program.Resource) bool {
	if reflect.TypeOf(res) != reflect.TypeOf(expected) {
		return false
	}
	switch res := res.(type) {
	case program.Constant:
		if reflect.TypeOf(res.Inner).Kind() == reflect.Ptr {
			t.Fatal("generated program contained a constant with a pointer value")
		}
		return core.ValueEquals(res.Inner, expected.(program.Constant).Inner)
	case program.Parameter:
		e := expected.(program.Parameter)
		return res.Typ == e.Typ && res.Name == e.Name
	case program.Metadata:
		e := expected.(program.Metadata)
		return res.SourceAccount == e.SourceAccount && res.Key == e.Key && res.Typ == e.Typ
	default:
		panic("invalid resource")
	}
}

func test(t *testing.T, c TestCase) {
	p, err := Compile(c.Case)

	if c.Expected.Error != "" {
		if err == nil {
			t.Error(errors.New("expected error and got none"))
			return
		} else if err.Error() == "" {
			t.Error(errors.New("error was not fed to the error listener"))
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
		} else if len(p.Resources) != len(c.Expected.Resources) {
			t.Error(fmt.Errorf("unexpected program resources (=/= lengths): %v", *p))
			return
		} else {
			for i := range c.Expected.Resources {
				if !checkResourceEquals(t, p.Resources[i], c.Expected.Resources[i]) {
					t.Error(fmt.Errorf("%v is not %v", p.Resources[i], c.Expected.Resources[i]))
					t.Error(fmt.Errorf("unexpected program resources: %v", *p))
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
			Resources: []program.Resource{},
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
			Resources: []program.Resource{},
		},
	})
}

func TestFail(t *testing.T) {
	test(t, TestCase{
		Case: "fail",
		Expected: CaseResult{
			Instructions: []byte{program.OP_FAIL},
			Resources:    []program.Resource{},
		},
	})
}

func TestConstant(t *testing.T) {
	user := core.Account("@user:001")
	test(t, TestCase{
		Case: "print @user:001",
		Expected: CaseResult{
			Instructions: []byte{program.OP_APUSH, 00, 00, program.OP_PRINT},
			Resources:    []program.Resource{program.Constant{Inner: user}},
		},
	})
}

func TestComments(t *testing.T) {
	test(t, TestCase{
		Case: `
/* This is a multi-line comment, it spans multiple lines
and /* doesn't choke on nested comments */ ! */
var $a: account
// this is a single-line comment
print $a
`,
		Expected: CaseResult{
			Instructions: []byte{program.OP_APUSH, 00, 00, program.OP_PRINT},
			Resources:    []program.Resource{program.Parameter{Typ: core.TYPE_ACCOUNT, Name: "a"}},
		},
	})
}

func TestUndeclaredVariable(t *testing.T) {
	test(t, TestCase{
		Case: "print $nope",
		Expected: CaseResult{
			Error: "declared",
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
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: core.NewAmountSpecific(43)}},
				program.Constant{Inner: core.Account("@foo")},
				program.Constant{Inner: core.Allotment{*big.NewRat(1, 8), *big.NewRat(7, 8)}},
				program.Constant{Inner: core.Account("@bar")},
				program.Constant{Inner: core.Account("@baz")},
			},
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
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: core.NewAmountSpecific(43)}},
				program.Constant{Inner: core.Account("@foo")},
				program.Constant{Inner: core.Allotment{*big.NewRat(125, 1000), *big.NewRat(375, 1000), *big.NewRat(1, 2)}},
				program.Constant{Inner: core.Account("@bar")},
				program.Constant{Inner: core.Account("@baz")},
				program.Constant{Inner: core.Account("@qux")},
			},
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
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: core.NewAmountSpecific(99)}},
				program.Constant{Inner: alice},
				program.Constant{Inner: bob},
			},
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
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: core.NewAmountAll()}},
				program.Constant{Inner: alice},
				program.Constant{Inner: bob},
			},
		},
	})
}

func TestMetadata(t *testing.T) {
	test(t, TestCase{
		Case: `var $sale: account
var $seller: account = meta($sale, "seller")
var $commission: portion = meta($seller, "commission")
transfer [DZD.2 53] from $sale
send remaining to $seller
send $commission to @platform
`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 03, 00,
				program.OP_APUSH, 00, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_APUSH, 04, 00,
				program.OP_IPUSH, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_MAKE_ALLOTMENT,
				program.OP_ALLOC,
				program.OP_APUSH, 01, 00,
				program.OP_SEND,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Parameter{Name: "sale", Typ: core.TYPE_ACCOUNT},
				program.Metadata{SourceAccount: core.NewAddress(0), Key: "seller", Typ: core.TYPE_ACCOUNT},
				program.Metadata{SourceAccount: core.NewAddress(1), Key: "commission", Typ: core.TYPE_PORTION},
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: core.NewAmountSpecific(53)}},
				program.Constant{Inner: core.NewPortionRemaining()},
				program.Constant{Inner: core.Account("@platform")},
			},
		},
	})
}

func TestSyntaxError(t *testing.T) {
	test(t, TestCase{
		Case: "print fail",
		Expected: CaseResult{
			Error: "mismatched input",
		},
	})
}

func TestLogicError(t *testing.T) {
	test(t, TestCase{
		Case: "transfer [DZD.2 200] from 200 to @bob",
		Expected: CaseResult{
			Error: "wrong type",
		},
	})
}

func TestPreventTakeAllFromWorld(t *testing.T) {
	test(t, TestCase{
		Case: "transfer [DZD.2 *] from @world to @foo",
		Expected: CaseResult{
			Error: "cannot",
		},
	})
}

func TestOverflowingAllocation(t *testing.T) {
	fmt.Println("case: >100%")
	test(t, TestCase{
		Case: `transfer [DZD.2 15] from @world
send 2/3 to @a
send 2/3 to @b
`,
		Expected: CaseResult{Error: "100%"},
	})

	fmt.Println("case: =100% + remaining")
	test(t, TestCase{
		Case: `transfer [DZD.2 15] from @world
send 1/2 to @a
send 1/2 to @b
keep remaining
`,
		Expected: CaseResult{Error: "100%"},
	})

	fmt.Println("case: >100% + remaining")
	test(t, TestCase{
		Case: `transfer [DZD.2 15] from @world
send 2/3 to @a
send 1/2 to @b
keep remaining
`,
		Expected: CaseResult{Error: "100%"},
	})

	fmt.Println("case: const remaining + remaining")
	test(t, TestCase{
		Case: `transfer [DZD.2 15] from @world
send 2/3 to @a
keep remaining
keep remaining
`,
		Expected: CaseResult{Error: "`remaining` in the same"},
	})

	fmt.Println("case: dyn remaining + remaining")
	test(t, TestCase{
		Case: `var $p: portion
transfer [DZD.2 15] from @world
send $p to @a
keep remaining
keep remaining
`,
		Expected: CaseResult{Error: "`remaining` in the same"},
	})

	fmt.Println("case: >100% + remaining + variable")
	test(t, TestCase{
		Case: `var $prop: portion
transfer [DZD.2 15] from @world
send 1/2 to @a
send 2/3 to @b
keep remaining
send $prop to @d
`,
		Expected: CaseResult{Error: "100%"},
	})

	fmt.Println("case: variable - remaining")
	test(t, TestCase{
		Case: `var $prop: portion
transfer [DZD.2 15] from @world
send 2/3 to @a
send $prop to @b
`,
		Expected: CaseResult{Error: "remaining"},
	})
}

func TestAllocationWrongDestination(t *testing.T) {
	test(t, TestCase{
		Case:     `transfer [DZD.2 15] from @world to [DZD.2 10]`,
		Expected: CaseResult{Error: "wrong type"},
	})
	test(t, TestCase{
		Case: `transfer [DZD.2 15] from @world
send 2/3 to @a
send 1/3 to [DZD.2 10]
`,
		Expected: CaseResult{Error: "account"},
	})
}

func TestAllocationInvalidPortion(t *testing.T) {
	test(t, TestCase{
		Case: `var $p: account
transfer [DZD.2 15] from @world
send 10% to @a
send $p to @b
keep remaining
`,
		Expected: CaseResult{Error: "type"},
	})
}
