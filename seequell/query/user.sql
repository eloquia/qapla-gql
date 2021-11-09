-- name: GetUser :one
SELECT * FROM core_qapla.users
WHERE user_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM core_qapla.users
ORDER BY last_name;

-- name: CreateUser :one
INSERT INTO core_qapla.users (
  first_name, last_name, email
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM core_qapla.users
WHERE user_id = $1;