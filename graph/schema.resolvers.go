package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"qaplagql/graph/generated"
	"qaplagql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, email string, password string) (*model.User, error) {
	return r.UserService.CreateUser(ctx, &model.NewUser{
		Email:    email,
		Password: password,
	})
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	return r.UserService.UpdateUser(ctx, &input)
}

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	return r.ProjectService.CreateProject(ctx, &input)
}

func (r *mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProject) (*model.Project, error) {
	// check to see if project with ID exists
	existingProject := r.projects[input.ID]

	if existingProject == nil {
		return nil, errors.New("Project not found")
	}

	project := &model.Project{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}
	r.projects[input.ID] = project

	return project, nil
}

func (r *mutationResolver) AssignUserToProject(ctx context.Context, userID string, projectID string) (*model.User, error) {
	return r.ProjectService.AddUserToProject(ctx, userID, projectID)
}

func (r *mutationResolver) SignIn(ctx context.Context, email string, password string) (*model.User, error) {
	return r.AuthService.SignIn(ctx, email, password)
}

func (r *projectResolver) Personnel(ctx context.Context, obj *model.Project) ([]*model.User, error) {
	return r.ProjectService.GetProjectPersonnel(ctx, obj.ID)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UserService.GetAll()
}

func (r *queryResolver) UserDetails(ctx context.Context, userID string) ([]*model.UserDetails, error) {
	return r.UserService.GetAllUserDetails(ctx)
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return r.UserService.GetById(id)
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	return r.ProjectService.GetAll(ctx)
}

func (r *queryResolver) GetProjectByID(ctx context.Context, id string) (*model.Project, error) {
	return r.ProjectService.GetById(ctx, id)
}

func (r *queryResolver) ProjectDetails(ctx context.Context, slug string) (*model.ProjectDetails, error) {
	return r.ProjectService.GetProjectDetails(ctx, slug)
}

func (r *userDetailsResolver) AssignedProjects(ctx context.Context, obj *model.UserDetails) ([]*model.Project, error) {
	return r.ProjectService.GetAssignedProjects(ctx, obj.ID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// UserDetails returns generated.UserDetailsResolver implementation.
func (r *Resolver) UserDetails() generated.UserDetailsResolver { return &userDetailsResolver{r} }

type mutationResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userDetailsResolver struct{ *Resolver }
