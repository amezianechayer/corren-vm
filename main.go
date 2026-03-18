package main

import (
	"encoding/json"
	"fmt"

	"github.com/amezianechayer/aurex-vm/script/compiler"
	"github.com/amezianechayer/aurex-vm/vm"
)

func main() {
	p, err := compiler.Compile(`
var $balance: account
var $payment: account
var $seller: account

transfer [DZD.2 15] from $balance then $payment to $seller
`)
	if err != nil {
		panic(err)
	}

	p.Print()

	machine := vm.NewMachine(p)

	var vars map[string]json.RawMessage
	json.Unmarshal([]byte(`{
		"balance": "@users:001",
		"payment": "@payments:001",
		"seller":  "@users:002"
	}`), &vars)

	err = machine.SetVarsFromJSON(vars)
	if err != nil {
		panic(err)
	}

	err = machine.SetBalances(map[string]map[string]uint64{
		"@users:001": {
			"DZD.2": 15,
		},
		"@payments:001": {
			"DZD.2": 0,
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
