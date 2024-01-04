// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: user_device.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUserDevice = `-- name: CreateUserDevice :one
INSERT INTO user_devices (
    user_id,
    device_id
) VALUES (
  $1, $2
)
RETURNING id, user_id, device_id
`

type CreateUserDeviceParams struct {
	UserID   int64
	DeviceID uuid.UUID
}

func (q *Queries) CreateUserDevice(ctx context.Context, arg CreateUserDeviceParams) (UserDevice, error) {
	row := q.db.QueryRow(ctx, createUserDevice, arg.UserID, arg.DeviceID)
	var i UserDevice
	err := row.Scan(&i.ID, &i.UserID, &i.DeviceID)
	return i, err
}

const listUserDevices = `-- name: ListUserDevices :many
SELECT device_id FROM user_devices
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListUserDevicesParams struct {
	UserID int64
	Limit  int32
	Offset int32
}

func (q *Queries) ListUserDevices(ctx context.Context, arg ListUserDevicesParams) ([]uuid.UUID, error) {
	rows, err := q.db.Query(ctx, listUserDevices, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []uuid.UUID{}
	for rows.Next() {
		var device_id uuid.UUID
		if err := rows.Scan(&device_id); err != nil {
			return nil, err
		}
		items = append(items, device_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
