package main

import (
	"encoding/json"
	"fmt"

	"github.com/amezianechayer/aurex-vm/script/compiler"
	"github.com/amezianechayer/aurex-vm/vm"
)

func main() {
	p, err := compiler.Compile(`
var $value: monetary

transfer $value from @a then @b
send 80% to @c
send 8% to @d
send 12% to @e
`)
	if err != nil {
		panic(err)
	}

	p.Print()

	machine := vm.NewMachine(p)

	var vars map[string]json.RawMessage
	json.Unmarshal([]byte(`{
		"value": {
			"asset": "DZD.2",
			"amount": 45
		}
	}`), &vars)

	fmt.Println(vars)

	err = machine.SetVarsFromJSON(vars)
	if err != nil {
		panic(err)
	}

	err = machine.SetBalances(map[string]map[string]uint64{
		"@a": {
			"DZD.2": 3,
		},
		"@b": {
			"DZD.2": 25,
		},
	})
	if err != nil {
		panic(err)
	}

	exit_code, err := machine.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Println("EXIT_CODE:", exit_code)
	fmt.Println(machine.Postings)
}
