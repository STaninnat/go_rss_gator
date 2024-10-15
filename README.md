# Gator

Golang RSS feed (CLI tool)

A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Features

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Installation and Requirements

- Clone the repository.

Before running the program, ensure you have the following installed:

- **[Go](https://golang.org/dl/)** (Make sure you have the latest): The primary language for building the project.
- **[PostgreSQL](https://www.postgresql.org/download/)** (Make sure you have the latest): Relational database used to store aggregated blog data.
- **SQLc**: Automatically generates Go code from SQL queries.
- **Goose**: Manages database migrations.
- **pgAdmin**: Database management tool to interact with PostgreSQL.

## Config

1. Run `chmod +x install.sh` and then `./install.sh`
2. Set up PostgreSQL and manually create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://<username>:<password>@<host>:<port>/<dbname>?<options>"
}
```
  *example:  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"*<br>
3. You'll need to run `go build -o gator` to generate the binary before you can run `gator <command>`.
4. Run the necessary database migrations with Goose.

## Usage

Help command to see all commands:
```bash
gator help
```

Create a new user:

```bash
gator register <name>
```

Add a feed:

```bash
gator addfeed <url>
```

Start the aggregator:

```bash
gator agg 30s
```

View the posts:

```bash
gator browse [limit]
```

There are a few other commands you'll need as well:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database

## Running the Gator Command

If you encounter an issue where the `gator` command is not recognized, you can run the binary directly using the following command:

```bash
./gator <command>
```