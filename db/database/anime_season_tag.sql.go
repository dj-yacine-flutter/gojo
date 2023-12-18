// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_season_tag.sql

package db

import (
	"context"
)

const createAnimeSeasonTag = `-- name: CreateAnimeSeasonTag :one
INSERT INTO anime_season_tags (season_id, tag_id)
VALUES ($1, $2)
RETURNING id, season_id, tag_id, created_at
`

type CreateAnimeSeasonTagParams struct {
	SeasonID int64
	TagID    int64
}

func (q *Queries) CreateAnimeSeasonTag(ctx context.Context, arg CreateAnimeSeasonTagParams) (AnimeSeasonTag, error) {
	row := q.db.QueryRow(ctx, createAnimeSeasonTag, arg.SeasonID, arg.TagID)
	var i AnimeSeasonTag
	err := row.Scan(
		&i.ID,
		&i.SeasonID,
		&i.TagID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeSeasonTag = `-- name: DeleteAnimeSeasonTag :exec
DELETE FROM anime_season_tags
WHERE season_id = $1 AND tag_id = $2
`

type DeleteAnimeSeasonTagParams struct {
	SeasonID int64
	TagID    int64
}

func (q *Queries) DeleteAnimeSeasonTag(ctx context.Context, arg DeleteAnimeSeasonTagParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeSeasonTag, arg.SeasonID, arg.TagID)
	return err
}

const listAnimeSeasonTags = `-- name: ListAnimeSeasonTags :many
SELECT id, season_id, tag_id, created_at FROM anime_season_tags
WHERE season_id = $1
`

func (q *Queries) ListAnimeSeasonTags(ctx context.Context, seasonID int64) ([]AnimeSeasonTag, error) {
	rows, err := q.db.Query(ctx, listAnimeSeasonTags, seasonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeSeasonTag{}
	for rows.Next() {
		var i AnimeSeasonTag
		if err := rows.Scan(
			&i.ID,
			&i.SeasonID,
			&i.TagID,
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
