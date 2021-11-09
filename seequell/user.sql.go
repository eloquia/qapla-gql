// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package seequell

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO core_qapla.users (
  first_name, last_name, email
) VALUES (
  $1, $2, $3
)
RETURNING user_id, first_name, last_name, email, created_at, updated_at
`

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CoreQaplaUser, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.FirstName, arg.LastName, arg.Email)
	var i CoreQaplaUser
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM core_qapla.users
WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID int64) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, userID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, first_name, last_name, email, created_at, updated_at FROM core_qapla.users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID int64) (CoreQaplaUser, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, userID)
	var i CoreQaplaUser
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, first_name, last_name, email, created_at, updated_at FROM core_qapla.users
ORDER BY last_name
`

func (q *Queries) ListUsers(ctx context.Context) ([]CoreQaplaUser, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CoreQaplaUser{}
	for rows.Next() {
		var i CoreQaplaUser
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}