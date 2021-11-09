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

// func (us *UserServiceInmem) GetAllShortUserDetails(ctx context.Context) ([]*model.UserDetailsShort, error) {
// 	var shorts []*model.UserDetailsShort

// 	return shorts, nil
// }
