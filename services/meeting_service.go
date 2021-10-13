package services

import (
	"context"
	"qaplagql/graph/model"
)

type MeetingService interface {
	Create(ctx context.Context, input model.NewMeeting) (*model.Meeting, error)
	GetById(ctx context.Context, meetingID string) (*model.Meeting, error)
}
