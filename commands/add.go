package commands

import (
	"strconv" // The built-in package `strconv` provides the number parsing.
	"github.com/FictProger/architecture-lab-4/engine"
)

type addCommand struct {
	arg1, arg2 int
}

func (add *addCommand) Execute(loop engine.Handler) {
	res := add.arg1 + add.arg2
	loop.Post(&printCommand{arg: strconv.Itoa(res)})
}