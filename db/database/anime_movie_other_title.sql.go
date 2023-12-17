// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_other_title.sql

package db

import (
	"context"
)

const createAnimeMovieOtherTitle = `-- name: CreateAnimeMovieOtherTitle :one
INSERT INTO anime_movie_other_titles (anime_id, title_text)
VALUES ($1, $2)
RETURNING id, anime_id, title_text, created_at
`

type CreateAnimeMovieOtherTitleParams struct {
	AnimeID   int64
	TitleText string
}

func (q *Queries) CreateAnimeMovieOtherTitle(ctx context.Context, arg CreateAnimeMovieOtherTitleParams) (AnimeMovieOtherTitle, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieOtherTitle, arg.AnimeID, arg.TitleText)
	var i AnimeMovieOtherTitle
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.TitleText,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeMovieOtherTitle = `-- name: DeleteAnimeMovieOtherTitle :exec
DELETE FROM anime_movie_other_titles
WHERE id = $1
`

func (q *Queries) DeleteAnimeMovieOtherTitle(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieOtherTitle, id)
	return err
}

const getAnimeMovieOtherTitles = `-- name: GetAnimeMovieOtherTitles :many
SELECT id, anime_id, title_text, created_at FROM anime_movie_other_titles
WHERE anime_id = $1
`

func (q *Queries) GetAnimeMovieOtherTitles(ctx context.Context, animeID int64) ([]AnimeMovieOtherTitle, error) {
	rows, err := q.db.Query(ctx, getAnimeMovieOtherTitles, animeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeMovieOtherTitle{}
	for rows.Next() {
		var i AnimeMovieOtherTitle
		if err := rows.Scan(
			&i.ID,
			&i.AnimeID,
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

const queryAnimeMovieOtherTitles = `-- name: QueryAnimeMovieOtherTitles :many
WITH search_documents AS (
  SELECT
    anime_id,
    title_text,
    to_tsvector('pg_catalog.english', title_text) AS title_text_tsv
  FROM anime_movie_other_titles
)
SELECT anime_id
FROM search_documents
WHERE (
  $1::text IS NOT NULL AND $1::text <> '' AND
  (
    SELECT bool_and(
      to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))) 
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

func (q *Queries) QueryAnimeMovieOtherTitles(ctx context.Context, dollar_1 string) ([]int64, error) {
	rows, err := q.db.Query(ctx, queryAnimeMovieOtherTitles, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var anime_id int64
		if err := rows.Scan(&anime_id); err != nil {
			return nil, err
		}
		items = append(items, anime_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}