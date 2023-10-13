// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: anime_movie.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeMovie = `-- name: CreateAnimeMovie :one
INSERT INTO anime_movie (
    original_title,
    aired,
    release_year,
    duration
)
VALUES ($1, $2, $3, $4)
RETURNING id, original_title, aired, release_year, duration, created_at
`

type CreateAnimeMovieParams struct {
	OriginalTitle string    `json:"original_title"`
	Aired         time.Time `json:"aired"`
	ReleaseYear   int32     `json:"release_year"`
	Duration      int32     `json:"duration"`
}

func (q *Queries) CreateAnimeMovie(ctx context.Context, arg CreateAnimeMovieParams) (AnimeMovie, error) {
	row := q.db.QueryRow(ctx, createAnimeMovie,
		arg.OriginalTitle,
		arg.Aired,
		arg.ReleaseYear,
		arg.Duration,
	)
	var i AnimeMovie
	err := row.Scan(
		&i.ID,
		&i.OriginalTitle,
		&i.Aired,
		&i.ReleaseYear,
		&i.Duration,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeMovie = `-- name: DeleteAnimeMovie :exec
DELETE FROM anime_movie 
WHERE id = $1
`

func (q *Queries) DeleteAnimeMovie(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovie, id)
	return err
}

const getAnimeMovie = `-- name: GetAnimeMovie :one
SELECT id, original_title, aired, release_year, duration, created_at FROM anime_movie 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeMovie(ctx context.Context, id int64) (AnimeMovie, error) {
	row := q.db.QueryRow(ctx, getAnimeMovie, id)
	var i AnimeMovie
	err := row.Scan(
		&i.ID,
		&i.OriginalTitle,
		&i.Aired,
		&i.ReleaseYear,
		&i.Duration,
		&i.CreatedAt,
	)
	return i, err
}

const updateAnimeMovie = `-- name: UpdateAnimeMovie :one
UPDATE anime_movie
SET
  original_title = COALESCE($1, original_title),
  aired = COALESCE($2, aired),
  release_year = COALESCE($3, release_year),
  duration = COALESCE($4, duration)
WHERE
  id = $5
RETURNING id, original_title, aired, release_year, duration, created_at
`

type UpdateAnimeMovieParams struct {
	OriginalTitle pgtype.Text        `json:"original_title"`
	Aired         pgtype.Timestamptz `json:"aired"`
	ReleaseYear   pgtype.Int4        `json:"release_year"`
	Duration      pgtype.Int4        `json:"duration"`
	ID            int64              `json:"id"`
}

func (q *Queries) UpdateAnimeMovie(ctx context.Context, arg UpdateAnimeMovieParams) (AnimeMovie, error) {
	row := q.db.QueryRow(ctx, updateAnimeMovie,
		arg.OriginalTitle,
		arg.Aired,
		arg.ReleaseYear,
		arg.Duration,
		arg.ID,
	)
	var i AnimeMovie
	err := row.Scan(
		&i.ID,
		&i.OriginalTitle,
		&i.Aired,
		&i.ReleaseYear,
		&i.Duration,
		&i.CreatedAt,
	)
	return i, err
}
