package seequell

import (
	"context"
	"database/sql"
	"errors"
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
	log.Printf("[DEBUG] GetProjectDetails")

	exists, err := ps.queries.ProjectExists(ctx, int64(id))
	if err != nil {
		return &model.Project{}, err
	}
	if !exists {
		return &model.Project{}, errors.New("Project with ID does not exist")
	}

	projectModel, err := ps.queries.GetProjectById(ctx, int64(id))
	if err != nil {
		return &model.Project{}, err
	}

	projectDomain := ProjectToDomain(projectModel)

	return projectDomain, nil
}

func (ps *ProjectServiceSql) GetProjectDetails(ctx context.Context, slug string) (*model.ProjectDetails, error) {
	log.Printf("[DEBUG] GetProjectDetails")

	if slug == "" {
		return &model.ProjectDetails{}, errors.New("Slug is required")
	}

	projectModel, err := ps.queries.GetProjectBySlug(ctx, slug)
	if err != nil {
		return &model.ProjectDetails{}, err
	}
	projectDomain := ProjectToDetailsDomain(projectModel)

	return projectDomain, nil
}

func (ps *ProjectServiceSql) AddUserToProject(ctx context.Context, userId int, projectId int) (*model.User, error) {
	log.Printf("[DEBUG] AddUserToProject")

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
