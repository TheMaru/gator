package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/TheMaru/gator/internal/config"
	"github.com/TheMaru/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("reading of config not possible: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	state := state{
		db:     dbQueries,
		config: &cfg,
	}
	commands := commands{}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerGetUsers)
	commands.register("agg", handlerAgg)

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
