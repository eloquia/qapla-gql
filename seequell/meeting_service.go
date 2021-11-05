package seequell

import (
	"context"
	"database/sql"
	"qaplagql/graph/model"
	"time"
)

type MeetingServiceSql struct {
	DB      *sql.DB
	queries *Queries
}

func NewMeetingService(client *sql.DB, queries *Queries) *MeetingServiceSql {
	return &MeetingServiceSql{
		DB:      client,
		queries: queries,
	}
}

func (me *MeetingServiceSql) CreatePersonMeeting(ctx context.Context, input model.NewUserMeeting) (*model.MeetingDetails, error) {
	panic("Not yet implemented")
}

func (me *MeetingServiceSql) CreateProjectMeeting(ctx context.Context, input model.NewProjectMeeting) (*model.MeetingDetails, error) {
	panic("Not yet implemented")
}

func (me *MeetingServiceSql) GetById(ctx context.Context, meetingID string) (*model.MeetingDetails, error) {
	panic("Not yet implemented")
}

func (me *MeetingServiceSql) GetByDate(ctx context.Context, date time.Time) ([]*model.MeetingDetails, error) {
	panic("Not yet implemented")
}

func (me *MeetingServiceSql) UpdatePersonMeeting(ctx context.Context, input model.UpdatedPeopleMeetingDetails) (*model.PeopleMeetingDetails, error) {
	panic("Not yet implemented")
}

func (me *MeetingServiceSql) UpdateMeetingItem(ctx context.Context, input model.UpdateMeetingItemRequest) (*model.MeetingItem, error) {
	panic("Not yet implemented")
}
