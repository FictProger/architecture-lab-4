package main

import (
	"github.com/FictProger/architecture-lab-4/commands"
	"github.com/FictProger/architecture-lab-4/engine"
)

func parse(cmd string) engine.Command {
	fields := string.Fields(cmd)
	name := fields[0]
	args := fields[1:]
	res := commands.Construct(name, args)
	return res
}

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	print := commands.PrintCommand{Arg: "hello"}
	eventLoop.Post(&print)

	add := commands.AddCommand{Arg1: 2, Arg2: 1}
	eventLoop.Post(&add)

	eventLoop.AwaitFinish()
}
