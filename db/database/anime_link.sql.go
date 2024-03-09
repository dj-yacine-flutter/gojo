// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_link.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeLink = `-- name: CreateAnimeLink :one
INSERT INTO anime_links (official_website, wikipedia_url, crunchyroll_url, social_media)
VALUES ($1, $2, $3, $4)
RETURNING  id, official_website, wikipedia_url, crunchyroll_url, social_media, created_at
`

type CreateAnimeLinkParams struct {
	OfficialWebsite string
	WikipediaUrl    string
	CrunchyrollUrl  string
	SocialMedia     []string
}

func (q *Queries) CreateAnimeLink(ctx context.Context, arg CreateAnimeLinkParams) (AnimeLink, error) {
	row := q.db.QueryRow(ctx, createAnimeLink,
		arg.OfficialWebsite,
		arg.WikipediaUrl,
		arg.CrunchyrollUrl,
		arg.SocialMedia,
	)
	var i AnimeLink
	err := row.Scan(
		&i.ID,
		&i.OfficialWebsite,
		&i.WikipediaUrl,
		&i.CrunchyrollUrl,
		&i.SocialMedia,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeLink = `-- name: DeleteAnimeLink :exec
DELETE FROM anime_links
WHERE id = $1
`

func (q *Queries) DeleteAnimeLink(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeLink, id)
	return err
}

const getAnimeLink = `-- name: GetAnimeLink :one
SELECT id, official_website, wikipedia_url, crunchyroll_url, social_media, created_at FROM anime_links
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeLink(ctx context.Context, id int64) (AnimeLink, error) {
	row := q.db.QueryRow(ctx, getAnimeLink, id)
	var i AnimeLink
	err := row.Scan(
		&i.ID,
		&i.OfficialWebsite,
		&i.WikipediaUrl,
		&i.CrunchyrollUrl,
		&i.SocialMedia,
		&i.CreatedAt,
	)
	return i, err
}

const updateAnimeLink = `-- name: UpdateAnimeLink :one
UPDATE anime_links
SET
  official_website = COALESCE($1, official_website),
  wikipedia_url = COALESCE($2, wikipedia_url),
  crunchyroll_url = COALESCE($3, crunchyroll_url),
  social_media = COALESCE($4, social_media)
WHERE
  id = $5
RETURNING id, official_website, wikipedia_url, crunchyroll_url, social_media, created_at
`

type UpdateAnimeLinkParams struct {
	OfficialWebsite pgtype.Text
	WikipediaUrl    pgtype.Text
	CrunchyrollUrl  pgtype.Text
	SocialMedia     []string
	ID              int64
}

func (q *Queries) UpdateAnimeLink(ctx context.Context, arg UpdateAnimeLinkParams) (AnimeLink, error) {
	row := q.db.QueryRow(ctx, updateAnimeLink,
		arg.OfficialWebsite,
		arg.WikipediaUrl,
		arg.CrunchyrollUrl,
		arg.SocialMedia,
		arg.ID,
	)
	var i AnimeLink
	err := row.Scan(
		&i.ID,
		&i.OfficialWebsite,
		&i.WikipediaUrl,
		&i.CrunchyrollUrl,
		&i.SocialMedia,
		&i.CreatedAt,
	)
	return i, err
}
