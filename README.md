# gator

rss feed aggregator to test SQL database connections and handling

## Prerequisites

In order to save data you'll need to install PostgreSQL first. Use your
favorite package manager to install PostgreSQL and add a database.

Since this is a Go programm if you want to build it or interact with the code
you'll need Go of course as well. So please install it, you can find
instructions at [Go install documentation](https://go.dev/doc/install).

### Installation via go install

In Go you can compile and install packages via the command `go install
[packages]` so to install this module you can use `go install
github.com/TheMaru/gator@latest` in order to install the latest version.

### Database migrations

To setup the databse correctly

## Config file

The program expects a config file in your Home directory with the name
`.gatorconfig.json` so please create the file and add the URL to your database

```json
{
    "db_url": "connection_string"
}
```

### Databse migrations

I've used goose in order to run database migrations, so please install it
if you don't have it already.
`go install github.com/pressly/goose/v3/cmd/goose`
cd into the sql/schema folder and run the goose up command:
```goose postgres <connection_string> up```

### How to run gator

Gator is a simple command line tool.
I've made my live easy and only implemented multiple users without a secure
login/password mechanism.

- Use the command `gator register <name>` to register a new user
- Use the command `gator login <name>` to login as the user.
- You can get a list of all currently registered Users via `gator users`
- add a rss feed via `gator addfeed <URL>`
- list feeds via `gator feeds`
- `gator follow <URL>` and `gator unfollow <URL>` will be used to
follow/unfollow a feed for the current user
- `gator following` lists all currently followed feeds
- `gator agg <duration>` This will start the aggregation of feeds. After
\<duration\>(a go duration string) the oldest feed will be updated and so on
until you stop the command
- `gator browse <amount>` lists the newest amount of feeds a user follwed

