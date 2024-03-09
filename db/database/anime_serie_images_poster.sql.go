// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_serie_images_poster.sql

package db

import (
	"context"
)

const createAnimeSeriePosterImage = `-- name: CreateAnimeSeriePosterImage :one
INSERT INTO anime_serie_poster_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING id, anime_id, image_id, created_at
`

type CreateAnimeSeriePosterImageParams struct {
	AnimeID int64
	ImageID int64
}

func (q *Queries) CreateAnimeSeriePosterImage(ctx context.Context, arg CreateAnimeSeriePosterImageParams) (AnimeSeriePosterImage, error) {
	row := q.db.QueryRow(ctx, createAnimeSeriePosterImage, arg.AnimeID, arg.ImageID)
	var i AnimeSeriePosterImage
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.ImageID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeSeriePosterImage = `-- name: DeleteAnimeSeriePosterImage :exec
DELETE FROM anime_serie_poster_images
WHERE anime_id = $1 AND image_id = $2
`

type DeleteAnimeSeriePosterImageParams struct {
	AnimeID int64
	ImageID int64
}

func (q *Queries) DeleteAnimeSeriePosterImage(ctx context.Context, arg DeleteAnimeSeriePosterImageParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeSeriePosterImage, arg.AnimeID, arg.ImageID)
	return err
}

const listAnimeSeriePosterImages = `-- name: ListAnimeSeriePosterImages :many
SELECT image_id
FROM anime_serie_poster_images
WHERE anime_id = $1
`

func (q *Queries) ListAnimeSeriePosterImages(ctx context.Context, animeID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeSeriePosterImages, animeID)
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
