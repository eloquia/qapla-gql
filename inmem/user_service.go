package inmem

// type UserServiceInmem struct {
// 	projects    map[string]*model.Project
// 	projectUser map[string][]string
// 	users       map[string]*model.UserDetails
// }

// func NewUserServiceInmem(userMap map[string]*model.UserDetails, projectMap map[string]*model.Project, projectUserMap map[string][]string) *UserServiceInmem {
// 	return &UserServiceInmem{
// 		users:       userMap,
// 		projects:    projectMap,
// 		projectUser: projectUserMap,
// 	}
// }

// func (userService *UserServiceInmem) CreateUser(ctx context.Context, input *model.NewUser) (*model.UserDetails, error) {
// 	// check to see if user with email exists
// 	isEmailUsed := false
// 	for _, user := range userService.users {
// 		if user.Email == input.Email {
// 			isEmailUsed = true
// 		}
// 	}
// 	if isEmailUsed {
// 		return nil, errors.New("Email address already in use")
// 	}

// 	// create user
// 	user := &model.UserDetails{
// 		ID:    common.RandomId(),
// 		Email: input.Email,
// 	}

// 	userService.users[user.ID] = user
// 	return user, nil
// }

// func (userService *UserServiceInmem) GetById(id string) (*model.UserDetails, error) {
// 	var foundUser *model.UserDetails
// 	for userId, user := range userService.users {
// 		if userId == id {
// 			foundUser = user
// 		}
// 	}

// 	if foundUser == nil {
// 		return &model.UserDetails{}, errors.New("User with ID does not exist")
// 	}

// 	return foundUser, nil
// }

// func (userService *UserServiceInmem) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
// 	_, err := userService.GetById(input.ID)
// 	if err != nil {
// 		return &model.User{}, err
// 	}

// 	user := &model.User{
// 		ID:        input.ID,
// 		FirstName: input.FirstName,
// 		LastName:  input.LastName,
// 	}

// 	userService.users[user.ID] = user

// 	return user, nil
// }

// func (userService *UserServiceInmem) GetAll(ctx context.Context) ([]*model.User, error) {
// 	var users []*model.User
// 	for _, user := range userService.users {
// 		users = append(users, user)
// 	}

// 	return users, nil
// }

// func (userService *UserServiceInmem) GetAllUserDetails(ctx context.Context) ([]*model.UserDetails, error) {
// 	var userDetails []*model.UserDetails
// 	for _, user := range userService.users {
// 		userDetail := &model.UserDetails{
// 			ID:        user.ID,
// 			FirstName: user.FirstName,
// 			LastName:  user.LastName,
// 			Email:     user.Email,
// 		}

// 		// fetch projects associated with User
// 		var assignedProjectIDs []string
// 		for projectID, assignedUserIDs := range userService.projectUser {
// 			// figure out if user is assigned to the project
// 			for _, assignedUserID := range assignedUserIDs {
// 				if assignedUserID == userDetail.ID {
// 					assignedProjectIDs = append(assignedProjectIDs, projectID)
// 				}
// 			}
// 		}

// 		var assignedProjects []*model.Project
// 		for _, pID := range assignedProjectIDs {
// 			project := userService.projects[pID]

// 			assignedProjects = append(assignedProjects, project)
// 		}

// 		userDetail.AssignedProjects = assignedProjects

// 		userDetails = append(userDetails, userDetail)
// 	}

// 	return userDetails, nil
// }

// func (us *UserServiceInmem) GetAllShortUserDetails(ctx context.Context) ([]*model.UserDetailsShort, error) {
// 	var shorts []*model.UserDetailsShort

// 	return shorts, nil
// }

// func (u *UserServiceInmem) AddPersonnel(ctx context.Context, input model.NewPersonnel) (*model.UserDetails, error) {
// 	// Make sure that email isn't already used
// 	email := input.Email
// 	isEmailExists := false
// 	for _, user := range u.users {
// 		if user.Email == email {
// 			isEmailExists = true
// 		}
// 	}

// 	if isEmailExists {
// 		return &model.UserDetails{}, errors.New("Email already in use")
// 	}

// 	// check for default values
// 	middleName := ""
// 	if input.MiddleName != &middleName {
// 		middleName = *input.MiddleName
// 	}

// 	goesBy := ""
// 	if input.GoesBy != &goesBy {
// 		goesBy = *input.GoesBy
// 	}

// 	gender := ""
// 	if input.Gender != &gender {
// 		gender = *input.Gender
// 	}

// 	ethnicity := ""
// 	if input.Ethnicity != &ethnicity {
// 		ethnicity = *input.Ethnicity
// 	}

// 	position := ""
// 	if input.Position != &position {
// 		position = *input.Position
// 	}

// 	institution := ""
// 	if input.Institution != &institution {
// 		institution = *input.Institution
// 	}

// 	userId := common.RandomId()
// 	user := &model.User{
// 		ID:          userId,
// 		FirstName:   input.FirstName,
// 		LastName:    input.LastName,
// 		MiddleName:  middleName,
// 		GoesBy:      goesBy,
// 		Email:       input.Email,
// 		Gender:      gender,
// 		Ethnicity:   ethnicity,
// 		Position:    position,
// 		Institution: institution,
// 	}

// 	u.users[userId] = user

// 	userDetails := &model.UserDetails{
// 		ID:        user.ID,
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 		Email:     user.Email,
// 	}

// 	return userDetails, nil
// }
