package inmem

// type AuthServiceInmem struct {
// 	UserSevice *UserServiceInmem
// }

// func (authService *AuthServiceInmem) SignIn(ctx context.Context, email string, password string) (*model.User, error) {
// 	var foundUser *model.User
// 	for _, user := range authService.UserSevice.users {
// 		if user.Email == email {
// 			foundUser = user
// 		}
// 	}

// 	if foundUser == nil {
// 		return nil, errors.New("Invalid credentials")
// 	}

// 	if foundUser.Password != password {
// 		return nil, errors.New("Invalid credentials")
// 	}

// 	return foundUser, nil
// }
