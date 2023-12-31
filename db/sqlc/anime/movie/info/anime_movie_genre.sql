-- name: CreateAnimeMovieGenre :one
INSERT INTO anime_movie_genres (anime_id, genre_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieGenre :one
SELECT * FROM anime_movie_genres
WHERE anime_id = $1 AND genre_id = $2;

-- name: ListAnimeMovieGenres :many
SELECT genre_id
FROM anime_movie_genres
WHERE anime_id = $1
ORDER BY id;

-- name: DeleteAnimeMovieGenre :exec
DELETE FROM anime_movie_genres
WHERE anime_id = $1 AND genre_id = $2;