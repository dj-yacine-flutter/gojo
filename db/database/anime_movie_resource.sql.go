// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_resource.sql

package db

import (
	"context"
)

const createAnimeMovieResource = `-- name: CreateAnimeMovieResource :one
INSERT INTO anime_movie_resources (anime_id, resource_id)
VALUES ($1, $2)
RETURNING id, anime_id, resource_id
`

type CreateAnimeMovieResourceParams struct {
	AnimeID    int64
	ResourceID int64
}

func (q *Queries) CreateAnimeMovieResource(ctx context.Context, arg CreateAnimeMovieResourceParams) (AnimeMovieResource, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieResource, arg.AnimeID, arg.ResourceID)
	var i AnimeMovieResource
	err := row.Scan(&i.ID, &i.AnimeID, &i.ResourceID)
	return i, err
}

const deleteAnimeMovieResource = `-- name: DeleteAnimeMovieResource :exec
DELETE FROM anime_movie_resources
WHERE anime_id = $1 AND resource_id = $2
`

type DeleteAnimeMovieResourceParams struct {
	AnimeID    int64
	ResourceID int64
}

func (q *Queries) DeleteAnimeMovieResource(ctx context.Context, arg DeleteAnimeMovieResourceParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieResource, arg.AnimeID, arg.ResourceID)
	return err
}

const getAnimeMovieResource = `-- name: GetAnimeMovieResource :one
SELECT id, anime_id, resource_id FROM anime_movie_resources
WHERE anime_id = $1
LIMIT 1
`

func (q *Queries) GetAnimeMovieResource(ctx context.Context, animeID int64) (AnimeMovieResource, error) {
	row := q.db.QueryRow(ctx, getAnimeMovieResource, animeID)
	var i AnimeMovieResource
	err := row.Scan(&i.ID, &i.AnimeID, &i.ResourceID)
	return i, err
}
