// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_serie_season.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeSeason = `-- name: CreateAnimeSeason :one
INSERT INTO anime_serie_seasons (
    anime_id,
    season_original_title,
    release_year,
    aired,
    rating,
    portrait_poster,
    portrait_blur_hash
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, anime_id, season_original_title, release_year, aired, portrait_poster, portrait_blur_hash, rating, created_at
`

type CreateAnimeSeasonParams struct {
	AnimeID             int64
	SeasonOriginalTitle string
	ReleaseYear         int32
	Aired               time.Time
	Rating              string
	PortraitPoster      string
	PortraitBlurHash    string
}

func (q *Queries) CreateAnimeSeason(ctx context.Context, arg CreateAnimeSeasonParams) (AnimeSerieSeason, error) {
	row := q.db.QueryRow(ctx, createAnimeSeason,
		arg.AnimeID,
		arg.SeasonOriginalTitle,
		arg.ReleaseYear,
		arg.Aired,
		arg.Rating,
		arg.PortraitPoster,
		arg.PortraitBlurHash,
	)
	var i AnimeSerieSeason
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.SeasonOriginalTitle,
		&i.ReleaseYear,
		&i.Aired,
		&i.PortraitPoster,
		&i.PortraitBlurHash,
		&i.Rating,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeSeason = `-- name: DeleteAnimeSeason :exec
DELETE FROM anime_serie_seasons
WHERE id = $1
`

func (q *Queries) DeleteAnimeSeason(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeSeason, id)
	return err
}

const getAnimeSeason = `-- name: GetAnimeSeason :one
SELECT id, anime_id, season_original_title, release_year, aired, portrait_poster, portrait_blur_hash, rating, created_at FROM anime_serie_seasons
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAnimeSeason(ctx context.Context, id int64) (AnimeSerieSeason, error) {
	row := q.db.QueryRow(ctx, getAnimeSeason, id)
	var i AnimeSerieSeason
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.SeasonOriginalTitle,
		&i.ReleaseYear,
		&i.Aired,
		&i.PortraitPoster,
		&i.PortraitBlurHash,
		&i.Rating,
		&i.CreatedAt,
	)
	return i, err
}

const listAnimeSeasonsByAnimeID = `-- name: ListAnimeSeasonsByAnimeID :many
SELECT id FROM anime_serie_seasons
WHERE anime_id = $1
ORDER BY release_year
LIMIT $2
OFFSET $3
`

type ListAnimeSeasonsByAnimeIDParams struct {
	AnimeID int64
	Limit   int32
	Offset  int32
}

func (q *Queries) ListAnimeSeasonsByAnimeID(ctx context.Context, arg ListAnimeSeasonsByAnimeIDParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeSeasonsByAnimeID, arg.AnimeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
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

const updateAnimeSeason = `-- name: UpdateAnimeSeason :one
UPDATE anime_serie_seasons
SET
  season_original_title = COALESCE($1, season_original_title),
  release_year = COALESCE($2, release_year),
  aired = COALESCE($3, aired),
  rating = COALESCE($4, rating),
  portrait_poster = COALESCE($5, portrait_poster),
  portrait_blur_hash = COALESCE($6, portrait_blur_hash)
WHERE
  id = $7
RETURNING id, anime_id, season_original_title, release_year, aired, portrait_poster, portrait_blur_hash, rating, created_at
`

type UpdateAnimeSeasonParams struct {
	SeasonOriginalTitle pgtype.Text
	ReleaseYear         pgtype.Int4
	Aired               pgtype.Timestamptz
	Rating              pgtype.Text
	PortraitPoster      pgtype.Text
	PortraitBlurHash    pgtype.Text
	ID                  int64
}

func (q *Queries) UpdateAnimeSeason(ctx context.Context, arg UpdateAnimeSeasonParams) (AnimeSerieSeason, error) {
	row := q.db.QueryRow(ctx, updateAnimeSeason,
		arg.SeasonOriginalTitle,
		arg.ReleaseYear,
		arg.Aired,
		arg.Rating,
		arg.PortraitPoster,
		arg.PortraitBlurHash,
		arg.ID,
	)
	var i AnimeSerieSeason
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.SeasonOriginalTitle,
		&i.ReleaseYear,
		&i.Aired,
		&i.PortraitPoster,
		&i.PortraitBlurHash,
		&i.Rating,
		&i.CreatedAt,
	)
	return i, err
}
