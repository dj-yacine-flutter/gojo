// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: genre.sql

package db

import (
	"context"
)

const createGenre = `-- name: CreateGenre :one
INSERT INTO genres (genre_name)
VALUES ($1)
ON CONFLICT (genre_name)
DO UPDATE SET genre_name = excluded.genre_name
RETURNING  id, genre_name, created_at
`

func (q *Queries) CreateGenre(ctx context.Context, genreName string) (Genre, error) {
	row := q.db.QueryRow(ctx, createGenre, genreName)
	var i Genre
	err := row.Scan(&i.ID, &i.GenreName, &i.CreatedAt)
	return i, err
}

const deleteGenre = `-- name: DeleteGenre :exec
DELETE FROM genres
WHERE id = $1
`

func (q *Queries) DeleteGenre(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteGenre, id)
	return err
}

const getGenre = `-- name: GetGenre :one
SELECT id, genre_name, created_at FROM genres
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGenre(ctx context.Context, id int32) (Genre, error) {
	row := q.db.QueryRow(ctx, getGenre, id)
	var i Genre
	err := row.Scan(&i.ID, &i.GenreName, &i.CreatedAt)
	return i, err
}

const listGenres = `-- name: ListGenres :many
SELECT id FROM genres
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListGenresParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListGenres(ctx context.Context, arg ListGenresParams) ([]int32, error) {
	rows, err := q.db.Query(ctx, listGenres, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGenre = `-- name: UpdateGenre :one
UPDATE genres
SET genre_name = $2
WHERE id = $1
RETURNING id, genre_name, created_at
`

type UpdateGenreParams struct {
	ID        int32
	GenreName string
}

func (q *Queries) UpdateGenre(ctx context.Context, arg UpdateGenreParams) (Genre, error) {
	row := q.db.QueryRow(ctx, updateGenre, arg.ID, arg.GenreName)
	var i Genre
	err := row.Scan(&i.ID, &i.GenreName, &i.CreatedAt)
	return i, err
}
