// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: meta.sql

package db

import (
	"context"
)

const createMeta = `-- name: CreateMeta :one
INSERT INTO metas (title, overview)
VALUES ($1, $2)
RETURNING  id, title, overview, created_at
`

type CreateMetaParams struct {
	Title    string
	Overview string
}

func (q *Queries) CreateMeta(ctx context.Context, arg CreateMetaParams) (Meta, error) {
	row := q.db.QueryRow(ctx, createMeta, arg.Title, arg.Overview)
	var i Meta
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.CreatedAt,
	)
	return i, err
}

const deleteMeta = `-- name: DeleteMeta :exec
DELETE FROM Metas
WHERE id = $1
`

func (q *Queries) DeleteMeta(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteMeta, id)
	return err
}

const getMeta = `-- name: GetMeta :one
SELECT id, title, overview, created_at FROM metas
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMeta(ctx context.Context, id int64) (Meta, error) {
	row := q.db.QueryRow(ctx, getMeta, id)
	var i Meta
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.CreatedAt,
	)
	return i, err
}

const updateMeta = `-- name: UpdateMeta :one
UPDATE metas
SET title = $2,
    overview = $3
WHERE id = $1
RETURNING  id, title, overview, created_at
`

type UpdateMetaParams struct {
	ID       int64
	Title    string
	Overview string
}

func (q *Queries) UpdateMeta(ctx context.Context, arg UpdateMetaParams) (Meta, error) {
	row := q.db.QueryRow(ctx, updateMeta, arg.ID, arg.Title, arg.Overview)
	var i Meta
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.CreatedAt,
	)
	return i, err
}
