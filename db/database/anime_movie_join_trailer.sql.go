// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_movie_join_trailer.sql

package db

import (
	"context"
)

const getAnimeMovieTrailersDirectly = `-- name: GetAnimeMovieTrailersDirectly :many
SELECT anime_trailers.id, anime_trailers.is_official, anime_trailers.host_name, anime_trailers.host_key, anime_trailers.created_at
FROM anime_trailers
JOIN anime_movie_trailers ON anime_trailers.id = anime_movie_trailers.trailer_id
WHERE anime_movie_trailers.anime_id = $1
`

func (q *Queries) GetAnimeMovieTrailersDirectly(ctx context.Context, animeID int64) ([]AnimeTrailer, error) {
	rows, err := q.db.Query(ctx, getAnimeMovieTrailersDirectly, animeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeTrailer{}
	for rows.Next() {
		var i AnimeTrailer
		if err := rows.Scan(
			&i.ID,
			&i.IsOfficial,
			&i.HostName,
			&i.HostKey,
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
