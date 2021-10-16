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

func (m *MeetingServiceInmem) GetByDate(ctx context.Context, datetime time.Time) ([]*model.Meeting, error) {
	log.Printf("Getting meetings by date: %+v", datetime)

	year := datetime.Year()
	month := datetime.Month()
	day := datetime.Day()
	// tz := datetime.z
	log.Printf("year: %+v, month: %+v, day: %+v", year, month, day)

	var meetings []*model.Meeting

	for _, meeting := range m.meetings {
		log.Printf("meeting: %+v", meeting)
		cy := meeting.StartTime.Year()
		cm := meeting.StartTime.Month()
		cd := meeting.StartTime.Day()

		if cy == year && cm == month && cd == day {
			meetings = append(meetings, meeting)
		}
	}

	return meetings, nil
}
