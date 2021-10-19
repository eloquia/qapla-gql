package services

import (
	"context"
	"qaplagql/graph/model"
	"time"
)

type MeetingService interface {
	CreatePersonMeeting(ctx context.Context, input model.NewUserMeeting) (*model.Meeting, error)
	CreateProjectMeeting(ctx context.Context, input model.NewProjectMeeting) (*model.Meeting, error)
	GetById(ctx context.Context, meetingID string) (*model.Meeting, error)
	GetByDate(ctx context.Context, date time.Time) ([]*model.MeetingDetails, error)
}
