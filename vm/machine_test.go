package vm

import (
	"fmt"
	"sync"
	"testing"

	"github.com/amezianechayer/aurex-vm/script/compiler"
	"github.com/amezianechayer/aurex/core"
)

type CaseResult struct {
	Printed  []uint64
	Postings []core.Posting
	ExitCode byte
}

type TestCase struct {
	Case     string
	Expected CaseResult
}

func test(t *testing.T, c TestCase) {
	p, err := compiler.Compile(c.Case)

	if err != nil {
		t.Error(fmt.Errorf("compile error: %v", err))
	}

	printed := []uint64{}

	var wg sync.WaitGroup
	wg.Add(1)

	machine := NewMachine(p)
	machine.Printer = func(c chan uint64) {
		for v := range c {
			printed = append(printed, v)
		}
		wg.Done()
	}
	machine.Execute(map[string]string{})

	wg.Wait()

	if len(printed) != len(c.Expected.Printed) {
		t.Error(fmt.Errorf("unexpected print output: %v", printed))
		return
	}
	for i := range printed {
		if printed[i] != c.Expected.Printed[i] {
			t.Error(fmt.Errorf("unexpected print output: %v", printed[i]))
			return
		}
	}
}

func TestFail(t *testing.T) {
	test(t, TestCase{
		Case: "fail",
		Expected: CaseResult{
			Printed:  []uint64{},
			Postings: []core.Posting{},
			ExitCode: EXIT_FAIL,
		},
	})
}

func TestPrint(t *testing.T) {
	test(t, TestCase{
		Case: "print 29 + 15 - 2",
		Expected: CaseResult{
			Printed:  []uint64{42},
			Postings: []core.Posting{},
			ExitCode: EXIT_OK,
		},
	})
}

func TestSend(t *testing.T) {
	test(t, TestCase{
		Case: "transfer [DZD.2 100] from @yanis to @ilyes",
		Expected: CaseResult{
			Printed: []uint64{},
			Postings: []core.Posting{
				{
					Asset:       "DZD.2",
					Amount:      100,
					Source:      "@yanis",
					Destination: "@ilyes",
				},
			},
			ExitCode: EXIT_OK,
		},
	})
}
