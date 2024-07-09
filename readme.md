# RSS Scraper

This is a simple RSS scraper that fetches RSS feeds from a list of URLs and stores them in a database.

## Features

* A user can create an account
* A user can create a feed
* A user can follow a feed
* A user can unfollow a feed
* A user can view all feeds they follow
* A user can view all posts from feeds they follow


## Libraries used

* [Sqlc](https://sqlc.dev/) - For type-safe SQL queries
* [ogen](https://ogen.dev/) - For generating Go code from OpenAPI specs
* [Goose](https://pressly.github.io/goose/) - For managing migrations

## Database Schema

```mermaid
erDiagram
    users {
       uuid Id PK
       text name
       varchar(64) api_key
       timestamp created_at
       timestamp update_at
    }
    users ||--|{ feed_follows : has
    users ||--|{ feeds : has
    feeds {
        uuid Id PK
        text name
        text url
        uuid user_id
        timestamp created_at
        timestamp update_at
    }
    feeds ||--|{ posts : has
    feeds ||--|{ feed_follows : has
    feed_follows {
        uuid Id PK
        uuid user_id
        uuid feed_id
        timestamp created_at
        timestamp update_at
    }
    posts {
        uuid Id PK
        string title
        string url
        text context
        timestamp published_at
        timestamp created_at
        timestamp update_at
    }
```


## How to run

### Install dependencies

```sh
go mod download
```

### Generate code from OpenAPI spec

```sh
go generate ./...
```

### Run the server

```sh
go run cmd/main.go
```

### Run migrations

```sh
go run cmd/migrate.go
```

### Run tests

```sh
go test ./...
```


