package inmem

import (
	"context"
	"errors"
	"log"
	"qaplagql/common"
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
	meetingItems    map[string]*model.MeetingItem
	meetingNotes    map[string]*model.MeetingNote
}

func NewMeetingServiceInmem(meetingMap map[string]*model.MeetingDetails, userMap map[string]*model.User, projectMap map[string]*model.Project, projectUserMap map[string][]string, meetingUserMap map[string][]string, meetingProjectMap map[string][]string, tagMap map[string]*model.MeetingNoteTag, meetingItemMap map[string]*model.MeetingItem, meetingNoteMap map[string]*model.MeetingNote) *MeetingServiceInmem {
	return &MeetingServiceInmem{
		meetings:        meetingMap,
		projects:        projectMap,
		projectUser:     projectUserMap,
		users:           userMap,
		meetingUsers:    meetingUserMap,
		meetingProjects: meetingProjectMap,
		tags:            tagMap,
		meetingItems:    meetingItemMap,
		meetingNotes:    meetingNoteMap,
	}
}

func (u *MeetingServiceInmem) CreatePersonMeeting(ctx context.Context, input model.NewUserMeeting) (*model.MeetingDetails, error) {
	log.Printf("Creating Person Meeting: %+v", input)
	meetingID := common.RandomId()

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

	// Create meeting items for each personnel
	var meetingItems []*model.MeetingItem
	for _, pID := range input.PeopleIDs {
		personnel := u.users[*pID]
		attendanceReason := ""
		meetingItem := &model.MeetingItem{
			ID:                      common.RandomId(),
			Personnel:               personnel,
			PlannedAttendanceStatus: "Attending",
			ActualAttendanceStatus:  "Attending",
			AttendanceReason:        &attendanceReason,
		}

		meetingItems = append(meetingItems, meetingItem)
	}

	meeting.MeetingItems = meetingItems

	u.meetings[meetingID] = meeting

	return meeting, nil
}

func (u *MeetingServiceInmem) CreateProjectMeeting(ctx context.Context, input model.NewProjectMeeting) (*model.MeetingDetails, error) {
	log.Printf("Creating Project Meeting: %+v", input)
	meetingID := common.RandomId()

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

	end := start.Add(time.Hour * time.Duration(24)).Add(-1)

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

func (ms *MeetingServiceInmem) UpdatePersonMeeting(ctx context.Context, input model.UpdatedPeopleMeetingDetails) (*model.PeopleMeetingDetails, error) {
	/*
		1. Check to see Meeting with ID exists
		2. Update Name
		3. Update Start Time
		4. Update list of people (if needed)
		5. Update each Meeting Item
	*/
	log.Printf("UpdatePersonMeeting %+v", input)

	foundPeopleMeetingDetails := ms.meetings[input.ID]
	if foundPeopleMeetingDetails == nil {
		return &model.PeopleMeetingDetails{}, errors.New("Meeting with ID not found")
	}

	foundPeopleMeetingDetails.Name = input.Name
	foundPeopleMeetingDetails.StartTime = input.StartTime

	var newlyAssignedPeople []*model.User
	for i := 0; i < len(input.People); i++ {
		person := ms.users[*input.People[i]]
		newlyAssignedPeople = append(newlyAssignedPeople, person)
	}

	foundPeopleMeetingDetails.People = newlyAssignedPeople

	// meeting items need to be created if they don't have an ID
	// else find and update
	for i := 0; i < len(input.MeetingItems); i++ {
		meetingItem := input.MeetingItems[i]
		if meetingItem.ID == "0" {

		} else {

		}
	}

	ms.meetings[input.ID] = foundPeopleMeetingDetails

	// convert to PeopleMeetingDetails
	peopleMeetingDetails := &model.PeopleMeetingDetails{
		ID:   foundPeopleMeetingDetails.ID,
		Name: foundPeopleMeetingDetails.Name,
	}

	return peopleMeetingDetails, nil
}

func (ms *MeetingServiceInmem) UpdateMeetingItem(ctx context.Context, input model.UpdateMeetingItemRequest) (*model.MeetingItem, error) {

	return &model.MeetingItem{}, nil
}
