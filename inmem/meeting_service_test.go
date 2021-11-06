package inmem

// func TestCreatePersonMeeting(t *testing.T) {
// 	// Setup
// 	meetingService := createMeetingService()
// 	ctx := context.Background()
// 	peopleID := "1"
// 	peopleIDs := []*string{&peopleID}
// 	input := model.NewUserMeeting{
// 		Name:            "Meeting in Unit Test",
// 		PeopleIDs:       peopleIDs,
// 		StartTime:       time.Time{},
// 		DurationMinutes: 30,
// 	}

// 	// Run test
// 	meetingDetails, err := meetingService.CreatePersonMeeting(ctx, input)
// 	log.Printf("Created Project Meeting Details %+v", meetingDetails)

// 	// Assert
// 	if err != nil {
// 		t.Log("Expected no error creating person meeting", err)
// 		t.Fail()
// 	}
// 	if meetingDetails == nil {
// 		t.Log("Expected meeting details to be returned", meetingDetails)
// 		t.Fail()
// 	}
// }

// func createMeetingService() *MeetingServiceInmem {
// 	usersMap := make(map[string]*model.User)
// 	usersMap["1"] = &model.User{
// 		ID:         "1",
// 		FirstName:  "Dale",
// 		LastName:   "Chang",
// 		GoesBy:     "Dale",
// 		MiddleName: "",
// 		Email:      "dale@eloquia.io",
// 		Password:   "notagoodpassword",
// 		Ethnicity:  "East Asian",
// 		Position:   "Software Engineer",
// 	}
// 	projectMap := make(map[string]*model.Project)
// 	projectMap["1"] = &model.Project{
// 		ID:          "1",
// 		Name:        "Test Project",
// 		Description: "Test project for testing purposes",
// 		Slug:        "test-project",
// 	}
// 	meetingMap := make(map[string]*model.MeetingDetails)
// 	tagMap := make(map[string]*model.MeetingNoteTag)

// 	projectUserMap := make(map[string][]string)
// 	meetingUserMap := make(map[string][]string)
// 	meetingProjectMap := make(map[string][]string)
// 	meetingItemMap := make(map[string]*model.MeetingItem)
// 	meetingNoteMap := make(map[string]*model.MeetingNote)

// 	meetingServiceInmem := NewMeetingServiceInmem(
// 		meetingMap,
// 		usersMap,
// 		projectMap,
// 		projectUserMap,
// 		meetingUserMap,
// 		meetingProjectMap,
// 		tagMap,
// 		meetingItemMap,
// 		meetingNoteMap,
// 	)
// 	return meetingServiceInmem
// }
