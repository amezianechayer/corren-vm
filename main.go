package main

import (
	"encoding/json"
	"fmt"

	"github.com/amezianechayer/aurex-vm/script/compiler"
	"github.com/amezianechayer/aurex-vm/vm"
)

func main() {
	p, err := compiler.Compile(`{
	var $rider: account
	var $driver: account
	var $value: monetary
}
transfer $value from $rider to $driver`) // changé : [DZD.2 999] → $value
	if err != nil {
		panic(err)
	}
	p.Print() // changé : machine.Program.Print() → p.Print()
	machine := vm.NewMachine(p)
	var vars map[string]json.RawMessage
	json.Unmarshal([]byte(`{
		"rider": "user:001",
		"driver": "user:002",
		"value": {
			"asset": "DZD",
			"amount": 999
		}
	}`), &vars)
	exit_code, err := machine.ExecuteFromJSON(vars)
	if err != nil {
		panic(err)
	}
	fmt.Println("EXIT_CODE:", exit_code)
	fmt.Println(machine.Postings)
}
