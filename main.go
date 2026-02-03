package main

import (
	"fmt"

	"github.com/amezianechayer/aurex-vm/script/compiler"
	"github.com/amezianechayer/aurex-vm/vm"
)

func main() {

	p, err := compiler.Compile(`calc 29 + 15 - 2 
fail`)

	fmt.Println(p)

	if err != nil {
		panic(err)
	}

	machine := vm.NewMachine(p)

	machine.Program.Print()
	exit_code := machine.Execute()
	fmt.Println(exit_code)
}
