package services

import (
	"context"
	"qaplagql/graph/model"
)

type ProjectService interface {
	CreateProject(ctx context.Context, input *model.NewProject) (*model.Project, error)
	GetAll(ctx context.Context) ([]*model.Project, error)
	GetById(ctx context.Context, id string) (*model.Project, error)
	GetProjectDetails(ctx context.Context, slug string) (*model.ProjectDetails, error)
	AddUserToProject(ctx context.Context, userId string, projectId string) (*model.User, error)
	GetProjectPersonnel(ctx context.Context, projectID string) ([]*model.User, error)
	GetAssignedProjects(ctx context.Context, userID string) ([]*model.Project, error)
}
