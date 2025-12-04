package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no arguments given")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("username could not be set: %v", err)
	}

	fmt.Println("user has been set to: " + cmd.args[0])
	return nil
}
