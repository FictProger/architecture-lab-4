package commands

import (
	"strconv" // The built-in package `strconv` provides the number parsing.

	"github.com/FictProger/architecture-lab-4/engine"
)

type AddCommand struct {
	Arg1, Arg2 int
}

func (add *AddCommand) Execute(loop engine.Handler) {
	res := add.Arg1 + add.Arg2
	loop.Post(&PrintCommand{Arg: strconv.Itoa(res)})
}
