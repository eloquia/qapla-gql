package seequell

import (
	"context"
	"database/sql"
	"log"
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

func (ps *ProjectServiceSql) GetById(ctx context.Context, id int) (*model.Project, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetProjectDetails(ctx context.Context, slug string) (*model.ProjectDetails, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) AddUserToProject(ctx context.Context, userId int, projectId int) (*model.User, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetProjectPersonnel(ctx context.Context, projectID int) ([]*model.User, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetAssignedProjects(ctx context.Context, userID int) ([]*model.Project, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) GetProjectListItems(ctx context.Context) ([]*model.ProjectListItem, error) {
	log.Printf("[DEBUG] GetProjectListItems")

	projectModels, err := ps.queries.GetAllProjects(ctx)
	if err != nil {
		return []*model.ProjectListItem{}, err
	}

	var projects []*model.ProjectListItem
	for i := 0; i < len(projectModels); i++ {
		project := ProjectToListItemDomain(projectModels[i])
		projects = append(projects, project)
	}

	return projects, nil
}

func (ps *ProjectServiceSql) GetAllProjectDetails(ctx context.Context) ([]*model.ProjectDetails, error) {
	panic("Not yet implemented")
}

func (ps *ProjectServiceSql) UpdateProject(ctx context.Context, input model.UpdateProject) (*model.Project, error) {
	panic("Not yet implemented")
}
