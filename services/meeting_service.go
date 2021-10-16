package services

import (
	"context"
	"qaplagql/graph/model"
	"time"
)

type MeetingService interface {
	Create(ctx context.Context, input model.NewMeeting) (*model.Meeting, error)
	GetById(ctx context.Context, meetingID string) (*model.Meeting, error)
	GetByDate(ctx context.Context, date time.Time) ([]*model.MeetingListItem, error)
}
