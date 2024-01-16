-- name: CreateAnimeMovieOtherTitle :one
INSERT INTO anime_movie_other_titles (anime_id, title_text)
VALUES ($1, $2)
ON CONFLICT (anime_id, title_text)
DO UPDATE SET 
    anime_id = excluded.anime_id
RETURNING *;

-- name: QueryAnimeMovieOtherTitles :many
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
  $1::varchar IS NOT NULL AND $1::varchar <> '' AND
  (
    SELECT bool_and(
      to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))) 
      @@ plainto_tsquery(word)
    )
    FROM unnest(regexp_split_to_array($1::varchar, '\\W+')) AS word
  )
  OR title_text ILIKE '%' || $1::varchar || '%'
)
ORDER BY
  ts_rank(title_text_tsv, phraseto_tsquery($1::varchar)) DESC,
  similarity(title_text, $1::varchar) DESC
  LIMIT $2
  OFFSET $3;

-- name: GetAnimeMovieOtherTitles :many
SELECT * FROM anime_movie_other_titles
WHERE anime_id = $1;

-- name: DeleteAnimeMovieOtherTitle :exec
DELETE FROM anime_movie_other_titles
WHERE id = $1;