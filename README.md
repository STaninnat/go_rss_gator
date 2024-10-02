# Blog Aggregator CLI

Golang RSS feed (CLI tool)

This project is a Go application that interacts with a PostgreSQL database.

## Features

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Requirements

Before running the program, ensure you have the following installed:

- **[Go](https://golang.org/dl/)** (version 1.19 or above): The primary language for building the project.
- **[PostgreSQL](https://www.postgresql.org/download/)** (version 12 or above): Relational database used to store aggregated blog data.
- **SQLc**: Automatically generates Go code from SQL queries.
- **Goose**: Manages database migrations.
- **pgAdmin**: Database management tool to interact with PostgreSQL.

## Installation

1. Clone the repository.
2. Set up PostgreSQL and Manually create a config file in your home directory, `~/.gatorconfig.json`, with the following content: 
`{"db_url": "postgres://example"}` or change path of `.gatorconfig.json` in `internal/config/config.go/getConfigFilePath`.
3. Run the necessary database migrations with Goose.
4. Run the application by using the following command: `./gator <command>`.
    - If you're using a different system, you'll need to run `go build -o gator` to generate the binary before you can run `./gator <command>`.
    - You can use `./gator help` to see all commands.

