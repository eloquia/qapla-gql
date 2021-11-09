package seequell

import (
	"context"
	"database/sql"
	"log"
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

func (us *UserServiceSql) CreateUser(ctx context.Context, input model.NewUser) (*model.UserDetails, error) {
	log.Printf("[DEBUG] CreateUsers")

	arg := CreateUserParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}
	user, err := us.queries.CreateUser(ctx, arg)
	if err != nil {
		return &model.UserDetails{}, err
	}

	// transform
	u := UserToDomain(user)

	return u, nil
}

func (us *UserServiceSql) GetById(ctx context.Context, id int) (*model.UserDetails, error) {
	log.Printf("[DEBUG] GetAll Users")

	id64 := int64(id)

	modelUser, err := us.queries.GetUser(ctx, id64)
	if err != nil {
		return &model.UserDetails{}, err
	}

	domainUser := UserToDomain(modelUser)

	return domainUser, nil
}

func (us *UserServiceSql) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	panic("Not yet implemented")
}

func (us *UserServiceSql) GetAll(ctx context.Context) ([]*model.UserDetails, error) {
	log.Printf("[DEBUG] GetAll Users")

	modelUsers, err := us.queries.ListUsers(ctx)
	if err != nil {
		return []*model.UserDetails{}, err
	}

	var domainUsers []*model.UserDetails
	for _, modelUser := range modelUsers {
		domainUser := UserToDomain(modelUser)
		domainUsers = append(domainUsers, domainUser)
	}

	return domainUsers, nil
}

func (us *UserServiceSql) GetAllUserDetails(ctx context.Context) ([]*model.UserDetails, error) {
	panic("Not yet implemented")
}

func (u *UserServiceSql) AddPersonnel(ctx context.Context, input model.NewPersonnel) (*model.UserDetails, error) {
	panic("Not yet implemented")
}

func (u *UserServiceSql) GetAllShortUserDetails(ctx context.Context) ([]*model.UserDetailsShort, error) {
	panic("Not yet implemented")
}
