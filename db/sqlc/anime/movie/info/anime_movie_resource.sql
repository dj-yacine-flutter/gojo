-- name: CreateAnimeMovieResource :one
INSERT INTO anime_movie_resources (anime_id, resource_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieResource :one
SELECT * FROM anime_movie_resources
WHERE anime_id = $1
LIMIT 1;

-- name: DeleteAnimeMovieResource :exec
DELETE FROM anime_movie_resources
WHERE anime_id = $1 AND resource_id = $2;