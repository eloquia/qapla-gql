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

-- name: GetAllProjects :many
SELECT * FROM core_qapla.projects
ORDER BY project_name;

-- name: GetProjectById :one
SELECT * FROM core_qapla.projects
WHERE project_id = $1;

-- name: ProjectExistsById :one
SELECT EXISTS (
  SELECT 1 FROM core_qapla.projects WHERE project_id = $1
);

-- name: GetProjectBySlug :one
SELECT * FROM core_qapla.projects
WHERE slug = $1;

-- name: ProjectExistsBySlug :one
SELECT EXISTS (
  SELECT 1 FROM core_qapla.projects WHERE slug = $1
);
