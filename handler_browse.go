package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/TheMaru/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		argLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = argLimit
	}

	postForUserParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), postForUserParams)
	if err != nil {
		return err
	}

	for i, post := range posts {
		fmt.Printf("Post #%d:\n", i+1)
		fmt.Printf("  Title:       %s\n", post.Title)
		fmt.Printf("  Description: %s\n", post.Description.String)
		fmt.Printf("  Published:   %s\n", post.PublishedAt.Time)
		fmt.Printf("  URL:         %s\n", post.Url)
		fmt.Println()
	}

	return nil
}
