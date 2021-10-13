package services

import (
	"context"
	"qaplagql/graph/model"
)

type UserService interface {
	Initialize() error
	CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error)
	GetAll() ([]*model.User, error)
	GetById(id string) (*model.User, error)
	UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error)
	GetAllUserDetails(ctx context.Context) ([]*model.UserDetails, error)
}
