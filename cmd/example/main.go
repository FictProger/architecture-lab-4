package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/FictProger/architecture-lab-4/commands"
	"github.com/FictProger/architecture-lab-4/engine"
)

func parse(cmd string) engine.Command {
	fields := strings.Fields(cmd)
	name := fields[0]
	args := fields[1:]
	res := commands.Construct(name, args)
	return res
}

func main() {
	inputFile := flag.String("f", "", "File to run with EventLoop")
	flag.Parse()

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open(*inputFile); err == nil {
		defer input.Close()

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			eventLoop.Post(cmd)
		}
	}

	eventLoop.AwaitFinish()
}
