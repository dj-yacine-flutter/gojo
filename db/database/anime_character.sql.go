// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_character.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeCharacter = `-- name: CreateAnimeCharacter :one
INSERT INTO anime_characters (full_name, about, role_playing, image_url, image_blur_hash, actors_id, pictures)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (full_name, about)
DO UPDATE SET 
    actors_id = array_remove(array_cat(anime_characters.actors_id, excluded.actors_id), NULL),
    pictures = array_remove(array_cat(anime_characters.pictures, excluded.pictures), NULL)
RETURNING id, full_name, about, role_playing, image_url, image_blur_hash, actors_id, pictures, created_at
`

type CreateAnimeCharacterParams struct {
	FullName      string
	About         string
	RolePlaying   string
	ImageUrl      string
	ImageBlurHash string
	ActorsID      []int64
	Pictures      []string
}

func (q *Queries) CreateAnimeCharacter(ctx context.Context, arg CreateAnimeCharacterParams) (AnimeCharacter, error) {
	row := q.db.QueryRow(ctx, createAnimeCharacter,
		arg.FullName,
		arg.About,
		arg.RolePlaying,
		arg.ImageUrl,
		arg.ImageBlurHash,
		arg.ActorsID,
		arg.Pictures,
	)
	var i AnimeCharacter
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.About,
		&i.RolePlaying,
		&i.ImageUrl,
		&i.ImageBlurHash,
		&i.ActorsID,
		&i.Pictures,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeCharacter = `-- name: DeleteAnimeCharacter :exec
DELETE FROM anime_characters
WHERE id = $1
`

func (q *Queries) DeleteAnimeCharacter(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeCharacter, id)
	return err
}

const getAnimeCharacter = `-- name: GetAnimeCharacter :one
SELECT id, full_name, about, role_playing, image_url, image_blur_hash, actors_id, pictures, created_at FROM anime_characters
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeCharacter(ctx context.Context, id int64) (AnimeCharacter, error) {
	row := q.db.QueryRow(ctx, getAnimeCharacter, id)
	var i AnimeCharacter
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.About,
		&i.RolePlaying,
		&i.ImageUrl,
		&i.ImageBlurHash,
		&i.ActorsID,
		&i.Pictures,
		&i.CreatedAt,
	)
	return i, err
}

const listAnimeCharacters = `-- name: ListAnimeCharacters :many
SELECT id, full_name, about, role_playing, image_url, image_blur_hash, actors_id, pictures, created_at FROM anime_characters
WHERE id = $1
`

func (q *Queries) ListAnimeCharacters(ctx context.Context, id int64) ([]AnimeCharacter, error) {
	rows, err := q.db.Query(ctx, listAnimeCharacters, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeCharacter{}
	for rows.Next() {
		var i AnimeCharacter
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.About,
			&i.RolePlaying,
			&i.ImageUrl,
			&i.ImageBlurHash,
			&i.ActorsID,
			&i.Pictures,
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

const updateAnimeCharacter = `-- name: UpdateAnimeCharacter :one
UPDATE anime_characters
SET
  full_name = COALESCE($1, full_name),
  about = COALESCE($2, about),
  role_playing = COALESCE($3, role_playing),
  image_url = COALESCE($4, image_url),
  image_blur_hash = COALESCE($5, image_blur_hash)
WHERE
  id = $6
RETURNING id, full_name, about, role_playing, image_url, image_blur_hash, actors_id, pictures, created_at
`

type UpdateAnimeCharacterParams struct {
	FullName      pgtype.Text
	About         pgtype.Text
	RolePlaying   pgtype.Text
	ImageUrl      pgtype.Text
	ImageBlurHash pgtype.Text
	ID            int64
}

func (q *Queries) UpdateAnimeCharacter(ctx context.Context, arg UpdateAnimeCharacterParams) (AnimeCharacter, error) {
	row := q.db.QueryRow(ctx, updateAnimeCharacter,
		arg.FullName,
		arg.About,
		arg.RolePlaying,
		arg.ImageUrl,
		arg.ImageBlurHash,
		arg.ID,
	)
	var i AnimeCharacter
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.About,
		&i.RolePlaying,
		&i.ImageUrl,
		&i.ImageBlurHash,
		&i.ActorsID,
		&i.Pictures,
		&i.CreatedAt,
	)
	return i, err
}
