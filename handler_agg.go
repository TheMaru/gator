package main

import (
	"context"
	"fmt"

	"github.com/TheMaru/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	tmpFeedDest := "https://www.wagslane.dev/index.xml"

	rssFeed, err := rss.FetchFeed(context.Background(), tmpFeedDest)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", rssFeed)
	return nil
}
