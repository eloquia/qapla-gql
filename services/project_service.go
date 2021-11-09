package services

import (
	"context"
	"qaplagql/graph/model"
)

type ProjectService interface {
	CreateProject(ctx context.Context, input *model.NewProject) (*model.Project, error)
	GetAll(ctx context.Context) ([]*model.ProjectDetails, error)
	GetById(ctx context.Context, id int) (*model.Project, error)
	GetProjectDetails(ctx context.Context, slug string) (*model.ProjectDetails, error)
	AddUserToProject(ctx context.Context, userId int, projectId int) (*model.User, error)
	GetProjectPersonnel(ctx context.Context, projectID int) ([]*model.User, error)
	GetAssignedProjects(ctx context.Context, userID int) ([]*model.Project, error)
	GetProjectListItems(ctx context.Context) ([]*model.ProjectListItem, error)
	GetAllProjectDetails(ctx context.Context) ([]*model.ProjectDetails, error)
	UpdateProject(ctx context.Context, input model.UpdateProject) (*model.Project, error)
}
