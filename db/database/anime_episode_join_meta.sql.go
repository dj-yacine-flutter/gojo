// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_episode_join_meta.sql

package db

import (
	"context"
)

const getAnimeEpisodeMetaWithLanguageDirectly = `-- name: GetAnimeEpisodeMetaWithLanguageDirectly :one
SELECT m.id, m.title, m.overview, m.created_at
FROM anime_episode_metas AS aem
JOIN metas AS m ON aem.meta_id = m.id
WHERE aem.episode_id = $1 AND aem.language_id = $2
LIMIT 1
`

type GetAnimeEpisodeMetaWithLanguageDirectlyParams struct {
	EpisodeID  int64
	LanguageID int32
}

func (q *Queries) GetAnimeEpisodeMetaWithLanguageDirectly(ctx context.Context, arg GetAnimeEpisodeMetaWithLanguageDirectlyParams) (Meta, error) {
	row := q.db.QueryRow(ctx, getAnimeEpisodeMetaWithLanguageDirectly, arg.EpisodeID, arg.LanguageID)
	var i Meta
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.CreatedAt,
	)
	return i, err
}
