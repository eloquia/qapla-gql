package seequell

import (
	"context"
	"database/sql"
	"qaplagql/graph/model"
)

type UserServiceSql struct {
	DB      *sql.DB
	queries *Queries
}

func NewUserService(client *sql.DB, queries *Queries) *UserServiceSql {
	return &UserServiceSql{
		DB:      client,
		queries: queries,
	}
}

func (us *UserServiceSql) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	panic("Not yet implemented")
}

func (us *UserServiceSql) GetById(id string) (*model.User, error) {
	panic("Not yet implemented")
}

func (userService *UserServiceSql) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	panic("Not yet implemented")
}

func (userService *UserServiceSql) GetAll() ([]*model.User, error) {
	panic("Not yet implemented")
}

func (userService *UserServiceSql) GetAllUserDetails(ctx context.Context) ([]*model.UserDetails, error) {
	panic("Not yet implemented")
}

func (u *UserServiceSql) AddPersonnel(ctx context.Context, input model.NewPersonnel) (*model.UserDetails, error) {
	panic("Not yet implemented")
}
