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
var $commission: portion

transfer $value from @users:001
send remaining to @seller
send $commission to @platform
`)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	machine := vm.NewMachine(p)

	var vars map[string]json.RawMessage
	json.Unmarshal([]byte(`{
		"value": {
			"asset": "DZD.2",
			"amount": 45
		},
		"commission": "12.5%"
	}`), &vars)

	err = machine.SetVarsFromJSON(vars)
	if err != nil {
		panic(err)
	}

	err = machine.SetBalances(map[string]map[string]uint64{
		"@users:001": {
			"DZD.2": 2500,
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
