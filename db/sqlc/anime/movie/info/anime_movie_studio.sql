-- name: CreateAnimeMovieStudio :one
INSERT INTO anime_movie_studios (anime_id, studio_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieStudio :one
SELECT * FROM anime_movie_studios
WHERE anime_id = $1 AND studio_id = $2;

-- name: ListAnimeMovieStudios :many
SELECT studio_id
FROM anime_movie_studios
WHERE anime_id = $1
ORDER BY id;

-- name: DeleteAnimeMovieStudio :exec
DELETE FROM anime_movie_studios
WHERE anime_id = $1 AND studio_id = $2;