// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_serie_join_meta.sql

package db

import (
	"context"
)

const getAnimeSerieMetaWithLanguageDirectly = `-- name: GetAnimeSerieMetaWithLanguageDirectly :one
SELECT m.id, m.title, m.overview, m.created_at
FROM anime_serie_metas AS asm
JOIN metas AS m ON asm.meta_id = m.id
WHERE asm.anime_id = $1 AND asm.language_id = $2
LIMIT 1
`

type GetAnimeSerieMetaWithLanguageDirectlyParams struct {
	AnimeID    int64
	LanguageID int32
}

func (q *Queries) GetAnimeSerieMetaWithLanguageDirectly(ctx context.Context, arg GetAnimeSerieMetaWithLanguageDirectlyParams) (Meta, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieMetaWithLanguageDirectly, arg.AnimeID, arg.LanguageID)
	var i Meta
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.CreatedAt,
	)
	return i, err
}
