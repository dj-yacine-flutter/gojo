// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_media.sql

package db

import (
	"context"
)

const createAnimeMovieMedia = `-- name: CreateAnimeMovieMedia :one
INSERT INTO anime_movie_media (anime_id, media_id)
VALUES ($1, $2)
RETURNING id, anime_id, media_id, created_at
`

type CreateAnimeMovieMediaParams struct {
	AnimeID int64
	MediaID int64
}

func (q *Queries) CreateAnimeMovieMedia(ctx context.Context, arg CreateAnimeMovieMediaParams) (AnimeMovieMedium, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieMedia, arg.AnimeID, arg.MediaID)
	var i AnimeMovieMedium
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.MediaID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeMovieMedia = `-- name: DeleteAnimeMovieMedia :exec
DELETE FROM anime_movie_media
WHERE anime_id = $1 AND media_id = $2
`

type DeleteAnimeMovieMediaParams struct {
	AnimeID int64
	MediaID int64
}

func (q *Queries) DeleteAnimeMovieMedia(ctx context.Context, arg DeleteAnimeMovieMediaParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieMedia, arg.AnimeID, arg.MediaID)
	return err
}

const getAnimeMovieMedia = `-- name: GetAnimeMovieMedia :one
SELECT id, anime_id, media_id, created_at FROM anime_movie_media
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeMovieMedia(ctx context.Context, id int64) (AnimeMovieMedium, error) {
	row := q.db.QueryRow(ctx, getAnimeMovieMedia, id)
	var i AnimeMovieMedium
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.MediaID,
		&i.CreatedAt,
	)
	return i, err
}

const getAnimeMovieMediaByAnimeID = `-- name: GetAnimeMovieMediaByAnimeID :one
SELECT id, anime_id, media_id, created_at FROM anime_movie_media
WHERE anime_id = $1 LIMIT 1
`

func (q *Queries) GetAnimeMovieMediaByAnimeID(ctx context.Context, animeID int64) (AnimeMovieMedium, error) {
	row := q.db.QueryRow(ctx, getAnimeMovieMediaByAnimeID, animeID)
	var i AnimeMovieMedium
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.MediaID,
		&i.CreatedAt,
	)
	return i, err
}

const listAnimeMovieMedias = `-- name: ListAnimeMovieMedias :many
SELECT media_id
FROM anime_movie_media
WHERE anime_id = $1
LIMIT $2
OFFSET $3
`

type ListAnimeMovieMediasParams struct {
	AnimeID int64
	Limit   int32
	Offset  int32
}

func (q *Queries) ListAnimeMovieMedias(ctx context.Context, arg ListAnimeMovieMediasParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeMovieMedias, arg.AnimeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var media_id int64
		if err := rows.Scan(&media_id); err != nil {
			return nil, err
		}
		items = append(items, media_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
