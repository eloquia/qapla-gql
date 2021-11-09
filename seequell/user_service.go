package seequell

import (
	"context"
	"database/sql"
	"errors"
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

func (us *UserServiceSql) CreateUser(ctx context.Context, input model.NewUser) (*model.UserDetailsShort, error) {
	log.Printf("[DEBUG] CreateUser")

	arg := CreateUserParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}
	user, err := us.queries.CreateUser(ctx, arg)
	if err != nil {
		return &model.UserDetailsShort{}, err
	}

	// transform
	u := UserToShortDomain(user)

	return u, nil
}

func (us *UserServiceSql) GetById(ctx context.Context, id int) (*model.UserDetails, error) {
	log.Printf("[DEBUG] Get User By ID")

	id64 := int64(id)

	modelUser, err := us.queries.GetUserDetails(ctx, id64)
	if err != nil {
		return &model.UserDetails{}, err
	}

	domainUser := UserToDetailedDomain(modelUser)

	return domainUser, nil
}

func (us *UserServiceSql) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	panic("Not yet implemented")
}

func (us *UserServiceSql) GetAll(ctx context.Context) ([]*model.UserDetailsShort, error) {
	log.Printf("[DEBUG] GetAll Users")

	modelUsers, err := us.queries.ListUsers(ctx)
	if err != nil {
		return []*model.UserDetailsShort{}, err
	}

	var domainUsers []*model.UserDetailsShort
	for _, modelUser := range modelUsers {
		domainUser := UserToShortDomain(modelUser)
		domainUsers = append(domainUsers, domainUser)
	}

	return domainUsers, nil
}

func (u *UserServiceSql) GetAllShortUserDetails(ctx context.Context) ([]*model.UserDetailsShort, error) {
	panic("Not yet implemented")
}

func (us *UserServiceSql) AddUserDetails(ctx context.Context, input model.UserDetailsInput) (*model.UserDetails, error) {
	log.Printf("[DEBUG] Add User Details")

	// check to see if user with ID exists
	res, err := us.queries.UserExists(ctx, int64(input.UserID))
	if err != nil {
		return &model.UserDetails{}, err
	}
	if !res {
		return &model.UserDetails{}, errors.New("User with ID does not exist")
	}

	// user with ID exists, so we can create a new entry
	params := AddUserDetailsParams{
		UserID: int64(input.UserID),
		MiddleName: sql.NullString{
			Valid:  true,
			String: *input.MiddleName,
		},
		GoesBy: sql.NullString{
			Valid:  true,
			String: *input.GoesBy,
		},
		Gender: sql.NullString{
			Valid:  true,
			String: *input.Gender,
		},
		Ethnicity: sql.NullString{
			Valid:  true,
			String: *input.Ethnicity,
		},
		Position: sql.NullString{
			Valid:  true,
			String: *input.Position,
		},
		Institution: sql.NullString{
			Valid:  true,
			String: *input.Institution,
		},
	}
	deets, err := us.queries.AddUserDetails(ctx, params)
	if err != nil {
		return &model.UserDetails{}, err
	}

	detailedDomain := ToUserDetailsDomain(deets)
	return detailedDomain, nil
}

func (us *UserServiceSql) IsEmailInUse(ctx context.Context, email string) (bool, error) {
	log.Printf("[DEBUG] IsEmailInUse")

	isInUse, err := us.queries.IsEmailInUse(ctx, email)
	if err != nil {
		return true, err
	}

	return isInUse, nil
}
