package commands

import (
	"fmt" // Package fmt implements formatted I/O with functions analogous to C's
	// printf and scanf. The format 'verbs' are derived from C's but are simpler.
	"github.com/FictProger/architecture-lab-4/engine"
)

type PrintCommand struct {
	Arg string
}

func (print *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(print.Arg)
}
