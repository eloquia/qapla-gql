package services

import (
	"context"
	"qaplagql/graph/model"
	"time"
)

type MeetingService interface {
	CreatePersonMeeting(ctx context.Context, input model.NewUserMeeting) (*model.MeetingDetails, error)
	CreateProjectMeeting(ctx context.Context, input model.NewProjectMeeting) (*model.MeetingDetails, error)
	GetById(ctx context.Context, meetingID int) (*model.MeetingDetails, error)
	GetByDate(ctx context.Context, date time.Time) ([]*model.MeetingDetails, error)
	UpdatePersonMeeting(ctx context.Context, input model.UpdatedPeopleMeetingDetails) (*model.PeopleMeetingDetails, error)
	UpdateMeetingItem(ctx context.Context, input model.UpdateMeetingItemRequest) (*model.MeetingItem, error)
}
