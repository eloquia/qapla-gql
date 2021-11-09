package seequell

import "qaplagql/graph/model"

func UserToDomain(user CoreQaplaUser) *model.UserDetails {
	return &model.UserDetails{
		ID:        int(user.UserID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
