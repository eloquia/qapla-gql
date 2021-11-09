package seequell

import "qaplagql/graph/model"

func UserToShortDomain(user CoreQaplaUser) *model.UserDetailsShort {
	return &model.UserDetailsShort{
		ID:        int(user.UserID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func UserToDetailedDomain(user GetUserDetailsRow) *model.UserDetails {
	return &model.UserDetails{
		ID:        int(user.UserID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func ToUserDetailsDomain(userDetails CoreQaplaUserDetail) *model.UserDetails {
	return &model.UserDetails{}
}

func ToUserDomain(userModel CoreQaplaUser) *model.User {
	return &model.User{
		ID:        int(userModel.UserID),
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}
