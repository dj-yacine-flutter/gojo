// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_serie_images_backdrop.sql

package db

import (
	"context"
)

const createAnimeSerieBackdropImage = `-- name: CreateAnimeSerieBackdropImage :one
INSERT INTO anime_serie_backdrop_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING id, anime_id, image_id, created_at
`

type CreateAnimeSerieBackdropImageParams struct {
	AnimeID int64
	ImageID int64
}

func (q *Queries) CreateAnimeSerieBackdropImage(ctx context.Context, arg CreateAnimeSerieBackdropImageParams) (AnimeSerieBackdropImage, error) {
	row := q.db.QueryRow(ctx, createAnimeSerieBackdropImage, arg.AnimeID, arg.ImageID)
	var i AnimeSerieBackdropImage
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.ImageID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeSerieBackdropImage = `-- name: DeleteAnimeSerieBackdropImage :exec
DELETE FROM anime_serie_backdrop_images
WHERE anime_id = $1 AND image_id = $2
`

type DeleteAnimeSerieBackdropImageParams struct {
	AnimeID int64
	ImageID int64
}

func (q *Queries) DeleteAnimeSerieBackdropImage(ctx context.Context, arg DeleteAnimeSerieBackdropImageParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeSerieBackdropImage, arg.AnimeID, arg.ImageID)
	return err
}

const getAnimeSerieBackdropImage = `-- name: GetAnimeSerieBackdropImage :one
SELECT id, anime_id, image_id, created_at FROM anime_serie_backdrop_images
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeSerieBackdropImage(ctx context.Context, id int64) (AnimeSerieBackdropImage, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieBackdropImage, id)
	var i AnimeSerieBackdropImage
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.ImageID,
		&i.CreatedAt,
	)
	return i, err
}

const getAnimeSerieBackdropImageByAnimeID = `-- name: GetAnimeSerieBackdropImageByAnimeID :one
SELECT id, anime_id, image_id, created_at FROM anime_serie_backdrop_images
WHERE anime_id = $1 LIMIT 1
`

func (q *Queries) GetAnimeSerieBackdropImageByAnimeID(ctx context.Context, animeID int64) (AnimeSerieBackdropImage, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieBackdropImageByAnimeID, animeID)
	var i AnimeSerieBackdropImage
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.ImageID,
		&i.CreatedAt,
	)
	return i, err
}

const listAnimeSerieBackdropImages = `-- name: ListAnimeSerieBackdropImages :many
SELECT image_id
FROM anime_serie_backdrop_images
WHERE anime_id = $1
LIMIT $2
OFFSET $3
`

type ListAnimeSerieBackdropImagesParams struct {
	AnimeID int64
	Limit   int32
	Offset  int32
}

func (q *Queries) ListAnimeSerieBackdropImages(ctx context.Context, arg ListAnimeSerieBackdropImagesParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeSerieBackdropImages, arg.AnimeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var image_id int64
		if err := rows.Scan(&image_id); err != nil {
			return nil, err
		}
		items = append(items, image_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}