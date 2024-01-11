// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: character_actor.sql

package db

import (
	"context"
)

const createAnimeCharacterActor = `-- name: CreateAnimeCharacterActor :exec
INSERT INTO anime_character_actors (character_id, actor_id)
VALUES ($1, $2)
ON CONFLICT (character_id, actor_id)
DO NOTHING
`

type CreateAnimeCharacterActorParams struct {
	CharacterID int64
	ActorID     int64
}

func (q *Queries) CreateAnimeCharacterActor(ctx context.Context, arg CreateAnimeCharacterActorParams) error {
	_, err := q.db.Exec(ctx, createAnimeCharacterActor, arg.CharacterID, arg.ActorID)
	return err
}

const deleteAnimeCharacterActor = `-- name: DeleteAnimeCharacterActor :exec
DELETE FROM anime_character_actors
WHERE id = $1
`

func (q *Queries) DeleteAnimeCharacterActor(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeCharacterActor, id)
	return err
}

const listAnimeCharacterActors = `-- name: ListAnimeCharacterActors :many
SELECT id, character_id, actor_id FROM anime_character_actors
WHERE character_id = $1
`

func (q *Queries) ListAnimeCharacterActors(ctx context.Context, characterID int64) ([]AnimeCharacterActor, error) {
	rows, err := q.db.Query(ctx, listAnimeCharacterActors, characterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeCharacterActor{}
	for rows.Next() {
		var i AnimeCharacterActor
		if err := rows.Scan(&i.ID, &i.CharacterID, &i.ActorID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
