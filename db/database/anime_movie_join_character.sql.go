// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_join_character.sql

package db

import (
	"context"
)

const getAnimeMovieCharactersDirectly = `-- name: GetAnimeMovieCharactersDirectly :many
SELECT anime_characters.id, anime_characters.full_name, anime_characters.about, anime_characters.role_playing, anime_characters.image_url, anime_characters.image_blur_hash, anime_characters.pictures, anime_characters.created_at
FROM anime_characters
JOIN anime_movie_characters ON anime_characters.id = anime_movie_characters.studio_id
WHERE anime_movie_characters.anime_id = $1
`

func (q *Queries) GetAnimeMovieCharactersDirectly(ctx context.Context, animeID int64) ([]AnimeCharacter, error) {
	rows, err := q.db.Query(ctx, getAnimeMovieCharactersDirectly, animeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeCharacter{}
	for rows.Next() {
		var i AnimeCharacter
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.About,
			&i.RolePlaying,
			&i.ImageUrl,
			&i.ImageBlurHash,
			&i.Pictures,
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
