// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_serie_episode.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeSerieEpisode = `-- name: CreateAnimeSerieEpisode :one
INSERT INTO anime_serie_episodes (
  episode_number,
  season_id,
  thumbnails,
  thumbnails_blur_hash
)
VALUES ($1, $2, $3, $4)
RETURNING id, episode_number, season_id, thumbnails, thumbnails_blur_hash, created_at
`

type CreateAnimeSerieEpisodeParams struct {
	EpisodeNumber      int32
	SeasonID           int64
	Thumbnails         string
	ThumbnailsBlurHash string
}

func (q *Queries) CreateAnimeSerieEpisode(ctx context.Context, arg CreateAnimeSerieEpisodeParams) (AnimeSerieEpisode, error) {
	row := q.db.QueryRow(ctx, createAnimeSerieEpisode,
		arg.EpisodeNumber,
		arg.SeasonID,
		arg.Thumbnails,
		arg.ThumbnailsBlurHash,
	)
	var i AnimeSerieEpisode
	err := row.Scan(
		&i.ID,
		&i.EpisodeNumber,
		&i.SeasonID,
		&i.Thumbnails,
		&i.ThumbnailsBlurHash,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeSerieEpisode = `-- name: DeleteAnimeSerieEpisode :exec
DELETE FROM anime_serie_episodes
WHERE id = $1
`

func (q *Queries) DeleteAnimeSerieEpisode(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeSerieEpisode, id)
	return err
}

const getAnimeSerieEpisodeByEpisodeID = `-- name: GetAnimeSerieEpisodeByEpisodeID :one
SELECT id, episode_number, season_id, thumbnails, thumbnails_blur_hash, created_at FROM anime_serie_episodes
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAnimeSerieEpisodeByEpisodeID(ctx context.Context, id int64) (AnimeSerieEpisode, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieEpisodeByEpisodeID, id)
	var i AnimeSerieEpisode
	err := row.Scan(
		&i.ID,
		&i.EpisodeNumber,
		&i.SeasonID,
		&i.Thumbnails,
		&i.ThumbnailsBlurHash,
		&i.CreatedAt,
	)
	return i, err
}

const listAnimeSerieEpisodesBySeasonID = `-- name: ListAnimeSerieEpisodesBySeasonID :many
SELECT id, episode_number, season_id, thumbnails, thumbnails_blur_hash, created_at FROM anime_serie_episodes
WHERE season_id = $1
ORDER BY episode_number
LIMIT $2
OFFSET $3
`

type ListAnimeSerieEpisodesBySeasonIDParams struct {
	SeasonID int64
	Limit    int32
	Offset   int32
}

func (q *Queries) ListAnimeSerieEpisodesBySeasonID(ctx context.Context, arg ListAnimeSerieEpisodesBySeasonIDParams) ([]AnimeSerieEpisode, error) {
	rows, err := q.db.Query(ctx, listAnimeSerieEpisodesBySeasonID, arg.SeasonID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeSerieEpisode{}
	for rows.Next() {
		var i AnimeSerieEpisode
		if err := rows.Scan(
			&i.ID,
			&i.EpisodeNumber,
			&i.SeasonID,
			&i.Thumbnails,
			&i.ThumbnailsBlurHash,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAnimeSerieEpisode = `-- name: UpdateAnimeSerieEpisode :one
UPDATE anime_serie_episodes
SET
  episode_number = COALESCE($1, episode_number),
  thumbnails = COALESCE($2, thumbnails),
  episode_number = COALESCE($1, episode_number),
  thumbnails_blur_hash = COALESCE($3, thumbnails_blur_hash)
WHERE
  id = $4
RETURNING id, episode_number, season_id, thumbnails, thumbnails_blur_hash, created_at
`

type UpdateAnimeSerieEpisodeParams struct {
	EpisodeNumber      pgtype.Int4
	Thumbnails         pgtype.Text
	ThumbnailsBlurHash pgtype.Text
	ID                 int64
}

func (q *Queries) UpdateAnimeSerieEpisode(ctx context.Context, arg UpdateAnimeSerieEpisodeParams) (AnimeSerieEpisode, error) {
	row := q.db.QueryRow(ctx, updateAnimeSerieEpisode,
		arg.EpisodeNumber,
		arg.Thumbnails,
		arg.ThumbnailsBlurHash,
		arg.ID,
	)
	var i AnimeSerieEpisode
	err := row.Scan(
		&i.ID,
		&i.EpisodeNumber,
		&i.SeasonID,
		&i.Thumbnails,
		&i.ThumbnailsBlurHash,
		&i.CreatedAt,
	)
	return i, err
}
