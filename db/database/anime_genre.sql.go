// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: anime_genre.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeGenre = `-- name: CreateAnimeGenre :exec
INSERT INTO anime_genre (anime_id, genre_id)
VALUES ($1, $2)
`

type CreateAnimeGenreParams struct {
	AnimeID int64       `json:"anime_id"`
	GenreID pgtype.Int4 `json:"genre_id"`
}

func (q *Queries) CreateAnimeGenre(ctx context.Context, arg CreateAnimeGenreParams) error {
	_, err := q.db.Exec(ctx, createAnimeGenre, arg.AnimeID, arg.GenreID)
	return err
}

const deleteAnimeGenre = `-- name: DeleteAnimeGenre :exec
DELETE FROM anime_genre
WHERE anime_id = $1 AND genre_id = $2
`

type DeleteAnimeGenreParams struct {
	AnimeID int64       `json:"anime_id"`
	GenreID pgtype.Int4 `json:"genre_id"`
}

func (q *Queries) DeleteAnimeGenre(ctx context.Context, arg DeleteAnimeGenreParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeGenre, arg.AnimeID, arg.GenreID)
	return err
}

const getAnimeGenre = `-- name: GetAnimeGenre :one
SELECT id, anime_id, genre_id FROM anime_genre
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeGenre(ctx context.Context, id int64) (AnimeGenre, error) {
	row := q.db.QueryRow(ctx, getAnimeGenre, id)
	var i AnimeGenre
	err := row.Scan(&i.ID, &i.AnimeID, &i.GenreID)
	return i, err
}

const listAnimeGenres = `-- name: ListAnimeGenres :many
SELECT genre_id
FROM anime_genre
WHERE anime_id = $1
LIMIT $2
OFFSET $3
`

type ListAnimeGenresParams struct {
	AnimeID int64 `json:"anime_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListAnimeGenres(ctx context.Context, arg ListAnimeGenresParams) ([]pgtype.Int4, error) {
	rows, err := q.db.Query(ctx, listAnimeGenres, arg.AnimeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []pgtype.Int4{}
	for rows.Next() {
		var genre_id pgtype.Int4
		if err := rows.Scan(&genre_id); err != nil {
			return nil, err
		}
		items = append(items, genre_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAnimeGenre = `-- name: UpdateAnimeGenre :one
UPDATE anime_genre
SET genre_id = $2
WHERE anime_id = $1
RETURNING id, anime_id, genre_id
`

type UpdateAnimeGenreParams struct {
	AnimeID int64       `json:"anime_id"`
	GenreID pgtype.Int4 `json:"genre_id"`
}

func (q *Queries) UpdateAnimeGenre(ctx context.Context, arg UpdateAnimeGenreParams) (AnimeGenre, error) {
	row := q.db.QueryRow(ctx, updateAnimeGenre, arg.AnimeID, arg.GenreID)
	var i AnimeGenre
	err := row.Scan(&i.ID, &i.AnimeID, &i.GenreID)
	return i, err
}
