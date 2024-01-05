// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: device.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createDevice = `-- name: CreateDevice :one
INSERT INTO devices (
    id,
    device_name,
    device_hash,
    operating_system,
    mac_address,
    client_ip,
    user_agent
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
ON CONFLICT (device_hash)
DO UPDATE SET user_agent = excluded.user_agent
RETURNING id, device_name, device_hash, operating_system, mac_address, client_ip, user_agent, is_banned, created_at
`

type CreateDeviceParams struct {
	ID              uuid.UUID
	DeviceName      string
	DeviceHash      string
	OperatingSystem string
	MacAddress      string
	ClientIp        string
	UserAgent       string
}

func (q *Queries) CreateDevice(ctx context.Context, arg CreateDeviceParams) (Device, error) {
	row := q.db.QueryRow(ctx, createDevice,
		arg.ID,
		arg.DeviceName,
		arg.DeviceHash,
		arg.OperatingSystem,
		arg.MacAddress,
		arg.ClientIp,
		arg.UserAgent,
	)
	var i Device
	err := row.Scan(
		&i.ID,
		&i.DeviceName,
		&i.DeviceHash,
		&i.OperatingSystem,
		&i.MacAddress,
		&i.ClientIp,
		&i.UserAgent,
		&i.IsBanned,
		&i.CreatedAt,
	)
	return i, err
}

const deleteDevice = `-- name: DeleteDevice :exec
DELETE FROM devices
WHERE id = $1
`

func (q *Queries) DeleteDevice(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteDevice, id)
	return err
}

const getDevice = `-- name: GetDevice :one
SELECT id, device_name, device_hash, operating_system, mac_address, client_ip, user_agent, is_banned, created_at FROM devices
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDevice(ctx context.Context, id uuid.UUID) (Device, error) {
	row := q.db.QueryRow(ctx, getDevice, id)
	var i Device
	err := row.Scan(
		&i.ID,
		&i.DeviceName,
		&i.DeviceHash,
		&i.OperatingSystem,
		&i.MacAddress,
		&i.ClientIp,
		&i.UserAgent,
		&i.IsBanned,
		&i.CreatedAt,
	)
	return i, err
}

const updateDevice = `-- name: UpdateDevice :exec
UPDATE devices
SET is_banned = COALESCE($2, is_banned)
WHERE id = $1
`

type UpdateDeviceParams struct {
	ID       uuid.UUID
	IsBanned pgtype.Bool
}

func (q *Queries) UpdateDevice(ctx context.Context, arg UpdateDeviceParams) error {
	_, err := q.db.Exec(ctx, updateDevice, arg.ID, arg.IsBanned)
	return err
}
