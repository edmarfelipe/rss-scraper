-- name: CreatePost :one
insert into posts (id, created_at, updated_at, title, url, content, published_at, feed_id)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning *;


-- name: TruncatePosts :exec
truncate table posts cascade;

-- name: GetPostsForUser :many
select
    posts.id,
    posts.created_at,
    posts.updated_at,
    posts.title,
    posts.url,
    posts.content,
    posts.published_at,
    posts.feed_id
from posts
join feed_follows on posts.feed_id = feed_follows.feed_id
where feed_follows.user_id = $1
order by posts.published_at desc
limit $2 OFFSET $3;
