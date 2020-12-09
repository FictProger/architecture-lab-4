package main

import (
	"github.com/FictProger/architecture-lab-4/commands"
	"github.com/FictProger/architecture-lab-4/engine"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	print := commands.PrintCommand{Arg: "hello"}
	eventLoop.Post(&print)

	eventLoop.AwaitFinish()
}
