package inmem

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"qaplagql/graph/model"
)

type UserServiceInmem struct {
	projects    map[string]*model.Project
	projectUser map[string][]string
	users       map[string]*model.User
}

func NewUserServiceInmem(userMap map[string]*model.User, projectMap map[string]*model.Project, projectUserMap map[string][]string) *UserServiceInmem {
	return &UserServiceInmem{
		users:       userMap,
		projects:    projectMap,
		projectUser: projectUserMap,
	}
}

func (userService *UserServiceInmem) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	// check to see if user with email exists
	isEmailUsed := false
	for _, user := range userService.users {
		if user.Email == input.Email {
			isEmailUsed = true
		}
	}
	if isEmailUsed {
		return nil, errors.New("Email address already in use")
	}

	// create user
	user := &model.User{
		ID:       fmt.Sprintf("%+v", rand.Int()),
		Email:    input.Email,
		Password: input.Password,
	}

	userService.users[user.ID] = user
	return user, nil
}

func (userService *UserServiceInmem) GetById(id string) (*model.User, error) {
	var foundUser *model.User
	for userId, user := range userService.users {
		if userId == id {
			foundUser = user
		}
	}

	if foundUser == nil {
		return &model.User{}, errors.New("User with ID does not exist")
	}

	return foundUser, nil
}

func (userService *UserServiceInmem) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	_, err := userService.GetById(input.ID)
	if err != nil {
		return &model.User{}, err
	}

	user := &model.User{
		ID:          input.ID,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		GoesBy:      input.GoesBy,
		MiddleName:  input.MiddleName,
		Gender:      input.Gender,
		Ethnicity:   input.Ethnicity,
		Position:    input.Position,
		Institution: input.Institution,
	}

	userService.users[user.ID] = user

	return user, nil
}

func (userService *UserServiceInmem) GetAll() ([]*model.User, error) {
	var users []*model.User
	for _, user := range userService.users {
		users = append(users, user)
	}

	return users, nil
}

func (userService *UserServiceInmem) GetAllUserDetails(ctx context.Context) ([]*model.UserDetails, error) {
	var userDetails []*model.UserDetails
	for _, user := range userService.users {
		userDetail := &model.UserDetails{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}

		// fetch projects associated with User
		var assignedProjectIDs []string
		for projectID, assignedUserIDs := range userService.projectUser {
			// figure out if user is assigned to the project
			for _, assignedUserID := range assignedUserIDs {
				if assignedUserID == userDetail.ID {
					assignedProjectIDs = append(assignedProjectIDs, projectID)
				}
			}
		}

		var assignedProjects []*model.Project
		for _, pID := range assignedProjectIDs {
			project := userService.projects[pID]

			assignedProjects = append(assignedProjects, project)
		}

		userDetail.AssignedProjects = assignedProjects

		userDetails = append(userDetails, userDetail)
	}

	return userDetails, nil
}
