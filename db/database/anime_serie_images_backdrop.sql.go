// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
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

const listAnimeSerieBackdropImages = `-- name: ListAnimeSerieBackdropImages :many
SELECT image_id
FROM anime_serie_backdrop_images
WHERE anime_id = $1
`

func (q *Queries) ListAnimeSerieBackdropImages(ctx context.Context, animeID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeSerieBackdropImages, animeID)
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
