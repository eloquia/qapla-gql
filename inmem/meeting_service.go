package inmem

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"qaplagql/graph/model"
	"time"
)

type MeetingServiceInmem struct {
	meetings        map[string]*model.Meeting
	projects        map[string]*model.Project
	projectUser     map[string][]string
	users           map[string]*model.User
	meetingUsers    map[string][]string
	meetingProjects map[string][]string
}

func NewMeetingServiceInmem(meetingMap map[string]*model.Meeting, userMap map[string]*model.User, projectMap map[string]*model.Project, projectUserMap map[string][]string, meetingUserMap map[string][]string, meetingProjectMap map[string][]string) *MeetingServiceInmem {
	return &MeetingServiceInmem{
		meetings:        meetingMap,
		projects:        projectMap,
		projectUser:     projectUserMap,
		users:           userMap,
		meetingUsers:    meetingUserMap,
		meetingProjects: meetingProjectMap,
	}
}

func (u *MeetingServiceInmem) Create(ctx context.Context, input model.NewMeeting) (*model.Meeting, error) {
	meetingID := fmt.Sprintf("%+v", rand.Int())
	meeting := &model.Meeting{
		ID:              meetingID,
		Name:            input.Name,
		StartTime:       input.StartTime,
		DurationMinutes: input.DurationMinutes,
	}

	u.meetings[meetingID] = meeting

	return meeting, nil
}

func (u *MeetingServiceInmem) GetById(ctx context.Context, meetingID string) (*model.Meeting, error) {
	foundMeeting := u.meetings[meetingID]

	if foundMeeting == nil {
		return &model.Meeting{}, errors.New("Meeting with ID not found")
	}

	return foundMeeting, nil
}

func (meetingService *MeetingServiceInmem) GetByDate(ctx context.Context, datetime time.Time) ([]*model.MeetingListItem, error) {
	log.Printf("Getting meetings by date: %+v", datetime)

	/*
		To figure out the meetings within a certain range,
		1. Calculate startDate
			1a. Calculate time.Duration using hour, minute, second, millisecond from day
		2. Calculate endDate
		3. For each meeting, compare with time.After and time.Before
	*/

	h := datetime.Hour()
	hInMs := -time.Hour * time.Duration(h)
	m := datetime.Minute()
	mInMs := -time.Minute * time.Duration(m)
	s := datetime.Second()
	sInMs := -time.Second * time.Duration(s)
	ns := time.Duration(-datetime.Nanosecond())

	start := datetime.Add(hInMs).Add(mInMs).Add(sInMs).Add(ns)
	log.Printf("start %+v", start)

	end := start.Add(time.Hour * time.Duration(24)).Add(-1)
	log.Printf("end %+v", end)

	var meetings []*model.MeetingListItem

	for _, meeting := range meetingService.meetings {

		if meeting.StartTime.After(start) && meeting.StartTime.Before(end) {
			meet := &model.MeetingListItem{
				ID:        meeting.ID,
				Name:      meeting.Name,
				StartTime: meeting.StartTime,
			}
			meetings = append(meetings, meet)
		}
	}

	return meetings, nil
}
