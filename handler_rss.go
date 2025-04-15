package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mdmdj/bootdev-gator/internal/database"
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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <feedname> <url>")
	}

	dbUser, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}

	feedName := cmd.Args[0]
	url := cmd.Args[1]
	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   feedName,
		Url:    url,
		UserID: dbUser.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	fmt.Println(newFeed)
	return nil

}

// get all feeds from database and output to console
func handlerFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}
	return nil
}
