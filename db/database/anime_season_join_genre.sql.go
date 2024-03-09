// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_season_join_genre.sql

package db

import (
	"context"
)

const getAnimeSeasonGenresDirectly = `-- name: GetAnimeSeasonGenresDirectly :many
SELECT genres.id, genres.genre_name, genres.created_at
FROM genres
JOIN anime_season_genres ON genres.id = anime_season_genres.genre_id
WHERE anime_season_genres.season_id = $1
`

func (q *Queries) GetAnimeSeasonGenresDirectly(ctx context.Context, seasonID int64) ([]Genre, error) {
	rows, err := q.db.Query(ctx, getAnimeSeasonGenresDirectly, seasonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Genre{}
	for rows.Next() {
		var i Genre
		if err := rows.Scan(&i.ID, &i.GenreName, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
