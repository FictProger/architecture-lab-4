package commands

import (
	"fmt"

	"github.com/FictProger/architecture-lab-4/engine"
)

func CommandCreate(name string, args []string) engine.Command {
	fmt.Println(name)
	fmt.Print(args)
	return nil
}
