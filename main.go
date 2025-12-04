package main

import (
	"log"
	"os"

	"github.com/TheMaru/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("reading of config not possible: %v", err)
	}
	state := state{&cfg}
	commands := commands{}
	commands.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("no command given, please provide a command!")
	}

	cmdName := args[1]
	arguments := args[2:]
	cmdToRun := command{
		name: cmdName,
		args: arguments,
	}

	err = commands.run(&state, cmdToRun)
	if err != nil {
		log.Fatal(err)
	}
}
