// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package seequell

import (
	"context"
	"database/sql"
)

const addUserDetails = `-- name: AddUserDetails :one
INSERT INTO core_qapla.user_details
  (user_id, middle_name, goes_by, gender, ethnicity, position, institution)
VALUES
  ($1, $2, $3, $4, $5, $6, $7)
RETURNING user_detail_id, user_id, middle_name, goes_by, gender, ethnicity, position, institution
`

type AddUserDetailsParams struct {
	UserID      int64          `json:"user_id"`
	MiddleName  sql.NullString `json:"middle_name"`
	GoesBy      sql.NullString `json:"goes_by"`
	Gender      sql.NullString `json:"gender"`
	Ethnicity   sql.NullString `json:"ethnicity"`
	Position    sql.NullString `json:"position"`
	Institution sql.NullString `json:"institution"`
}

func (q *Queries) AddUserDetails(ctx context.Context, arg AddUserDetailsParams) (CoreQaplaUserDetail, error) {
	row := q.queryRow(ctx, q.addUserDetailsStmt, addUserDetails,
		arg.UserID,
		arg.MiddleName,
		arg.GoesBy,
		arg.Gender,
		arg.Ethnicity,
		arg.Position,
		arg.Institution,
	)
	var i CoreQaplaUserDetail
	err := row.Scan(
		&i.UserDetailID,
		&i.UserID,
		&i.MiddleName,
		&i.GoesBy,
		&i.Gender,
		&i.Ethnicity,
		&i.Position,
		&i.Institution,
	)
	return i, err
}

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

const getUserDetails = `-- name: GetUserDetails :one
SELECT users.user_id,
       users.first_name,
       users.last_name,
       users.email,
       details.user_detail_id,
       details.middle_name,
       details.goes_by,
       details.gender,
       details.ethnicity,
       details.position,
       details.institution
FROM core_qapla.user_details details
JOIN core_qapla.users users
ON users.user_id = details.user_id
WHERE users.user_id = $1
`

type GetUserDetailsRow struct {
	UserID       int64          `json:"user_id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Email        string         `json:"email"`
	UserDetailID int64          `json:"user_detail_id"`
	MiddleName   sql.NullString `json:"middle_name"`
	GoesBy       sql.NullString `json:"goes_by"`
	Gender       sql.NullString `json:"gender"`
	Ethnicity    sql.NullString `json:"ethnicity"`
	Position     sql.NullString `json:"position"`
	Institution  sql.NullString `json:"institution"`
}

func (q *Queries) GetUserDetails(ctx context.Context, userID int64) (GetUserDetailsRow, error) {
	row := q.queryRow(ctx, q.getUserDetailsStmt, getUserDetails, userID)
	var i GetUserDetailsRow
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.UserDetailID,
		&i.MiddleName,
		&i.GoesBy,
		&i.Gender,
		&i.Ethnicity,
		&i.Position,
		&i.Institution,
	)
	return i, err
}

const isEmailInUse = `-- name: IsEmailInUse :one
SELECT EXISTS (
  SELECT 1 FROM core_qapla.users WHERE email = $1
)
`

func (q *Queries) IsEmailInUse(ctx context.Context, email string) (bool, error) {
	row := q.queryRow(ctx, q.isEmailInUseStmt, isEmailInUse, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
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

const updateUser = `-- name: UpdateUser :one
UPDATE core_qapla.users
SET first_name = $2,
    last_name = $3,
    email = $4
WHERE user_id = $1
RETURNING user_id, first_name, last_name, email, created_at, updated_at
`

type UpdateUserParams struct {
	UserID    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (CoreQaplaUser, error) {
	row := q.queryRow(ctx, q.updateUserStmt, updateUser,
		arg.UserID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
	)
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

const userExists = `-- name: UserExists :one
SELECT EXISTS (
  SELECT 1 FROM core_qapla.users WHERE user_id = $1
)
`

func (q *Queries) UserExists(ctx context.Context, userID int64) (bool, error) {
	row := q.queryRow(ctx, q.userExistsStmt, userExists, userID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
