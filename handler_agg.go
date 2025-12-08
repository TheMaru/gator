package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/TheMaru/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("duration string as argument for request timeout needed")
	}
	duratiom, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s\n", duratiom.Abs().String())

	ticker := time.NewTicker(duratiom)
	for ; ; <-ticker.C {
		err = rss.ScrapeFeeds(s.db)
		if err != nil {
			return err
		}
	}

	return nil
}
