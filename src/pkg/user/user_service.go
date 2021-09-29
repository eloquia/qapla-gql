package user

import "qaplagql/graph/model"

type UserService interface {
	Initialize() error
	CreateUser(firstName string, lastName string, goesBy string, middleName string, email string, gender string, ethnicity string, position string, institution string, isActive bool) (*model.User, error)
}
