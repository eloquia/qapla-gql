package services

import (
	"context"
	"qaplagql/graph/model"
)

type AuthService interface {
	SignIn(ctx context.Context, email string, password string) (*model.User, error)
}
