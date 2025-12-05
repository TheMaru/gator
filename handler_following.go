package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	fmt.Printf("User: %s currently follows:\n", s.config.CurrentUserName)
	for _, follow := range follows {
		fmt.Printf("  - %s\n", follow.FeedName)
	}
	return nil
}
