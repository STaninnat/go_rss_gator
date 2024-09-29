package main

import "fmt"

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchfeed("https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feef: %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
