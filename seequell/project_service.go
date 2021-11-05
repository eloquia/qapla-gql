package seequell

import (
	"context"
	"database/sql"
	"qaplagql/graph/model"
)

type ProjectServiceSql struct {
	DB      *sql.DB
	queries *Queries
}

func NewProjectService(client *sql.DB, queries *Queries) *ProjectServiceSql {
	return &ProjectServiceSql{
		DB:      client,
		queries: queries,
	}
}

func (ps *ProjectServiceSql) CreateProject(ctx context.Context, input *model.NewProject) (*model.Project, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetAll(ctx context.Context) ([]*model.ProjectDetails, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetById(ctx context.Context, id string) (*model.Project, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetProjectDetails(ctx context.Context, slug string) (*model.ProjectDetails, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) AddUserToProject(ctx context.Context, userId string, projectId string) (*model.User, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetProjectPersonnel(ctx context.Context, projectID string) ([]*model.User, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetAssignedProjects(ctx context.Context, userID string) ([]*model.Project, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetProjectListItems(ctx context.Context) ([]*model.ProjectListItem, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetAllProjectDetails(ctx context.Context) ([]*model.ProjectDetails, error) {
	panic("Not yet implemented")
}
