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

-- name: AddUserDetails :one
INSERT INTO core_qapla.user_details
  (user_id, middle_name, goes_by, gender, ethnicity, position, institution)
VALUES
  ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetUserDetails :one
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
WHERE users.user_id = $1;

-- name: UserExists :one
SELECT EXISTS (
  SELECT 1 FROM core_qapla.users WHERE user_id = $1
);

-- name: IsEmailInUse :one
SELECT EXISTS (
  SELECT 1 FROM core_qapla.users WHERE email = $1
);
