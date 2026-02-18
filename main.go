package main

import (
	"fmt"

	"github.com/amezianechayer/aurex-vm/script/compiler"
	"github.com/amezianechayer/aurex-vm/vm"
)

func main() {

	p, err := compiler.Compile(`print 29 + 15 - 2
transfer [DZD.2 100] from @yanis to @ilyes
fail`)

	// fmt.Println(p)

	if err != nil {
		panic(err)
	}

	machine := vm.NewMachine(p)

	machine.Program.Print()
	exit_code := machine.Execute(map[string]string{})
	fmt.Println(exit_code)
	fmt.Println(machine.Postings)
}
