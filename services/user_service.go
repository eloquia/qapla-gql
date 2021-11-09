package services

import (
	"context"
	"qaplagql/graph/model"
)

type UserService interface {
	CreateUser(ctx context.Context, input model.NewUser) (*model.UserDetails, error)
	GetAll(ctx context.Context) ([]*model.UserDetails, error)
	GetById(ctx context.Context, id int) (*model.UserDetails, error)
	UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error)
	GetAllUserDetails(ctx context.Context) ([]*model.UserDetails, error)
	GetAllShortUserDetails(ctx context.Context) ([]*model.UserDetailsShort, error)
	AddPersonnel(ctx context.Context, input model.NewPersonnel) (*model.UserDetails, error)
}
