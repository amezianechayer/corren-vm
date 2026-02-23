package main

import (
	"fmt"

	"github.com/amezianechayer/aurex-vm/core"
	"github.com/amezianechayer/aurex-vm/script/compiler"
	"github.com/amezianechayer/aurex-vm/vm"
)

func main() {
	// exemple simple sans variables
	// p, err := compiler.Compile(`
	// print 29 + 15 - 2
	// transfer [DZD.2 100] from @alice to @bob
	// fail`)

	p, err := compiler.Compile(`vars {
		account $rider
		account $driver
	}
	transfer [DZD.2 999] from $rider to $driver`)
	if err != nil {
		panic(err)
	}

	machine := vm.NewMachine(p)
	machine.Program.Print()

	exit_code := machine.Execute(map[string]core.Value{
		"rider":  core.Account("user:001"),
		"driver": core.Account("user:002"),
	})

	fmt.Println(exit_code)
	fmt.Println(machine.Postings)
}
