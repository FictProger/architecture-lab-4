package commands

import (
	"strconv" // The built-in package `strconv` provides the number parsing.

	"github.com/FictProger/architecture-lab-4/engine"
)

type addCommand struct {
	Arg1, Arg2 int
}

func (add *addCommand) Execute(loop engine.Handler) {
	res := add.Arg1 + add.Arg2
	loop.Post(&printCommand{Arg: strconv.Itoa(res)})
}
