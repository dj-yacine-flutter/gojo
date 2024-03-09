// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
          username, 
          email, 
          hashed_password, 
          full_name
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, username, email, hashed_password, full_name, password_changed_at, created_at, is_email_verified, role
`

type CreateUserParams struct {
	Username       string
	Email          string
	HashedPassword string
	FullName       string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
		arg.FullName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.FullName,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
		&i.Role,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, hashed_password, full_name, password_changed_at, created_at, is_email_verified, role FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.FullName,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
		&i.Role,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, hashed_password, full_name, password_changed_at, created_at, is_email_verified, role FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.FullName,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
		&i.Role,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, hashed_password, full_name, password_changed_at, created_at, is_email_verified, role FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.HashedPassword,
			&i.FullName,
			&i.PasswordChangedAt,
			&i.CreatedAt,
			&i.IsEmailVerified,
			&i.Role,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE($1, hashed_password),
  password_changed_at = COALESCE($2, password_changed_at),
  full_name = COALESCE($3, full_name),
  email = COALESCE($4, email),
  is_email_verified = COALESCE($5, is_email_verified)
WHERE
  username = $6
RETURNING id, username, email, hashed_password, full_name, password_changed_at, created_at, is_email_verified, role
`

type UpdateUserParams struct {
	HashedPassword    pgtype.Text
	PasswordChangedAt pgtype.Timestamptz
	FullName          pgtype.Text
	Email             pgtype.Text
	IsEmailVerified   pgtype.Bool
	Username          string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.FullName,
		arg.Email,
		arg.IsEmailVerified,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.FullName,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
		&i.Role,
	)
	return i, err
}
