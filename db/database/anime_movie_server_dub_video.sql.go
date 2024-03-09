// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_movie_server_dub_video.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeMovieServerDubVideo = `-- name: CreateAnimeMovieServerDubVideo :one
INSERT INTO anime_movie_server_dub_videos (server_id, video_id)
VALUES ($1, $2)
RETURNING id, server_id, video_id, created_at
`

type CreateAnimeMovieServerDubVideoParams struct {
	ServerID int64
	VideoID  int64
}

func (q *Queries) CreateAnimeMovieServerDubVideo(ctx context.Context, arg CreateAnimeMovieServerDubVideoParams) (AnimeMovieServerDubVideo, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieServerDubVideo, arg.ServerID, arg.VideoID)
	var i AnimeMovieServerDubVideo
	err := row.Scan(
		&i.ID,
		&i.ServerID,
		&i.VideoID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeMovieServerDubVideo = `-- name: DeleteAnimeMovieServerDubVideo :exec
DELETE FROM anime_movie_server_dub_videos
WHERE id = $1
`

func (q *Queries) DeleteAnimeMovieServerDubVideo(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieServerDubVideo, id)
	return err
}

const getAnimeMovieServerDubVideo = `-- name: GetAnimeMovieServerDubVideo :one
SELECT id, server_id, video_id, created_at FROM anime_movie_server_dub_videos
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAnimeMovieServerDubVideo(ctx context.Context, id int64) (AnimeMovieServerDubVideo, error) {
	row := q.db.QueryRow(ctx, getAnimeMovieServerDubVideo, id)
	var i AnimeMovieServerDubVideo
	err := row.Scan(
		&i.ID,
		&i.ServerID,
		&i.VideoID,
		&i.CreatedAt,
	)
	return i, err
}

const listAnimeMovieServerDubVideos = `-- name: ListAnimeMovieServerDubVideos :many
SELECT id, server_id, video_id, created_at FROM anime_movie_server_dub_videos
WHERE server_id = $1
ORDER BY id
`

func (q *Queries) ListAnimeMovieServerDubVideos(ctx context.Context, serverID int64) ([]AnimeMovieServerDubVideo, error) {
	rows, err := q.db.Query(ctx, listAnimeMovieServerDubVideos, serverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeMovieServerDubVideo{}
	for rows.Next() {
		var i AnimeMovieServerDubVideo
		if err := rows.Scan(
			&i.ID,
			&i.ServerID,
			&i.VideoID,
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

const updateAnimeMovieServerDubVideo = `-- name: UpdateAnimeMovieServerDubVideo :one
UPDATE anime_movie_server_dub_videos
SET 
  server_id = COALESCE($1, server_id),
  video_id = COALESCE($2, video_id)
WHERE
  id = $3
RETURNING id, server_id, video_id, created_at
`

type UpdateAnimeMovieServerDubVideoParams struct {
	ServerID pgtype.Int8
	VideoID  pgtype.Int8
	ID       int64
}

func (q *Queries) UpdateAnimeMovieServerDubVideo(ctx context.Context, arg UpdateAnimeMovieServerDubVideoParams) (AnimeMovieServerDubVideo, error) {
	row := q.db.QueryRow(ctx, updateAnimeMovieServerDubVideo, arg.ServerID, arg.VideoID, arg.ID)
	var i AnimeMovieServerDubVideo
	err := row.Scan(
		&i.ID,
		&i.ServerID,
		&i.VideoID,
		&i.CreatedAt,
	)
	return i, err
}
