// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_movie_images_logo.sql

package db

import (
	"context"
)

const createAnimeMovieLogoImage = `-- name: CreateAnimeMovieLogoImage :one
INSERT INTO anime_movie_logo_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING id, anime_id, image_id, created_at
`

type CreateAnimeMovieLogoImageParams struct {
	AnimeID int64
	ImageID int64
}

func (q *Queries) CreateAnimeMovieLogoImage(ctx context.Context, arg CreateAnimeMovieLogoImageParams) (AnimeMovieLogoImage, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieLogoImage, arg.AnimeID, arg.ImageID)
	var i AnimeMovieLogoImage
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.ImageID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeMovieLogoImage = `-- name: DeleteAnimeMovieLogoImage :exec
DELETE FROM anime_movie_logo_images
WHERE anime_id = $1 AND image_id = $2
`

type DeleteAnimeMovieLogoImageParams struct {
	AnimeID int64
	ImageID int64
}

func (q *Queries) DeleteAnimeMovieLogoImage(ctx context.Context, arg DeleteAnimeMovieLogoImageParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieLogoImage, arg.AnimeID, arg.ImageID)
	return err
}

const listAnimeMovieLogoImages = `-- name: ListAnimeMovieLogoImages :many
SELECT image_id
FROM anime_movie_logo_images
WHERE anime_id = $1
`

func (q *Queries) ListAnimeMovieLogoImages(ctx context.Context, animeID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeMovieLogoImages, animeID)
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
