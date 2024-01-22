// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_join_meta.sql

package db

import (
	"context"
)

const getAnimeMovieMetaWithLanguageDirectly = `-- name: GetAnimeMovieMetaWithLanguageDirectly :one
SELECT m.id, m.title, m.overview, m.created_at
FROM anime_movie_metas AS amm
JOIN metas AS m ON amm.meta_id = m.id
WHERE amm.anime_id = $1 AND amm.language_id = $2
LIMIT 1
`

type GetAnimeMovieMetaWithLanguageDirectlyParams struct {
	AnimeID    int64
	LanguageID int32
}

func (q *Queries) GetAnimeMovieMetaWithLanguageDirectly(ctx context.Context, arg GetAnimeMovieMetaWithLanguageDirectlyParams) (Meta, error) {
	row := q.db.QueryRow(ctx, getAnimeMovieMetaWithLanguageDirectly, arg.AnimeID, arg.LanguageID)
	var i Meta
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.CreatedAt,
	)
	return i, err
}
