package main

import (
	"context"
	"fmt"

	"github.com/TheMaru/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("User: %s currently follows:\n", user.Name)
	for _, follow := range follows {
		fmt.Printf("  - %s\n", follow.FeedName)
	}
	return nil
}
