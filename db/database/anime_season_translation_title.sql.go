// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_season_translation_title.sql

package db

import (
	"context"
)

const createAnimeSeasonTranslationTitle = `-- name: CreateAnimeSeasonTranslationTitle :one
INSERT INTO anime_season_translation_titles (season_id, title_text)
VALUES ($1, $2)
RETURNING id, season_id, title_text, created_at
`

type CreateAnimeSeasonTranslationTitleParams struct {
	SeasonID  int64
	TitleText string
}

func (q *Queries) CreateAnimeSeasonTranslationTitle(ctx context.Context, arg CreateAnimeSeasonTranslationTitleParams) (AnimeSeasonTranslationTitle, error) {
	row := q.db.QueryRow(ctx, createAnimeSeasonTranslationTitle, arg.SeasonID, arg.TitleText)
	var i AnimeSeasonTranslationTitle
	err := row.Scan(
		&i.ID,
		&i.SeasonID,
		&i.TitleText,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeSeasonTranslationTitle = `-- name: DeleteAnimeSeasonTranslationTitle :exec
DELETE FROM anime_season_translation_titles
WHERE id = $1
`

func (q *Queries) DeleteAnimeSeasonTranslationTitle(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeSeasonTranslationTitle, id)
	return err
}

const getAnimeSeasonTranslationTitles = `-- name: GetAnimeSeasonTranslationTitles :many
SELECT id, season_id, title_text, created_at FROM anime_season_translation_titles
WHERE season_id = $1
`

func (q *Queries) GetAnimeSeasonTranslationTitles(ctx context.Context, seasonID int64) ([]AnimeSeasonTranslationTitle, error) {
	rows, err := q.db.Query(ctx, getAnimeSeasonTranslationTitles, seasonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeSeasonTranslationTitle{}
	for rows.Next() {
		var i AnimeSeasonTranslationTitle
		if err := rows.Scan(
			&i.ID,
			&i.SeasonID,
			&i.TitleText,
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

const queryAnimeSeasonTranslationTitles = `-- name: QueryAnimeSeasonTranslationTitles :many
WITH search_documents AS (
  SELECT
    season_id,
    title_text,
    to_tsvector('pg_catalog.simple', title_text) AS title_text_tsv
  FROM anime_season_translation_titles
)
SELECT season_id
FROM search_documents
WHERE (
  $1::text IS NOT NULL AND $1::text <> '' AND
  (
    SELECT bool_and(
      to_tsvector('pg_catalog.simple', lower(translate(title_text, '[:punct:]', ''))) 
      @@ plainto_tsquery(word)
    )
    FROM unnest(regexp_split_to_array($1::text, '\\W+')) AS word
  )
  OR title_text ILIKE '%' || $1::text || '%'
)
ORDER BY
  ts_rank(title_text_tsv, phraseto_tsquery($1::text)) DESC,
  similarity(title_text, $1::text) DESC
`

func (q *Queries) QueryAnimeSeasonTranslationTitles(ctx context.Context, dollar_1 string) ([]int64, error) {
	rows, err := q.db.Query(ctx, queryAnimeSeasonTranslationTitles, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var season_id int64
		if err := rows.Scan(&season_id); err != nil {
			return nil, err
		}
		items = append(items, season_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
