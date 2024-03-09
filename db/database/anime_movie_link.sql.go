// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_movie_link.sql

package db

import (
	"context"
)

const createAnimeMovieLink = `-- name: CreateAnimeMovieLink :one
INSERT INTO anime_movie_links (anime_id, link_id)
VALUES ($1, $2)
RETURNING id, anime_id, link_id
`

type CreateAnimeMovieLinkParams struct {
	AnimeID int64
	LinkID  int64
}

func (q *Queries) CreateAnimeMovieLink(ctx context.Context, arg CreateAnimeMovieLinkParams) (AnimeMovieLink, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieLink, arg.AnimeID, arg.LinkID)
	var i AnimeMovieLink
	err := row.Scan(&i.ID, &i.AnimeID, &i.LinkID)
	return i, err
}

const deleteAnimeMovieLink = `-- name: DeleteAnimeMovieLink :exec
DELETE FROM anime_movie_links
WHERE anime_id = $1 AND link_id = $2
`

type DeleteAnimeMovieLinkParams struct {
	AnimeID int64
	LinkID  int64
}

func (q *Queries) DeleteAnimeMovieLink(ctx context.Context, arg DeleteAnimeMovieLinkParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieLink, arg.AnimeID, arg.LinkID)
	return err
}

const getAnimeMovieLink = `-- name: GetAnimeMovieLink :one
SELECT id, anime_id, link_id FROM anime_movie_links
WHERE anime_id = $1
LIMIT 1
`

func (q *Queries) GetAnimeMovieLink(ctx context.Context, animeID int64) (AnimeMovieLink, error) {
	row := q.db.QueryRow(ctx, getAnimeMovieLink, animeID)
	var i AnimeMovieLink
	err := row.Scan(&i.ID, &i.AnimeID, &i.LinkID)
	return i, err
}
