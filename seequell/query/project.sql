-- name: CreateProject :one
INSERT INTO core_qapla.projects
  (project_name, project_desc)
VALUES
  ($1, $2)
RETURNING *;

-- name: UpdateProject :one
UPDATE core_qapla.projects
SET project_name = $2,
    project_desc = $3
WHERE project_id = $1
RETURNING *;

-- name: AssignUserToProject :one
INSERT INTO core_qapla.project_user 
  (project_id, user_id)
VALUES
  ($1, $2)
RETURNING *;
