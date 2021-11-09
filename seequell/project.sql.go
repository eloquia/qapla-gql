// Code generated by sqlc. DO NOT EDIT.
// source: project.sql

package seequell

import (
	"context"
	"database/sql"
)

const assignUserToProject = `-- name: AssignUserToProject :one
INSERT INTO core_qapla.project_user 
  (project_id, user_id)
VALUES
  ($1, $2)
RETURNING user_id, project_id
`

type AssignUserToProjectParams struct {
	ProjectID int64 `json:"project_id"`
	UserID    int64 `json:"user_id"`
}

func (q *Queries) AssignUserToProject(ctx context.Context, arg AssignUserToProjectParams) (CoreQaplaProjectUser, error) {
	row := q.queryRow(ctx, q.assignUserToProjectStmt, assignUserToProject, arg.ProjectID, arg.UserID)
	var i CoreQaplaProjectUser
	err := row.Scan(&i.UserID, &i.ProjectID)
	return i, err
}

const createProject = `-- name: CreateProject :one
INSERT INTO core_qapla.projects
  (project_name, project_desc)
VALUES
  ($1, $2)
RETURNING project_id, project_name, project_desc, created_at, updated_at, slug
`

type CreateProjectParams struct {
	ProjectName string         `json:"project_name"`
	ProjectDesc sql.NullString `json:"project_desc"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (CoreQaplaProject, error) {
	row := q.queryRow(ctx, q.createProjectStmt, createProject, arg.ProjectName, arg.ProjectDesc)
	var i CoreQaplaProject
	err := row.Scan(
		&i.ProjectID,
		&i.ProjectName,
		&i.ProjectDesc,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Slug,
	)
	return i, err
}

const getAllProjects = `-- name: GetAllProjects :many
SELECT project_id, project_name, project_desc, created_at, updated_at, slug FROM core_qapla.projects
ORDER BY project_name
`

func (q *Queries) GetAllProjects(ctx context.Context) ([]CoreQaplaProject, error) {
	rows, err := q.query(ctx, q.getAllProjectsStmt, getAllProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CoreQaplaProject{}
	for rows.Next() {
		var i CoreQaplaProject
		if err := rows.Scan(
			&i.ProjectID,
			&i.ProjectName,
			&i.ProjectDesc,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Slug,
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

const getProjectById = `-- name: GetProjectById :one
SELECT project_id, project_name, project_desc, created_at, updated_at, slug FROM core_qapla.projects
WHERE project_id = $1
`

func (q *Queries) GetProjectById(ctx context.Context, projectID int64) (CoreQaplaProject, error) {
	row := q.queryRow(ctx, q.getProjectByIdStmt, getProjectById, projectID)
	var i CoreQaplaProject
	err := row.Scan(
		&i.ProjectID,
		&i.ProjectName,
		&i.ProjectDesc,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Slug,
	)
	return i, err
}

const getProjectBySlug = `-- name: GetProjectBySlug :one
SELECT project_id, project_name, project_desc, created_at, updated_at, slug FROM core_qapla.projects
WHERE slug = $1
`

func (q *Queries) GetProjectBySlug(ctx context.Context, slug string) (CoreQaplaProject, error) {
	row := q.queryRow(ctx, q.getProjectBySlugStmt, getProjectBySlug, slug)
	var i CoreQaplaProject
	err := row.Scan(
		&i.ProjectID,
		&i.ProjectName,
		&i.ProjectDesc,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Slug,
	)
	return i, err
}

const projectExistsById = `-- name: ProjectExistsById :one
SELECT EXISTS (
  SELECT 1 FROM core_qapla.projects WHERE project_id = $1
)
`

func (q *Queries) ProjectExistsById(ctx context.Context, projectID int64) (bool, error) {
	row := q.queryRow(ctx, q.projectExistsByIdStmt, projectExistsById, projectID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const projectExistsBySlug = `-- name: ProjectExistsBySlug :one
SELECT EXISTS (
  SELECT 1 FROM core_qapla.projects WHERE slug = $1
)
`

func (q *Queries) ProjectExistsBySlug(ctx context.Context, slug string) (bool, error) {
	row := q.queryRow(ctx, q.projectExistsBySlugStmt, projectExistsBySlug, slug)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const updateProject = `-- name: UpdateProject :one
UPDATE core_qapla.projects
SET project_name = $2,
    project_desc = $3
WHERE project_id = $1
RETURNING project_id, project_name, project_desc, created_at, updated_at, slug
`

type UpdateProjectParams struct {
	ProjectID   int64          `json:"project_id"`
	ProjectName string         `json:"project_name"`
	ProjectDesc sql.NullString `json:"project_desc"`
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (CoreQaplaProject, error) {
	row := q.queryRow(ctx, q.updateProjectStmt, updateProject, arg.ProjectID, arg.ProjectName, arg.ProjectDesc)
	var i CoreQaplaProject
	err := row.Scan(
		&i.ProjectID,
		&i.ProjectName,
		&i.ProjectDesc,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Slug,
	)
	return i, err
}
