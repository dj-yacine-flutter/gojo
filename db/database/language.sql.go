// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: language.sql

package db

import (
	"context"
)

const createLanguage = `-- name: CreateLanguage :one
INSERT INTO languages (language_name, language_code)
VALUES ($1, $2)
RETURNING  id, language_code, language_name, created_at
`

type CreateLanguageParams struct {
	LanguageName string `json:"language_name"`
	LanguageCode string `json:"language_code"`
}

func (q *Queries) CreateLanguage(ctx context.Context, arg CreateLanguageParams) (Language, error) {
	row := q.db.QueryRow(ctx, createLanguage, arg.LanguageName, arg.LanguageCode)
	var i Language
	err := row.Scan(
		&i.ID,
		&i.LanguageCode,
		&i.LanguageName,
		&i.CreatedAt,
	)
	return i, err
}

const deleteLanguage = `-- name: DeleteLanguage :exec
DELETE FROM languages
WHERE id = $1
`

func (q *Queries) DeleteLanguage(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteLanguage, id)
	return err
}

const getLanguage = `-- name: GetLanguage :one
SELECT id, language_code, language_name, created_at FROM languages
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLanguage(ctx context.Context, id int32) (Language, error) {
	row := q.db.QueryRow(ctx, getLanguage, id)
	var i Language
	err := row.Scan(
		&i.ID,
		&i.LanguageCode,
		&i.LanguageName,
		&i.CreatedAt,
	)
	return i, err
}

const listLanguages = `-- name: ListLanguages :many
SELECT id, language_code, language_name, created_at FROM languages
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListLanguagesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListLanguages(ctx context.Context, arg ListLanguagesParams) ([]Language, error) {
	rows, err := q.db.Query(ctx, listLanguages, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Language{}
	for rows.Next() {
		var i Language
		if err := rows.Scan(
			&i.ID,
			&i.LanguageCode,
			&i.LanguageName,
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

const updateLanguage = `-- name: UpdateLanguage :one
UPDATE languages
SET language_code = $2,
    language_name = $3
WHERE id = $1
RETURNING id, language_code, language_name, created_at
`

type UpdateLanguageParams struct {
	ID           int32  `json:"id"`
	LanguageCode string `json:"language_code"`
	LanguageName string `json:"language_name"`
}

func (q *Queries) UpdateLanguage(ctx context.Context, arg UpdateLanguageParams) (Language, error) {
	row := q.db.QueryRow(ctx, updateLanguage, arg.ID, arg.LanguageCode, arg.LanguageName)
	var i Language
	err := row.Scan(
		&i.ID,
		&i.LanguageCode,
		&i.LanguageName,
		&i.CreatedAt,
	)
	return i, err
}
