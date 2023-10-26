// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_serie_studio.sql

package db

import (
	"context"
)

const createAnimeSerieStudio = `-- name: CreateAnimeSerieStudio :one
INSERT INTO anime_serie_studios (anime_id, studio_id)
VALUES ($1, $2)
RETURNING id, anime_id, studio_id
`

type CreateAnimeSerieStudioParams struct {
	AnimeID  int64 `json:"anime_id"`
	StudioID int32 `json:"studio_id"`
}

func (q *Queries) CreateAnimeSerieStudio(ctx context.Context, arg CreateAnimeSerieStudioParams) (AnimeSerieStudio, error) {
	row := q.db.QueryRow(ctx, createAnimeSerieStudio, arg.AnimeID, arg.StudioID)
	var i AnimeSerieStudio
	err := row.Scan(&i.ID, &i.AnimeID, &i.StudioID)
	return i, err
}

const deleteAnimeSerieStudio = `-- name: DeleteAnimeSerieStudio :exec
DELETE FROM anime_serie_studios
WHERE anime_id = $1 AND studio_id = $2
`

type DeleteAnimeSerieStudioParams struct {
	AnimeID  int64 `json:"anime_id"`
	StudioID int32 `json:"studio_id"`
}

func (q *Queries) DeleteAnimeSerieStudio(ctx context.Context, arg DeleteAnimeSerieStudioParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeSerieStudio, arg.AnimeID, arg.StudioID)
	return err
}

const getAnimeSerieStudio = `-- name: GetAnimeSerieStudio :one
SELECT id, anime_id, studio_id FROM anime_serie_studios
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeSerieStudio(ctx context.Context, id int64) (AnimeSerieStudio, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieStudio, id)
	var i AnimeSerieStudio
	err := row.Scan(&i.ID, &i.AnimeID, &i.StudioID)
	return i, err
}

const listAnimeSerieStudios = `-- name: ListAnimeSerieStudios :many
SELECT studio_id
FROM anime_serie_studios
WHERE anime_id = $1
LIMIT $2
OFFSET $3
`

type ListAnimeSerieStudiosParams struct {
	AnimeID int64 `json:"anime_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListAnimeSerieStudios(ctx context.Context, arg ListAnimeSerieStudiosParams) ([]int32, error) {
	rows, err := q.db.Query(ctx, listAnimeSerieStudios, arg.AnimeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var studio_id int32
		if err := rows.Scan(&studio_id); err != nil {
			return nil, err
		}
		items = append(items, studio_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
