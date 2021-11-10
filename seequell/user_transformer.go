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

func UserDetailsToDomain(userDetailModel CoreQaplaUserDetail, userModel CoreQaplaUser) *model.UserDetails {
	deets := &model.UserDetails{
		ID:          int(userDetailModel.UserID),
		MiddleName:  &userDetailModel.MiddleName.String,
		GoesBy:      &userDetailModel.GoesBy.String,
		Gender:      &userDetailModel.Gender.String,
		Ethnicity:   &userDetailModel.Ethnicity.String,
		Position:    &userDetailModel.Position.String,
		Institution: &userDetailModel.Institution.String,
	}

	if userModel.FirstName != "" {
		deets.FirstName = userModel.FirstName
	}
	if userModel.LastName != "" {
		deets.LastName = userModel.LastName
	}
	if userModel.Email != "" {
		deets.Email = userModel.Email
	}

	return deets
}
