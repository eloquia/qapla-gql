package services

import (
	"context"
	"qaplagql/graph/model"
)

type TagService interface {
	GetAll(ctx context.Context) ([]*model.MeetingNoteTag, error)
}
