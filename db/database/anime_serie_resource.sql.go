// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_serie_resource.sql

package db

import (
	"context"
)

const createAnimeSerieResource = `-- name: CreateAnimeSerieResource :one
INSERT INTO anime_serie_resources (anime_id, resource_id)
VALUES ($1, $2)
RETURNING id, anime_id, resource_id
`

type CreateAnimeSerieResourceParams struct {
	AnimeID    int64 `json:"anime_id"`
	ResourceID int64 `json:"resource_id"`
}

func (q *Queries) CreateAnimeSerieResource(ctx context.Context, arg CreateAnimeSerieResourceParams) (AnimeSerieResource, error) {
	row := q.db.QueryRow(ctx, createAnimeSerieResource, arg.AnimeID, arg.ResourceID)
	var i AnimeSerieResource
	err := row.Scan(&i.ID, &i.AnimeID, &i.ResourceID)
	return i, err
}

const deleteAnimeSerieResource = `-- name: DeleteAnimeSerieResource :exec
DELETE FROM anime_serie_resources
WHERE anime_id = $1 AND resource_id = $2
`

type DeleteAnimeSerieResourceParams struct {
	AnimeID    int64 `json:"anime_id"`
	ResourceID int64 `json:"resource_id"`
}

func (q *Queries) DeleteAnimeSerieResource(ctx context.Context, arg DeleteAnimeSerieResourceParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeSerieResource, arg.AnimeID, arg.ResourceID)
	return err
}

const getAnimeSerieResource = `-- name: GetAnimeSerieResource :one
SELECT id, anime_id, resource_id FROM anime_serie_resources
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeSerieResource(ctx context.Context, id int64) (AnimeSerieResource, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieResource, id)
	var i AnimeSerieResource
	err := row.Scan(&i.ID, &i.AnimeID, &i.ResourceID)
	return i, err
}

const listAnimeSerieResources = `-- name: ListAnimeSerieResources :many
SELECT resource_id
FROM anime_serie_resources
WHERE anime_id = $1
LIMIT $2
OFFSET $3
`

type ListAnimeSerieResourcesParams struct {
	AnimeID int64 `json:"anime_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListAnimeSerieResources(ctx context.Context, arg ListAnimeSerieResourcesParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeSerieResources, arg.AnimeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var resource_id int64
		if err := rows.Scan(&resource_id); err != nil {
			return nil, err
		}
		items = append(items, resource_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
