// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_serie_torrent.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeSerieTorrent = `-- name: CreateAnimeSerieTorrent :one
INSERT INTO anime_serie_torrents (language_id, file_name, torrent_hash, torrent_file, seeds, peers, leechers, size_bytes, quality)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, file_name, language_id, torrent_hash, torrent_file, seeds, peers, leechers, size_bytes, quality, created_at
`

type CreateAnimeSerieTorrentParams struct {
	LanguageID  int32  `json:"language_id"`
	FileName    string `json:"file_name"`
	TorrentHash string `json:"torrent_hash"`
	TorrentFile string `json:"torrent_file"`
	Seeds       int32  `json:"seeds"`
	Peers       int32  `json:"peers"`
	Leechers    int32  `json:"leechers"`
	SizeBytes   int64  `json:"size_bytes"`
	Quality     string `json:"quality"`
}

func (q *Queries) CreateAnimeSerieTorrent(ctx context.Context, arg CreateAnimeSerieTorrentParams) (AnimeSerieTorrent, error) {
	row := q.db.QueryRow(ctx, createAnimeSerieTorrent,
		arg.LanguageID,
		arg.FileName,
		arg.TorrentHash,
		arg.TorrentFile,
		arg.Seeds,
		arg.Peers,
		arg.Leechers,
		arg.SizeBytes,
		arg.Quality,
	)
	var i AnimeSerieTorrent
	err := row.Scan(
		&i.ID,
		&i.FileName,
		&i.LanguageID,
		&i.TorrentHash,
		&i.TorrentFile,
		&i.Seeds,
		&i.Peers,
		&i.Leechers,
		&i.SizeBytes,
		&i.Quality,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeSerieTorrent = `-- name: DeleteAnimeSerieTorrent :exec
DELETE FROM anime_serie_torrents
WHERE id = $1
`

func (q *Queries) DeleteAnimeSerieTorrent(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeSerieTorrent, id)
	return err
}

const getAnimeSerieTorrent = `-- name: GetAnimeSerieTorrent :one
SELECT id, file_name, language_id, torrent_hash, torrent_file, seeds, peers, leechers, size_bytes, quality, created_at FROM anime_serie_torrents
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAnimeSerieTorrent(ctx context.Context, id int64) (AnimeSerieTorrent, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieTorrent, id)
	var i AnimeSerieTorrent
	err := row.Scan(
		&i.ID,
		&i.FileName,
		&i.LanguageID,
		&i.TorrentHash,
		&i.TorrentFile,
		&i.Seeds,
		&i.Peers,
		&i.Leechers,
		&i.SizeBytes,
		&i.Quality,
		&i.CreatedAt,
	)
	return i, err
}

const listAnimeSerieTorrents = `-- name: ListAnimeSerieTorrents :many
SELECT id, file_name, language_id, torrent_hash, torrent_file, seeds, peers, leechers, size_bytes, quality, created_at FROM anime_serie_torrents
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAnimeSerieTorrentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAnimeSerieTorrents(ctx context.Context, arg ListAnimeSerieTorrentsParams) ([]AnimeSerieTorrent, error) {
	rows, err := q.db.Query(ctx, listAnimeSerieTorrents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeSerieTorrent{}
	for rows.Next() {
		var i AnimeSerieTorrent
		if err := rows.Scan(
			&i.ID,
			&i.FileName,
			&i.LanguageID,
			&i.TorrentHash,
			&i.TorrentFile,
			&i.Seeds,
			&i.Peers,
			&i.Leechers,
			&i.SizeBytes,
			&i.Quality,
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

const updateAnimeSerieTorrent = `-- name: UpdateAnimeSerieTorrent :one
UPDATE anime_serie_torrents
SET
    language_id = COALESCE($2, language_id),
    file_name = COALESCE($3, file_name),
    torrent_hash = COALESCE($4, torrent_hash),
    torrent_file = COALESCE($5, torrent_file),
    seeds = COALESCE($6, seeds),
    peers = COALESCE($7, peers),
    leechers = COALESCE($8, leechers),
    size_bytes = COALESCE($9, size_bytes),
    quality = COALESCE($10, quality)
WHERE id = $1
RETURNING id, file_name, language_id, torrent_hash, torrent_file, seeds, peers, leechers, size_bytes, quality, created_at
`

type UpdateAnimeSerieTorrentParams struct {
	ID          int64       `json:"id"`
	LanguageID  pgtype.Int4 `json:"language_id"`
	FileName    pgtype.Text `json:"file_name"`
	TorrentHash pgtype.Text `json:"torrent_hash"`
	TorrentFile pgtype.Text `json:"torrent_file"`
	Seeds       pgtype.Int4 `json:"seeds"`
	Peers       pgtype.Int4 `json:"peers"`
	Leechers    pgtype.Int4 `json:"leechers"`
	SizeBytes   pgtype.Int8 `json:"size_bytes"`
	Quality     pgtype.Text `json:"quality"`
}

func (q *Queries) UpdateAnimeSerieTorrent(ctx context.Context, arg UpdateAnimeSerieTorrentParams) (AnimeSerieTorrent, error) {
	row := q.db.QueryRow(ctx, updateAnimeSerieTorrent,
		arg.ID,
		arg.LanguageID,
		arg.FileName,
		arg.TorrentHash,
		arg.TorrentFile,
		arg.Seeds,
		arg.Peers,
		arg.Leechers,
		arg.SizeBytes,
		arg.Quality,
	)
	var i AnimeSerieTorrent
	err := row.Scan(
		&i.ID,
		&i.FileName,
		&i.LanguageID,
		&i.TorrentHash,
		&i.TorrentFile,
		&i.Seeds,
		&i.Peers,
		&i.Leechers,
		&i.SizeBytes,
		&i.Quality,
		&i.CreatedAt,
	)
	return i, err
}
