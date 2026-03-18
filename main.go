package main

import (
	"encoding/json"
	"fmt"

	"github.com/amezianechayer/aurex-vm/script/compiler"
	"github.com/amezianechayer/aurex-vm/vm"
)

func main() {
	p, err := compiler.Compile(`
var $wallet: account
var $payment: account
var $driver: account
var $value: monetary

transfer $value from $wallet then $payment
send 80% to $driver
send 2/25 to @bank
send 12% to @bank2
`)
	if err != nil {
		panic(err)
	}

	p.Print()

	machine := vm.NewMachine(p)

	var vars map[string]json.RawMessage
	json.Unmarshal([]byte(`{
		"wallet":  "user:001",
		"payment": "payments:001",
		"driver":  "user:002",
		"value": {
			"asset":  "DZD",
			"amount": 15
		}
	}`), &vars)

	exit_code, err := machine.ExecuteFromJSON(vars)
	if err != nil {
		panic(err)
	}

	fmt.Println("EXIT_CODE:", exit_code)
	fmt.Println(machine.Postings)
}
