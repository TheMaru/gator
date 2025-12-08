package rss

import (
	"context"
	"database/sql"
	"time"

	"github.com/TheMaru/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

	const layout = "Mon, 02 Jan 2006 15:04:05 -0700"

	for _, item := range feedData.Channel.Item {
		t, err := time.Parse(layout, item.PubDate)
		if err != nil {
			return err
		}
		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: t, Valid: true},
			FeedID:      markedFeed.ID,
		}

		_, err = db.CreatePost(context.Background(), postParams)
		if err != nil {
			if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
				continue
			}
			return err
		}
	}

	return nil
}
