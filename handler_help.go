package main

import "fmt"

type clicommand struct {
	name        string
	description string
	usage       string
}

func getCommands() map[string]clicommand {
	return map[string]clicommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			usage:       "./gator help",
		},
		"login": {
			name:        "login",
			description: "Sets the current user in the config",
			usage:       "./gator login <username>",
		},
		"register": {
			name:        "register",
			description: "Adds a new user to the database",
			usage:       "./gator resigter <username>",
		},
		"users": {
			name:        "users",
			description: "lists all the users in the database",
			usage:       "./gator users",
		},
		"reset": {
			name:        "reset",
			description: "Manually reset the state of the database",
			usage:       "./gator reset",
		},
		"agg": {
			name:        "aggregator",
			description: "Fetch the RSS feeds. This command takes a duration as an argument and uses it to set up a ticker that defines the interval for repeatedly fetching data from the feed. <time_between_reqs> like 1s, 1m, 1h",
			usage:       "./gator agg <time_between_reqs>",
		},
		"addfeed": {
			name:        "add feed",
			description: "Connect the feed to the current user",
			usage:       "./gator addfeed <name_of_the_feed> <url>",
		},
		"feeds": {
			name:        "feeds",
			description: "Inspect all the feeds in the database",
			usage:       "./gator feeds",
		},
		"follow": {
			name:        "follow",
			description: "Create a new feed follow record for the current user",
			usage:       "./gator follow <url>",
		},
		"following": {
			name:        "following",
			description: "Print all the names of the feeds the current user follows",
			usage:       "./gator following",
		},
		"unfollow": {
			name:        "unfollow",
			description: "Unfollow the feeds the cuurent user is following",
			usage:       "./gator unfollow <url>",
		},
		"browse": {
			name:        "browse",
			description: "View all the posts from the feeds the user follows. This command take an optional limit parameter(It is the number of posts to pull). if it's not provided, default the limit to 2",
			usage:       "./gator browse <limit>",
		},
	}
}

func handlerHelp(s *state, cmd command) error {
	fmt.Println("==================================================================================")
	fmt.Println()
	fmt.Println("Welcome to RSS Aggregator")
	fmt.Println("Usage:")
	fmt.Println()
	for _, clicmd := range getCommands() {
		fmt.Printf("	%s: %s\n", clicmd.name, clicmd.description)
		fmt.Printf("	usage: %s\n\n", clicmd.usage)
	}

	fmt.Println("==================================================================================")
	return nil
}
