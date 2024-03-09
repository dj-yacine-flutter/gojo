// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_serie_meta.sql

package db

import (
	"context"
)

const createAnimeSerieMeta = `-- name: CreateAnimeSerieMeta :one
INSERT INTO anime_serie_metas (anime_id, language_id, meta_id)
VALUES ($1, $2, $3)
RETURNING id, anime_id, language_id, meta_id
`

type CreateAnimeSerieMetaParams struct {
	AnimeID    int64
	LanguageID int32
	MetaID     int64
}

func (q *Queries) CreateAnimeSerieMeta(ctx context.Context, arg CreateAnimeSerieMetaParams) (AnimeSerieMeta, error) {
	row := q.db.QueryRow(ctx, createAnimeSerieMeta, arg.AnimeID, arg.LanguageID, arg.MetaID)
	var i AnimeSerieMeta
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.LanguageID,
		&i.MetaID,
	)
	return i, err
}

const deleteAnimeSerieMeta = `-- name: DeleteAnimeSerieMeta :exec
DELETE FROM anime_serie_metas
WHERE anime_id = $1 AND language_id = $2
`

type DeleteAnimeSerieMetaParams struct {
	AnimeID    int64
	LanguageID int32
}

func (q *Queries) DeleteAnimeSerieMeta(ctx context.Context, arg DeleteAnimeSerieMetaParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeSerieMeta, arg.AnimeID, arg.LanguageID)
	return err
}

const getAnimeSerieMeta = `-- name: GetAnimeSerieMeta :one
SELECT meta_id
FROM anime_serie_metas
WHERE anime_id = $1 AND language_id = $2
`

type GetAnimeSerieMetaParams struct {
	AnimeID    int64
	LanguageID int32
}

func (q *Queries) GetAnimeSerieMeta(ctx context.Context, arg GetAnimeSerieMetaParams) (int64, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieMeta, arg.AnimeID, arg.LanguageID)
	var meta_id int64
	err := row.Scan(&meta_id)
	return meta_id, err
}

const getAnimeSerieMetaByID = `-- name: GetAnimeSerieMetaByID :one
SELECT id, anime_id, language_id, meta_id 
FROM anime_serie_metas
WHERE id = $1
ORDER BY id
`

func (q *Queries) GetAnimeSerieMetaByID(ctx context.Context, id int64) (AnimeSerieMeta, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieMetaByID, id)
	var i AnimeSerieMeta
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.LanguageID,
		&i.MetaID,
	)
	return i, err
}

const listAnimeSerieMetas = `-- name: ListAnimeSerieMetas :many
SELECT meta_id
FROM anime_serie_metas
WHERE anime_id = $1
ORDER BY id
`

func (q *Queries) ListAnimeSerieMetas(ctx context.Context, animeID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeSerieMetas, animeID)
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

const updateAnimeSerieMeta = `-- name: UpdateAnimeSerieMeta :one
UPDATE anime_serie_metas
SET meta_id = $3
WHERE anime_id = $1 AND language_id = $2
RETURNING id, anime_id, language_id, meta_id
`

type UpdateAnimeSerieMetaParams struct {
	AnimeID    int64
	LanguageID int32
	MetaID     int64
}

func (q *Queries) UpdateAnimeSerieMeta(ctx context.Context, arg UpdateAnimeSerieMetaParams) (AnimeSerieMeta, error) {
	row := q.db.QueryRow(ctx, updateAnimeSerieMeta, arg.AnimeID, arg.LanguageID, arg.MetaID)
	var i AnimeSerieMeta
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.LanguageID,
		&i.MetaID,
	)
	return i, err
}
