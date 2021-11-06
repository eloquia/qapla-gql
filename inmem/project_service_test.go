package inmem

// func TestAddUserToProject(t *testing.T) {
// 	// Setup
// 	userMap := make(map[string]*model.User)
// 	userMap["1"] = &model.User{
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
// 	projectUserMap := make(map[string][]string)
// 	ctx := context.Background()
// 	userService := &UserServiceInmem{
// 		users:       userMap,
// 		projects:    projectMap,
// 		projectUser: projectUserMap,
// 	}
// 	projectService := &ProjectServiceInmem{
// 		users:       userMap,
// 		projects:    projectMap,
// 		projectUser: projectUserMap,
// 	}

// 	users, err := userService.GetAll()
// 	if err != nil {
// 		t.Log("Expected no error getting users", err)
// 		t.Fail()
// 	}
// 	if len(users) < 1 {
// 		t.Log("Expected at least one user to exist")
// 		t.Fail()
// 	}

// 	projects, err := projectService.GetAll(ctx)
// 	if err != nil {
// 		t.Log("Expected no error getting projects", err)
// 		t.Fail()
// 	}
// 	if len(projects) < 1 {
// 		t.Log("Expected at least one project to exist")
// 		t.Fail()
// 	}

// 	lastUser := users[len(users)-1]
// 	lastProject := projects[len(projects)-1]

// 	// Test
// 	_, err = projectService.AddUserToProject(ctx, lastUser.ID, lastProject.ID)
// 	if err != nil {
// 		t.Log("Expected no error assigning user to project", err)
// 		t.Fail()
// 	}
// }

// func TestGetProjectPersonnel(t *testing.T) {
// 	// Setup
// 	userMap := make(map[string]*model.User)
// 	userMap["1"] = &model.User{
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
// 	projectUserMap := make(map[string][]string)
// 	ctx := context.Background()
// 	userService := &UserServiceInmem{
// 		users:       userMap,
// 		projects:    projectMap,
// 		projectUser: projectUserMap,
// 	}
// 	projectService := &ProjectServiceInmem{
// 		users:       userMap,
// 		projects:    projectMap,
// 		projectUser: projectUserMap,
// 	}

// 	users, err := userService.GetAll()
// 	if err != nil {
// 		t.Log("Expected no error getting users", err)
// 		t.Fail()
// 	}
// 	if len(users) < 1 {
// 		t.Log("Expected at least one user to exist")
// 		t.Fail()
// 	}

// 	projects, err := projectService.GetAll(ctx)
// 	if err != nil {
// 		t.Log("Expected no error getting projects", err)
// 		t.Fail()
// 	}
// 	if len(projects) < 1 {
// 		t.Log("Expected at least one project to exist")
// 		t.Fail()
// 	}

// 	lastUser := users[len(users)-1]
// 	lastProject := projects[len(projects)-1]

// 	_, err = projectService.AddUserToProject(ctx, lastUser.ID, lastProject.ID)
// 	if err != nil {
// 		t.Log("Expected no error assigning user to project", err)
// 		t.Fail()
// 	}

// 	// Test
// 	usersInProject, err := projectService.GetProjectPersonnel(ctx, lastProject.ID)
// 	if err != nil {
// 		t.Log("Expected no error while getting personnel assigned to a project", err)
// 		t.Fail()
// 	}

// 	if len(usersInProject) == 0 {
// 		t.Log("Expected at least one user to be assigned to project")
// 		t.Fail()
// 	}
// }
