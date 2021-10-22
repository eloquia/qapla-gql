package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"qaplagql/graph/generated"
	"qaplagql/graph/model"
	"time"
)

func (r *mutationResolver) CreateUser(ctx context.Context, email string, password string) (*model.User, error) {
	return r.UserService.CreateUser(ctx, &model.NewUser{
		Email:    email,
		Password: password,
	})
}

func (r *mutationResolver) CreatePersonnel(ctx context.Context, input model.NewPersonnel) (*model.UserDetails, error) {
	return r.UserService.AddPersonnel(ctx, input)
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

func (r *mutationResolver) CreateUserMeeting(ctx context.Context, input model.NewUserMeeting) (*model.MeetingDetails, error) {
	return r.MeetingService.CreatePersonMeeting(ctx, input)
}

func (r *mutationResolver) CreateProjectMeeting(ctx context.Context, input model.NewProjectMeeting) (*model.MeetingDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Personnel(ctx context.Context, obj *model.Project) ([]*model.User, error) {
	return r.ProjectService.GetProjectPersonnel(ctx, obj.ID)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UserService.GetAll()
}

func (r *queryResolver) UserDetails(ctx context.Context) ([]*model.UserDetails, error) {
	return r.UserService.GetAllUserDetails(ctx)
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return r.UserService.GetById(id)
}

func (r *queryResolver) ProjectListItems(ctx context.Context) ([]*model.ProjectListItem, error) {
	return r.ProjectService.GetProjectListItems(ctx)
}

func (r *queryResolver) GetProjectByID(ctx context.Context, id string) (*model.Project, error) {
	return r.ProjectService.GetById(ctx, id)
}

func (r *queryResolver) ProjectDetails(ctx context.Context, slug string) (*model.ProjectDetails, error) {
	return r.ProjectService.GetProjectDetails(ctx, slug)
}

func (r *queryResolver) ProjectDetailsList(ctx context.Context) ([]*model.ProjectDetails, error) {
	return r.ProjectService.GetAllProjectDetails(ctx)
}

func (r *queryResolver) MeetingByID(ctx context.Context, id string) (*model.MeetingDetails, error) {
	return r.MeetingService.GetById(ctx, id)
}

func (r *queryResolver) MeetingsByDate(ctx context.Context, date time.Time) ([]*model.MeetingDetails, error) {
	return r.MeetingService.GetByDate(ctx, date)
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.MeetingNoteTag, error) {
	return r.TagService.GetAll(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
