-- name: CreateUser :one
insert into users (id, created_at, updated_at, name, api_key)
values ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
returning *;

-- name: GetByAPIKey :one
select * from users where api_key = $1;

-- name: GetFeeds :many
select * from feeds;

-- name: TruncateUsers :exec
truncate table users cascade;