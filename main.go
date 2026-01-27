package main

import (
	"aurex-vm/script/compiler"
	"aurex-vm/vm"
)

func main() {

	p := compiler.Compile(`29 + 15 - 2`)

	machine := vm.NewMachine(p)

	machine.Program.Print()
	machine.Execute()
}
