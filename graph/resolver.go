package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"qaplagql/graph/model"
	"qaplagql/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users    map[string]*model.User
	projects map[string]*model.Project

	UserService    services.UserService
	ProjectService services.ProjectService
	AuthService    services.AuthService
}
