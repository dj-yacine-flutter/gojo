// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_meta.sql

package db

import (
	"context"
)

const createAnimeMovieMeta = `-- name: CreateAnimeMovieMeta :one
INSERT INTO anime_movie_metas (anime_id, language_id, meta_id)
VALUES ($1, $2, $3)
RETURNING id, anime_id, language_id, meta_id
`

type CreateAnimeMovieMetaParams struct {
	AnimeID    int64 `json:"anime_id"`
	LanguageID int32 `json:"language_id"`
	MetaID     int64 `json:"meta_id"`
}

func (q *Queries) CreateAnimeMovieMeta(ctx context.Context, arg CreateAnimeMovieMetaParams) (AnimeMovieMeta, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieMeta, arg.AnimeID, arg.LanguageID, arg.MetaID)
	var i AnimeMovieMeta
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.LanguageID,
		&i.MetaID,
	)
	return i, err
}

const deleteAnimeMovieMeta = `-- name: DeleteAnimeMovieMeta :exec
DELETE FROM anime_movie_metas
WHERE anime_id = $1 AND language_id = $2
`

type DeleteAnimeMovieMetaParams struct {
	AnimeID    int64 `json:"anime_id"`
	LanguageID int32 `json:"language_id"`
}

func (q *Queries) DeleteAnimeMovieMeta(ctx context.Context, arg DeleteAnimeMovieMetaParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieMeta, arg.AnimeID, arg.LanguageID)
	return err
}

const getAnimeMovieMeta = `-- name: GetAnimeMovieMeta :one
SELECT meta_id
FROM anime_movie_metas
WHERE anime_id = $1 AND language_id = $2
`

type GetAnimeMovieMetaParams struct {
	AnimeID    int64 `json:"anime_id"`
	LanguageID int32 `json:"language_id"`
}

func (q *Queries) GetAnimeMovieMeta(ctx context.Context, arg GetAnimeMovieMetaParams) (int64, error) {
	row := q.db.QueryRow(ctx, getAnimeMovieMeta, arg.AnimeID, arg.LanguageID)
	var meta_id int64
	err := row.Scan(&meta_id)
	return meta_id, err
}

const listAnimeMovieMetas = `-- name: ListAnimeMovieMetas :many
SELECT meta_id
FROM anime_movie_metas
WHERE anime_id = $1
LIMIT $2
OFFSET $3
`

type ListAnimeMovieMetasParams struct {
	AnimeID int64 `json:"anime_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListAnimeMovieMetas(ctx context.Context, arg ListAnimeMovieMetasParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeMovieMetas, arg.AnimeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var meta_id int64
		if err := rows.Scan(&meta_id); err != nil {
			return nil, err
		}
		items = append(items, meta_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAnimeMovieMeta = `-- name: UpdateAnimeMovieMeta :one
UPDATE anime_movie_metas
SET meta_id = $3
WHERE anime_id = $1 AND language_id = $2
RETURNING id, anime_id, language_id, meta_id
`

type UpdateAnimeMovieMetaParams struct {
	AnimeID    int64 `json:"anime_id"`
	LanguageID int32 `json:"language_id"`
	MetaID     int64 `json:"meta_id"`
}

func (q *Queries) UpdateAnimeMovieMeta(ctx context.Context, arg UpdateAnimeMovieMetaParams) (AnimeMovieMeta, error) {
	row := q.db.QueryRow(ctx, updateAnimeMovieMeta, arg.AnimeID, arg.LanguageID, arg.MetaID)
	var i AnimeMovieMeta
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.LanguageID,
		&i.MetaID,
	)
	return i, err
}
