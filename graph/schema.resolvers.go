package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"qaplagql/graph/generated"
	"qaplagql/graph/model"
	"time"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.UserDetailsShort, error) {
	return r.UserService.CreateUser(ctx, input)
}

func (r *mutationResolver) CreateAuth(ctx context.Context, input model.NewUserAuth) (*model.UserDetailsShort, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePersonnel(ctx context.Context, input model.NewPersonnel) (*model.UserDetailsShort, error) {
	return r.UserService.AddPersonnel(ctx, input)
}

func (r *mutationResolver) AddUserDetails(ctx context.Context, input model.UserDetailsInput) (*model.UserDetails, error) {
	return r.UserService.AddUserDetails(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	return r.UserService.UpdateUser(ctx, &input)
}

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	return r.ProjectService.CreateProject(ctx, &input)
}

func (r *mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProject) (*model.Project, error) {
	return r.ProjectService.UpdateProject(ctx, input)
}

func (r *mutationResolver) AssignUserToProject(ctx context.Context, userID int, projectID int) (*model.User, error) {
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

func (r *mutationResolver) UpdateUserMeeting(ctx context.Context, input model.UpdatedPeopleMeetingDetails) (*model.PeopleMeetingDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMeetingItem(ctx context.Context, input model.UpdateMeetingItemRequest) (*model.MeetingItem, error) {
	return r.MeetingService.UpdateMeetingItem(ctx, input)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.UserDetailsShort, error) {
	return r.UserService.GetAll(ctx)
}

func (r *queryResolver) GetUserByID(ctx context.Context, id int) (*model.UserDetails, error) {
	return r.UserService.GetById(ctx, id)
}

func (r *queryResolver) IsEmailInUse(ctx context.Context, email string) (bool, error) {
	return r.UserService.IsEmailInUse(ctx, email)
}

func (r *queryResolver) ProjectListItems(ctx context.Context) ([]*model.ProjectListItem, error) {
	return r.ProjectService.GetProjectListItems(ctx)
}

func (r *queryResolver) GetProjectByID(ctx context.Context, id int) (*model.Project, error) {
	return r.ProjectService.GetById(ctx, id)
}

func (r *queryResolver) ProjectDetails(ctx context.Context, slug string) (*model.ProjectDetails, error) {
	return r.ProjectService.GetProjectDetails(ctx, slug)
}

func (r *queryResolver) MeetingByID(ctx context.Context, id int) (*model.MeetingDetails, error) {
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) ProjectDetailsList(ctx context.Context) ([]*model.ProjectDetails, error) {
	return r.ProjectService.GetAllProjectDetails(ctx)
}
func (r *queryResolver) UserDetails(ctx context.Context) ([]*model.UserDetailsShort, error) {
	return r.UserService.GetAllShortUserDetails(ctx)
}
