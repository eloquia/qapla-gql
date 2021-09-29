package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"qaplagql/graph/generated"
	"qaplagql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		GoesBy:      input.GoesBy,
		MiddleName:  input.MiddleName,
		Email:       input.Email,
		Gender:      input.Gender,
		Ethnicity:   input.Ethnicity,
		Position:    input.Position,
		Institution: input.Institution,
		IsActive:    input.IsActive,
	}
	proposedId := fmt.Sprintf("T%d", rand.Int())
	// keep generating new IDs until a unique one is created
	for {
		if r.users[proposedId] != nil {
			proposedId = fmt.Sprintf("T%d", rand.Int())
		}
	}
	user.ID = proposedId
	r.users[proposedId] = user
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	project := &model.Project{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: input.Description,
	}
	proposedId := fmt.Sprintf("T%d", rand.Int())
	// keep generating new IDs until a unique one is created
	for {
		if r.projects[proposedId] != nil {
			proposedId = fmt.Sprintf("T%d", rand.Int())
		}
	}
	project.ID = proposedId
	r.projects[proposedId] = project
	return project, nil
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
		Description: *input.Description,
	}
	r.projects[input.ID] = project

	return project, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 0, len(r.users))
	for _, value := range r.users {
		users = append(users, value)
	}
	return users, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	projects := make([]*model.Project, 0, len(r.projects))
	for _, value := range r.projects {
		projects = append(projects, value)
	}
	return projects, nil
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
type projectResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
