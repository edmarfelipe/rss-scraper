// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, name, url, user_id)
values ($1, $2, $3, $4, $5, $6)
returning id, created_at, updated_at, name, url, user_id, last_fetched_at
`

type CreateFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const createFeedFollow = `-- name: CreateFeedFollow :one
insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
values ($1, $2, $3, $4, $5)
returning id, created_at, updated_at, user_id, feed_id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
delete from feed_follows where id = $1 and user_id = $2
`

type DeleteFeedFollowParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.ID, arg.UserID)
	return err
}

const getFeedByUser = `-- name: GetFeedByUser :many
select id, created_at, updated_at, name, url, user_id from feeds where user_id = $1
`

type GetFeedByUserRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) GetFeedByUser(ctx context.Context, userID uuid.UUID) ([]GetFeedByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedByUserRow
	for rows.Next() {
		var i GetFeedByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNextFeedsToFetch = `-- name: GetNextFeedsToFetch :many
select id, created_at, updated_at, name, url, user_id
from feeds
order by last_fetched_at asc nulls first
limit $1
`

type GetNextFeedsToFetchRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) GetNextFeedsToFetch(ctx context.Context, limit int32) ([]GetNextFeedsToFetchRow, error) {
	rows, err := q.db.QueryContext(ctx, getNextFeedsToFetch, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNextFeedsToFetchRow
	for rows.Next() {
		var i GetNextFeedsToFetchRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markFeedAsFetched = `-- name: MarkFeedAsFetched :exec
update feeds set last_fetched_at = now(), updated_at = now() where id = $1 returning id, created_at, updated_at, name, url, user_id, last_fetched_at
`

func (q *Queries) MarkFeedAsFetched(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, markFeedAsFetched, id)
	return err
}

const truncateFeeds = `-- name: TruncateFeeds :exec
truncate table feeds cascade
`

func (q *Queries) TruncateFeeds(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, truncateFeeds)
	return err
}
