// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_serie_join_link.sql

package db

import (
	"context"
)

const getAnimeSerieLinksDirectly = `-- name: GetAnimeSerieLinksDirectly :one
SELECT anime_links.id, anime_links.official_website, anime_links.wikipedia_url, anime_links.crunchyroll_url, anime_links.social_media, anime_links.created_at
FROM anime_links
JOIN anime_serie_links ON anime_links.id = anime_serie_links.link_id
WHERE anime_serie_links.anime_id = $1
LIMIT 1
`

func (q *Queries) GetAnimeSerieLinksDirectly(ctx context.Context, animeID int64) (AnimeLink, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieLinksDirectly, animeID)
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
