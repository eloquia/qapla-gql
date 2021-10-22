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
	meetings        map[string]*model.MeetingDetails
	projects        map[string]*model.Project
	projectUser     map[string][]string
	users           map[string]*model.User
	meetingUsers    map[string][]string
	meetingProjects map[string][]string
	tags            map[string]*model.MeetingNoteTag
}

func NewMeetingServiceInmem(meetingMap map[string]*model.MeetingDetails, userMap map[string]*model.User, projectMap map[string]*model.Project, projectUserMap map[string][]string, meetingUserMap map[string][]string, meetingProjectMap map[string][]string, tagMap map[string]*model.MeetingNoteTag) *MeetingServiceInmem {
	return &MeetingServiceInmem{
		meetings:        meetingMap,
		projects:        projectMap,
		projectUser:     projectUserMap,
		users:           userMap,
		meetingUsers:    meetingUserMap,
		meetingProjects: meetingProjectMap,
		tags:            tagMap,
	}
}

func (u *MeetingServiceInmem) CreatePersonMeeting(ctx context.Context, input model.NewUserMeeting) (*model.MeetingDetails, error) {
	log.Printf("Creating Person Meeting: %+v", input)
	meetingID := fmt.Sprintf("%+v", rand.Int())

	var users []*model.User
	for _, userID := range input.PeopleIDs {
		user := u.users[*userID]
		if user != nil {
			users = append(users, user)
		}
	}

	meeting := &model.MeetingDetails{
		ID:              meetingID,
		Name:            input.Name,
		StartTime:       input.StartTime,
		DurationMinutes: input.DurationMinutes,
		People:          users,
	}

	u.meetings[meetingID] = meeting

	return meeting, nil
}

func (u *MeetingServiceInmem) CreateProjectMeeting(ctx context.Context, input model.NewProjectMeeting) (*model.MeetingDetails, error) {
	log.Printf("Creating Person Meeting: %+v", input)
	meetingID := fmt.Sprintf("%+v", rand.Int())

	var projects []*model.Project
	for _, projectID := range input.ProjectIDs {
		project := u.projects[*projectID]
		if project != nil {
			projects = append(projects, project)
		}
	}

	meeting := &model.MeetingDetails{
		ID:              meetingID,
		Name:            input.Name,
		StartTime:       input.StartTime,
		DurationMinutes: input.DurationMinutes,
		Projects:        projects,
	}

	u.meetings[meetingID] = meeting

	return meeting, nil
}

func (u *MeetingServiceInmem) GetById(ctx context.Context, meetingID string) (*model.MeetingDetails, error) {
	foundMeeting := u.meetings[meetingID]

	if foundMeeting == nil {
		return &model.MeetingDetails{}, errors.New("Meeting with ID not found")
	}

	return foundMeeting, nil
}

func (meetingService *MeetingServiceInmem) GetByDate(ctx context.Context, datetime time.Time) ([]*model.MeetingDetails, error) {
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

	var meetings []*model.MeetingDetails

	for _, meeting := range meetingService.meetings {

		if meeting.StartTime.After(start) && meeting.StartTime.Before(end) {
			meet := &model.MeetingDetails{
				ID:              meeting.ID,
				Name:            meeting.Name,
				StartTime:       meeting.StartTime,
				DurationMinutes: meeting.DurationMinutes,
				MeetingItems:    meeting.MeetingItems,
			}

			if len(meeting.People) > 0 {
				meet.People = meeting.People
			}
			meetings = append(meetings, meet)
		}
	}

	log.Printf("Found total of %+v meetings", len(meetings))

	return meetings, nil
}
