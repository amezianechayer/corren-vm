package main

import (
	"fmt"

	"github.com/amezianechayer/aurex-vm/script/compiler"
	"github.com/amezianechayer/aurex-vm/vm"
)

func main() {

	p := compiler.Compile(`calc 29 + 15 - 2
fail`)

	machine := vm.NewMachine(p)

	machine.Program.Print()
	exit_code := machine.Execute()
	fmt.Println(exit_code)
}
