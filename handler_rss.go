package main

import (
	"context"
	"fmt"
)

// this will be a commandfunc
func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}

	fmt.Println(feed)
	return nil
}
