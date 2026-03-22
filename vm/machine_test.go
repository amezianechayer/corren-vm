package vm

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"testing"

	"github.com/amezianechayer/aurex-vm/core"
	"github.com/amezianechayer/aurex-vm/script/compiler"
	ledger "github.com/amezianechayer/aurex/core"
)

type CaseResult struct {
	Printed  []core.Value
	Postings []ledger.Posting
	ExitCode byte
	Error    string
}

func test(
	t *testing.T,
	code string,
	variables map[string]core.Value,
	meta map[string]map[string]core.Value,
	balances map[string]map[string]uint64,
	expected CaseResult,
) {
	testimpl(t, code, expected, func(m *Machine) (byte, error) {
		err := m.SetVars(variables)
		if err != nil {
			return 0, err
		}
		{
			ch, err := m.ResolveResources()
			if err != nil {
				return 0, err
			}
			for req := range ch {
				if req.Error != nil {
					panic(req.Error)
				}
				val := meta[req.Account][req.Key]
				req.Response <- val
			}
		}
		{
			ch, err := m.ResolveBalances()
			if err != nil {
				return 0, err
			}
			for req := range ch {
				if req.Error != nil {
					panic(req.Error)
				}
				val := balances[req.Account][req.Asset]
				req.Response <- val
			}
		}
		return m.Execute()
	})
}

func testJSON(
	t *testing.T,
	code string,
	variables string,
	meta map[string]map[string]core.Value,
	balances map[string]map[string]uint64,
	expected CaseResult,
) {
	testimpl(t, code, expected, func(m *Machine) (byte, error) {
		var v map[string]json.RawMessage
		err := json.Unmarshal([]byte(variables), &v)
		if err != nil {
			return 0, err
		}
		err = m.SetVarsFromJSON(v)
		if err != nil {
			return 0, err
		}
		{
			ch, err := m.ResolveResources()
			if err != nil {
				return 0, err
			}
			for req := range ch {
				if req.Error != nil {
					panic(req.Error)
				}
				val, ok := meta[req.Account][req.Key]
				if !ok {
					t.Fatalf("case error: missing metadata %q of %v", req.Key, req.Account)
				}
				req.Response <- val
			}
		}
		{
			ch, err := m.ResolveBalances()
			if err != nil {
				return 0, err
			}
			for req := range ch {
				if req.Error != nil {
					return 0, req.Error
				}
				val, ok := balances[req.Account][req.Asset]
				if !ok {
					return 0, fmt.Errorf("missing %v balance of %v", req.Asset, req.Account)
				}
				req.Response <- val
			}
		}
		return m.Execute()
	})
}

func testimpl(t *testing.T, code string, expected CaseResult, exec func(*Machine) (byte, error)) {
	p, err := compiler.Compile(code)
	if err != nil {
		t.Error(fmt.Errorf("compile error: %v", err))
		return
	}
	printed := []core.Value{}
	var wg sync.WaitGroup
	wg.Add(1)
	machine := NewMachine(p)
	machine.Printer = func(c chan core.Value) {
		for v := range c {
			printed = append(printed, v)
		}
		wg.Done()
	}
	exit_code, err := exec(machine)

	if err != nil && expected.Error != "" {
		if !strings.Contains(err.Error(), expected.Error) {
			t.Error(fmt.Errorf("unexpected execution error: %v", err))
			return
		} else {
			return
		}
	} else if err != nil {
		t.Error(fmt.Errorf("did not expect an execution error: %v", err))
		return
	} else if expected.Error != "" {
		t.Error(fmt.Errorf("expected an execution error"))
		return
	}

	if exit_code != expected.ExitCode {
		t.Error(fmt.Errorf("unexpected exit code: %v", exit_code))
		return
	}

	if len(machine.Postings) != len(expected.Postings) {
		t.Error(fmt.Errorf("unexpected postings output: %v", machine.Postings))
		return
	} else {
		for i := range machine.Postings {
			if machine.Postings[i] != expected.Postings[i] {
				t.Error(fmt.Errorf("unexpected postings output: %v", machine.Postings[i]))
				return
			}
		}
	}

	wg.Wait()

	if len(printed) != len(expected.Printed) {
		t.Error(fmt.Errorf("unexpected print output: %v", printed))
		return
	} else {
		for i := range printed {
			if printed[i] != expected.Printed[i] {
				t.Error(fmt.Errorf("unexpected print output: %v", printed[i]))
				return
			}
		}
	}
}

func TestFail(t *testing.T) {
	test(t, "fail", map[string]core.Value{}, map[string]map[string]core.Value{}, map[string]map[string]uint64{},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{}, ExitCode: EXIT_FAIL})
}

func TestPrint(t *testing.T) {
	test(t, "print 29 + 15 - 2", map[string]core.Value{}, map[string]map[string]core.Value{}, map[string]map[string]uint64{},
		CaseResult{Printed: []core.Value{core.Number(42)}, Postings: []ledger.Posting{}, ExitCode: EXIT_OK})
}

func TestTransfer(t *testing.T) {
	test(t, "transfer [DZD.2 100] from @alice to @bob", map[string]core.Value{},
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@alice": {"DZD.2": 100}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 100, Source: "@alice", Destination: "@bob"},
		}, ExitCode: EXIT_OK})
}

func TestVariables(t *testing.T) {
	test(t,
		`var $rider: account
var $driver: account
transfer [DZD.2 999] from $rider to $driver`,
		map[string]core.Value{"rider": core.Account("@users:001"), "driver": core.Account("@users:002")},
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@users:001": {"DZD.2": 1000}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 999, Source: "@users:001", Destination: "@users:002"},
		}, ExitCode: EXIT_OK})
}

func TestVariablesJSONInvalid(t *testing.T) {
	testJSON(t,
		`var $p: portion
transfer [DZD.2 999] from @world
send $p to @b
keep remaining
`,
		`{"p": "3/2"}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{},
		CaseResult{Error: "portion must be"})
}

func TestVariablesJSON(t *testing.T) {
	testJSON(t,
		`var $rider: account
var $driver: account
transfer [DZD.2 999] from $rider to $driver`,
		`{"rider": "@users:001", "driver": "@users:002"}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@users:001": {"DZD.2": 1000}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 999, Source: "@users:001", Destination: "@users:002"},
		}, ExitCode: EXIT_OK})
}

func TestSource(t *testing.T) {
	testJSON(t,
		`var $balance: account
var $payment: account
var $seller: account
transfer [DZD.2 15] from $balance then $payment to $seller`,
		`{"balance": "@users:001", "payment": "@payments:001", "seller": "@users:002"}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@users:001": {"DZD.2": 3}, "@payments:001": {"DZD.2": 12}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 3, Source: "@users:001", Destination: "@users:002"},
			{Asset: "DZD.2", Amount: 12, Source: "@payments:001", Destination: "@users:002"},
		}, ExitCode: EXIT_OK})
}

func TestAllocation(t *testing.T) {
	testJSON(t,
		`var $rider: account
var $driver: account
transfer [DZD.2 15] from $rider
send 80% to $driver
send 8% to @a
send 12% to @b
`,
		`{"rider": "@users:001", "driver": "@users:002"}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@users:001": {"DZD.2": 15}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 13, Source: "@users:001", Destination: "@users:002"},
			{Asset: "DZD.2", Amount: 1, Source: "@users:001", Destination: "@a"},
			{Asset: "DZD.2", Amount: 1, Source: "@users:001", Destination: "@b"},
		}, ExitCode: EXIT_OK})
}

func TestDynamicAllocation(t *testing.T) {
	testJSON(t,
		`var $p: portion
transfer [DZD.2 15] from @a
send 80% to @b
send $p to @c
keep remaining
`,
		`{"p": "15%"}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@a": {"DZD.2": 15}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 13, Source: "@a", Destination: "@b"},
			{Asset: "DZD.2", Amount: 2, Source: "@a", Destination: "@c"},
		}, ExitCode: EXIT_OK})
}

func TestSendAll(t *testing.T) {
	testJSON(t, "transfer [DZD.2 *] from @users:001 to @platform", `{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@users:001": {"DZD.2": 17}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 17, Source: "@users:001", Destination: "@platform"},
		}, ExitCode: EXIT_OK})
}

func TestSendAllMulti(t *testing.T) {
	testJSON(t, "transfer [DZD.2 *] from @users:001:wallet then @users:001:credit to @platform", `{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"@users:001:wallet": {"DZD.2": 19},
			"@users:001:credit": {"DZD.2": 22},
		},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 19, Source: "@users:001:wallet", Destination: "@platform"},
			{Asset: "DZD.2", Amount: 22, Source: "@users:001:credit", Destination: "@platform"},
		}, ExitCode: EXIT_OK})
}

func TestInsufficientFunds(t *testing.T) {
	testJSON(t,
		`var $balance: account
var $payment: account
var $seller: account
transfer [DZD.2 16] from $balance then $payment to $seller`,
		`{"balance": "@users:001", "payment": "@payments:001", "seller": "@users:002"}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@users:001": {"DZD.2": 3}, "@payments:001": {"DZD.2": 12}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{}, ExitCode: EXIT_FAIL})
}

func TestMissingBalance(t *testing.T) {
	testJSON(t, "transfer [DZD.2 15] from @alice to @bob", `{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@users:001": {"DZD.2": 3}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{}, ExitCode: 0, Error: "missing"})
}

func TestMissingWorldBalance(t *testing.T) {
	testJSON(t, "transfer [DZD.2 15] from @world to @a", `{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 15, Source: "@world", Destination: "@a"},
		}, ExitCode: EXIT_OK})
}

func TestWorldSource(t *testing.T) {
	testJSON(t,
		`var $a: account
var $b: account
transfer [DZD.2 15] from $a then @world to $b`,
		`{"a": "@a", "b": "@b"}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@a": {"DZD.2": 1}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 1, Source: "@a", Destination: "@b"},
			{Asset: "DZD.2", Amount: 14, Source: "@world", Destination: "@b"},
		}, ExitCode: EXIT_OK})
}

func TestNoEmptyPostings(t *testing.T) {
	testJSON(t,
		`transfer [DZD.2 2] from @world
send 90% to @a
send 10% to @b
`,
		`{}`, map[string]map[string]core.Value{}, map[string]map[string]uint64{},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 2, Source: "@world", Destination: "@a"},
		}, ExitCode: EXIT_OK})
}

func TestNoEmptyPostings2(t *testing.T) {
	testJSON(t, "transfer [DZD.2 *] from @foo to @bar", `{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@foo": {"DZD.2": 0}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{}, ExitCode: EXIT_OK})
}

func TestAllocateDontTakeTooMuch(t *testing.T) {
	testJSON(t,
		`var $u1: account
var $u2: account
transfer [CREDIT 200] from $u1 then $u2
send 1/2 to @foo
send 1/2 to @bar
`,
		`{"u1": "@users:001", "u2": "@users:002"}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{"@users:001": {"CREDIT": 100}, "@users:002": {"CREDIT": 100}},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "CREDIT", Amount: 100, Source: "@users:001", Destination: "@foo"},
			{Asset: "CREDIT", Amount: 100, Source: "@users:002", Destination: "@bar"},
		}, ExitCode: EXIT_OK})
}

func TestNeededBalances(t *testing.T) {
	p, err := compiler.Compile(`
var $a: account
transfer [DZD.2 15] from $a then @b then @world to @c
`)
	if err != nil {
		t.Fatalf("did not expect error on Compile, got: %v", err)
	}

	m := NewMachine(p)
	err = m.SetVars(map[string]core.Value{"a": core.Account("@a")})
	if err != nil {
		t.Fatalf("did not expect error on SetVars, got: %v", err)
	}
	{
		ch, err := m.ResolveResources()
		if err != nil {
			t.Fatalf("did not expect error on ResolveResources, got: %v", err)
		}
		for range ch {
			t.Fatalf("did not expect to need any metadata")
		}
	}

	expected := map[string]map[string]struct{}{
		"@a": {"DZD.2": {}},
		"@b": {"DZD.2": {}},
	}
	{
		ch, err := m.ResolveBalances()
		if err != nil {
			t.Fatalf("did not expect error on ResolveBalances, got: %v", err)
		}
		for req := range ch {
			if req.Error != nil {
				t.Fatalf("did not expect error in balance request: %v", req.Error)
			}
			if _, ok := expected[req.Account][req.Asset]; ok {
				delete(expected[req.Account], req.Asset)
				if len(expected[req.Account]) == 0 {
					delete(expected, req.Account)
				}
				req.Response <- 0
			} else {
				t.Fatalf("did not expect to need %v balance of %v", req.Asset, req.Account)
			}
		}
	}
	if len(expected) != 0 {
		t.Fatalf("some balances were not requested: %v", expected)
	}
}

func TestMetadata(t *testing.T) {
	commission, _ := core.NewPortionSpecific(*big.NewRat(125, 1000))
	testJSON(t,
		`var $sale: account
var $seller: account = meta($sale, "seller")
var $commission: portion = meta($seller, "commission")
transfer [DZD.2 100] from $sale
send remaining to $seller
send $commission to @platform
`,
		`{"sale": "@sales:042"}`,
		map[string]map[string]core.Value{
			"@sales:042": {
				"seller": core.Account("@users:053"),
			},
			"@users:053": {
				"commission": *commission,
			},
		},
		map[string]map[string]uint64{
			"@sales:042": {"DZD.2": 2500},
			"@users:053": {"DZD.2": 500},
		},
		CaseResult{Printed: []core.Value{}, Postings: []ledger.Posting{
			{Asset: "DZD.2", Amount: 88, Source: "@sales:042", Destination: "@users:053"},
			{Asset: "DZD.2", Amount: 12, Source: "@sales:042", Destination: "@platform"},
		}, ExitCode: EXIT_OK})
}
