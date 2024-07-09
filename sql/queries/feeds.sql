-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, name, url, user_id)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetFeedByUser :many
select id, created_at, updated_at, name, url, user_id from feeds where user_id = $1;

-- name: CreateFeedFollow :one
insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
values ($1, $2, $3, $4, $5)
returning *;

-- name: DeleteFeedFollow :exec
delete from feed_follows where id = $1 and user_id = $2;

-- name: GetNextFeedsToFetch :many
select id, created_at, updated_at, name, url, user_id
from feeds
order by last_fetched_at asc nulls first
limit $1;

-- name: MarkFeedAsFetched :exec
update feeds set last_fetched_at = now(), updated_at = now() where id = $1 returning *;

-- name: TruncateFeeds :exec
truncate table feeds cascade;