// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_video.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeMovieVideo = `-- name: CreateAnimeMovieVideo :one
INSERT INTO anime_movie_videos (language_id, autority, referer, link, quality)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, language_id, autority, referer, link, quality, created_at
`

type CreateAnimeMovieVideoParams struct {
	LanguageID int32
	Autority   string
	Referer    string
	Link       string
	Quality    string
}

func (q *Queries) CreateAnimeMovieVideo(ctx context.Context, arg CreateAnimeMovieVideoParams) (AnimeMovieVideo, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieVideo,
		arg.LanguageID,
		arg.Autority,
		arg.Referer,
		arg.Link,
		arg.Quality,
	)
	var i AnimeMovieVideo
	err := row.Scan(
		&i.ID,
		&i.LanguageID,
		&i.Autority,
		&i.Referer,
		&i.Link,
		&i.Quality,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeMovieVideo = `-- name: DeleteAnimeMovieVideo :exec
DELETE FROM anime_movie_videos
WHERE id = $1
`

func (q *Queries) DeleteAnimeMovieVideo(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieVideo, id)
	return err
}

const getAnimeMovieVideo = `-- name: GetAnimeMovieVideo :one
SELECT id, language_id, autority, referer, link, quality, created_at FROM anime_movie_videos
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAnimeMovieVideo(ctx context.Context, id int64) (AnimeMovieVideo, error) {
	row := q.db.QueryRow(ctx, getAnimeMovieVideo, id)
	var i AnimeMovieVideo
	err := row.Scan(
		&i.ID,
		&i.LanguageID,
		&i.Autority,
		&i.Referer,
		&i.Link,
		&i.Quality,
		&i.CreatedAt,
	)
	return i, err
}

const listAnimeMovieVideos = `-- name: ListAnimeMovieVideos :many
SELECT id, language_id, autority, referer, link, quality, created_at FROM anime_movie_videos
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAnimeMovieVideosParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListAnimeMovieVideos(ctx context.Context, arg ListAnimeMovieVideosParams) ([]AnimeMovieVideo, error) {
	rows, err := q.db.Query(ctx, listAnimeMovieVideos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeMovieVideo{}
	for rows.Next() {
		var i AnimeMovieVideo
		if err := rows.Scan(
			&i.ID,
			&i.LanguageID,
			&i.Autority,
			&i.Referer,
			&i.Link,
			&i.Quality,
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

const updateAnimeMovieVideo = `-- name: UpdateAnimeMovieVideo :one
UPDATE anime_movie_videos
SET
    language_id = COALESCE($2, language_id),
    autority = COALESCE($3, autority),
    referer = COALESCE($4, referer),
    link = COALESCE($5, link),
    quality = COALESCE($6, quality)
WHERE id = $1
RETURNING id, language_id, autority, referer, link, quality, created_at
`

type UpdateAnimeMovieVideoParams struct {
	ID         int64
	LanguageID pgtype.Int4
	Autority   pgtype.Text
	Referer    pgtype.Text
	Link       pgtype.Text
	Quality    pgtype.Text
}

func (q *Queries) UpdateAnimeMovieVideo(ctx context.Context, arg UpdateAnimeMovieVideoParams) (AnimeMovieVideo, error) {
	row := q.db.QueryRow(ctx, updateAnimeMovieVideo,
		arg.ID,
		arg.LanguageID,
		arg.Autority,
		arg.Referer,
		arg.Link,
		arg.Quality,
	)
	var i AnimeMovieVideo
	err := row.Scan(
		&i.ID,
		&i.LanguageID,
		&i.Autority,
		&i.Referer,
		&i.Link,
		&i.Quality,
		&i.CreatedAt,
	)
	return i, err
}
