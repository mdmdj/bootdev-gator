package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mdmdj/bootdev-gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	dbUser, err := s.db.GetUserByName(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	err = s.cfg.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	if name == "" {
		return fmt.Errorf("username cannot be empty")
	}

	dbUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), Name: name})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User registered successfully!", dbUser.ID, dbUser.Name)

	return nil
}

func handlerResetUsers(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("couldn't reset users: %w", err)
	}

	err = s.cfg.SetUser("")
	if err != nil {
		return fmt.Errorf("error resetting users, couldn't set current user: %w", err)
	}

	fmt.Println("All users have been reset!")
	return nil
}

func handlerUsers(s *state, _ command) error {
	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}

	for _, user := range dbUsers {
		if s.cfg.CurrentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}

	}

	return nil
}
