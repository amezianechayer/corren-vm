package compiler

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/amezianechayer/corren-vm/core"
	"github.com/amezianechayer/corren-vm/vm/program"
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
					t.Error(fmt.Errorf("%v: %v is not %v: %v",
						p.Resources[i], reflect.TypeOf(p.Resources[i]).Name(),
						c.Expected.Resources[i], reflect.TypeOf(c.Expected.Resources[i]).Name(),
					))
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
			Instructions: []byte{program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00, program.OP_PRINT},
			Resources:    []program.Resource{},
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
		Case:     "fail",
		Expected: CaseResult{Instructions: []byte{program.OP_FAIL}, Resources: []program.Resource{}},
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

func TestCRLF(t *testing.T) {
	test(t, TestCase{
		Case: "print @a\r\nprint @b",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00, program.OP_PRINT,
				program.OP_APUSH, 01, 00, program.OP_PRINT,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Account("@a")},
				program.Constant{Inner: core.Account("@b")},
			},
		},
	})
}

func TestComments(t *testing.T) {
	test(t, TestCase{
		Case: `
/* multi-line comment */
var $a: account
// single-line comment
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
		Case:     "print $nope",
		Expected: CaseResult{Error: "declared"},
	})
}

func TestAllocationFractions(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 43] (
	from @foo
	to {
		1/8 to @bar
		7/8 to @baz
	}
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE,
				program.OP_SWAP,
				program.OP_REPAY,
				program.OP_APUSH, 02, 00,
				program.OP_APUSH, 03, 00,
				program.OP_IPUSH, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_MAKE_ALLOTMENT,
				program.OP_ALLOC,
				program.OP_APUSH, 04, 00,
				program.OP_SEND,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: 43}},
				program.Constant{Inner: core.Account("@foo")},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(7, 8)}},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(1, 8)}},
				program.Constant{Inner: core.Account("@bar")},
				program.Constant{Inner: core.Account("@baz")},
			},
		},
	})
}

func TestAllocationPercentages(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 43] (
	from @foo
	to {
		12.5% to @bar
		37.5% to @baz
		50% to @qux
	}
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE,
				program.OP_SWAP,
				program.OP_REPAY,
				program.OP_APUSH, 02, 00,
				program.OP_APUSH, 03, 00,
				program.OP_APUSH, 04, 00,
				program.OP_IPUSH, 03, 00, 00, 00, 00, 00, 00, 00,
				program.OP_MAKE_ALLOTMENT,
				program.OP_ALLOC,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
				program.OP_APUSH, 06, 00,
				program.OP_SEND,
				program.OP_APUSH, 07, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: 43}},
				program.Constant{Inner: core.Account("@foo")},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(1, 2)}},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(3, 8)}},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(1, 8)}},
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
		Case: `transfer [DZD.2 99] (
	from @alice
	to @bob
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE,
				program.OP_SWAP,
				program.OP_REPAY,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: 99}},
				program.Constant{Inner: alice},
				program.Constant{Inner: bob},
			},
		},
	})
}

func TestTransferAll(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 *] (
	from @alice
	to @bob
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Asset("DZD.2")},
				program.Constant{Inner: core.Account("@alice")},
				program.Constant{Inner: core.Account("@bob")},
			},
		},
	})
}

func TestMetadata(t *testing.T) {
	test(t, TestCase{
		Case: `var $sale: account
var $seller: account = meta($sale, "seller")
var $commission: portion = meta($seller, "commission")
transfer [DZD.2 53] (
	from $sale
	to {
		remaining to $seller
		$commission to @platform
	}
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 03, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 03, 00,
				program.OP_TAKE,
				program.OP_SWAP,
				program.OP_REPAY,
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
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: 53}},
				program.Constant{Inner: core.NewPortionRemaining()},
				program.Constant{Inner: core.Account("@platform")},
			},
		},
	})
}

func TestSyntaxError(t *testing.T) {
	test(t, TestCase{
		Case:     "print fail",
		Expected: CaseResult{Error: "mismatched input"},
	})
}

func TestLogicError(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 200] (
	from 200
	to @bob
)`,
		Expected: CaseResult{Error: "wrong type"},
	})
}

func TestPreventTakeAllFromWorld(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 *] (
	from @world
	to @foo
)`,
		Expected: CaseResult{Error: "cannot"},
	})
}

func TestPreventAddToBottomlessSource(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 1000] (
	from {
		@a
		@world
		@c
	}
	to @out
)`,
		Expected: CaseResult{Error: "world"},
	})
}

func TestPreventSourceAlreadyEmptied(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 1000] (
	from {
		{
			@a
			@world
		}
		@a
	}
	to @out
)`,
		Expected: CaseResult{Error: "@a"},
	})
}

func TestOverflowingAllocation(t *testing.T) {
	fmt.Println("case: >100%")
	test(t, TestCase{
		Case: `transfer [DZD.2 15] (
	from @world
	to {
		2/3 to @a
		2/3 to @b
	}
)`,
		Expected: CaseResult{Error: "100%"},
	})

	fmt.Println("case: =100% + remaining")
	test(t, TestCase{
		Case: `transfer [DZD.2 15] (
	from @world
	to {
		1/2 to @a
		1/2 to @b
		remaining to @c
	}
)`,
		Expected: CaseResult{Error: "100%"},
	})

	fmt.Println("case: variable - remaining")
	test(t, TestCase{
		Case: `var $prop: portion
transfer [DZD.2 15] (
	from @world
	to {
		2/3 to @a
		$prop to @b
	}
)`,
		Expected: CaseResult{Error: "remaining"},
	})
}

func TestAllocationWrongDestination(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 15] (
	from @world
	to [DZD.2 10]
)`,
		Expected: CaseResult{Error: "account"},
	})
}

func TestAllocationInvalidPortion(t *testing.T) {
	test(t, TestCase{
		Case: `var $p: account
transfer [DZD.2 15] (
	from @world
	to {
		10% to @a
		$p to @b
		remaining to @c
	}
)`,
		Expected: CaseResult{Error: "type"},
	})
}

func TestComplexDestination(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 43] (
	from @a
	to {
		1/8 to {
			max [DZD.2 10] to @b
			@c
		}
		7/8 to @d
	}
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE,
				program.OP_SWAP,
				program.OP_REPAY,
				program.OP_APUSH, 02, 00,
				program.OP_APUSH, 03, 00,
				program.OP_IPUSH, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_MAKE_ALLOTMENT,
				program.OP_ALLOC,
				program.OP_APUSH, 04, 00,
				program.OP_TAKE_MAX,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
				program.OP_APUSH, 06, 00,
				program.OP_SEND,
				program.OP_APUSH, 07, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: 43}},
				program.Constant{Inner: core.Account("@a")},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(7, 8)}},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(1, 8)}},
				program.Constant{Inner: core.Monetary{Asset: "DZD.2", Amount: 10}},
				program.Constant{Inner: core.Account("@b")},
				program.Constant{Inner: core.Account("@c")},
				program.Constant{Inner: core.Account("@d")},
			},
		},
	})
}

func TestPreventAddToBottomlessSource2(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 1000] (
	from {
		{
			@a
			@world
		}
		{
			@b
			@world
		}
	}
	to @out
)`,
		Expected: CaseResult{Error: "world"},
	})
}

func TestCappedDestination(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 15] (
	from @world
	to max [DZD.2 15] to @a
)`,
		Expected: CaseResult{Error: "cap"},
	})
}

func TestCappedDestination2(t *testing.T) {
	test(t, TestCase{
		Case: `transfer [DZD.2 15] (
	from @world
	to {
		50% to max [DZD.2 15] to @a
		50% to @b
	}
)`,
		Expected: CaseResult{Error: "cap"},
	})
}
