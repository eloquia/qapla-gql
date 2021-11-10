package services

import (
	"context"
	"qaplagql/graph/model"
)

type UserService interface {
	// Create
	CreateUser(ctx context.Context, input model.NewUser) (*model.UserDetailsShort, error)

	// Read
	GetAll(ctx context.Context) ([]*model.UserDetailsShort, error)
	GetById(ctx context.Context, id int) (*model.UserDetails, error)
	GetAllShortUserDetails(ctx context.Context) ([]*model.UserDetailsShort, error)
	IsEmailInUse(ctx context.Context, email string) (bool, error)

	// Update
	UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error)
	AddUserDetails(ctx context.Context, input model.UserDetailsInput) (*model.UserDetails, error)
	UpdateUserDetails(ctx context.Context, input model.UserDetailsInput) (*model.UserDetails, error)

	// Delete

}
