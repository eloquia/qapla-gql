package seequell

import (
	"context"
	"database/sql"
	"qaplagql/graph/model"
)

type TagServiceSql struct {
	DB      *sql.DB
	queries *Queries
}

func NewTagService(client *sql.DB, queries *Queries) *TagServiceSql {
	return &TagServiceSql{
		DB:      client,
		queries: queries,
	}
}

func (as *TagServiceSql) GetAll(ctx context.Context) ([]*model.MeetingNoteTag, error) {
	panic("Not yet implemented")
}
