package commands

import (
	"strconv"

	"github.com/FictProger/architecture-lab-4/engine"
)

func Construct(name string, args []string) engine.Command {
	switch name {
	case "print":
		if len(args) != 1 {
			return &printCommand{"function " + name + " received no arguments"}
		}
		return &printCommand{args[0]}
	case "add":
		if len(args) != 2 {
			return &printCommand{"function " + name + " received no arguments"}
		}
		arg1, _ := strconv.ParseInt(args[0], 10, 64)
		arg2, _ := strconv.ParseInt(args[1], 10, 64)
		return &addCommand{Arg1: int(arg1), Arg2: int(arg2)}
	default:
		return &printCommand{"wrong command"}
	}
}
