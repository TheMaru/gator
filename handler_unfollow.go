package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/TheMaru/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("url arg is needed for this command")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	delFeedFollowParams := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}

	s.db.DeleteFeedFollow(context.Background(), delFeedFollowParams)
	fmt.Printf("Unfollowed Feed:\n  Name: %s,\n  URL: %s\n", feed.Name, feed.Url)
	return nil
}
