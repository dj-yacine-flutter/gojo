// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_tag.sql

package db

import (
	"context"
)

const createAnimeMovieTag = `-- name: CreateAnimeMovieTag :one
INSERT INTO anime_movie_tags (anime_id, tag_id)
VALUES ($1, $2)
RETURNING id, anime_id, tag_id, created_at
`

type CreateAnimeMovieTagParams struct {
	AnimeID int64
	TagID   int64
}

func (q *Queries) CreateAnimeMovieTag(ctx context.Context, arg CreateAnimeMovieTagParams) (AnimeMovieTag, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieTag, arg.AnimeID, arg.TagID)
	var i AnimeMovieTag
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.TagID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeMovieTag = `-- name: DeleteAnimeMovieTag :exec
DELETE FROM anime_movie_tags
WHERE anime_id = $1 AND tag_id = $2
`

type DeleteAnimeMovieTagParams struct {
	AnimeID int64
	TagID   int64
}

func (q *Queries) DeleteAnimeMovieTag(ctx context.Context, arg DeleteAnimeMovieTagParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieTag, arg.AnimeID, arg.TagID)
	return err
}

const listAnimeMovieTags = `-- name: ListAnimeMovieTags :many
SELECT tag_id FROM anime_movie_tags
WHERE anime_id = $1
`

func (q *Queries) ListAnimeMovieTags(ctx context.Context, animeID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeMovieTags, animeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var tag_id int64
		if err := rows.Scan(&tag_id); err != nil {
			return nil, err
		}
		items = append(items, tag_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
