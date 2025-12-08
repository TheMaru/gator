package rss

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/TheMaru/gator/internal/database"
)

func ScrapeFeeds(db *database.Queries) error {
	oldestFeed, err := db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	markFeedParams := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now()},
		UpdatedAt:     time.Now(),
		ID:            oldestFeed.ID,
	}

	markedFeed, err := db.MarkFeedFetched(context.Background(), markFeedParams)
	if err != nil {
		return err
	}

	feedData, err := FetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range feedData.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
