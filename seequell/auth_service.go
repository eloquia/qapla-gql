package seequell

import (
	"context"
	"database/sql"
	"qaplagql/graph/model"
)

type AuthServiceSqlc struct {
	DB      *sql.DB
	queries *Queries
}

func NewAuthService(client *sql.DB, queries *Queries) *AuthServiceSqlc {
	return &AuthServiceSqlc{
		DB:      client,
		queries: queries,
	}
}

func (as *AuthServiceSqlc) SignIn(ctx context.Context, email string, password string) (*model.User, error) {
	panic("Not yet implemented")
}
