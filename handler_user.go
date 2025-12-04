package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no arguments given")
	}

	name := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("username could not be set: %v", err)
	}

	fmt.Println("user has been set to: " + cmd.args[0])
	return nil
}
