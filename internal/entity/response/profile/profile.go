package profile

import "github.com/awgst/datings/internal/entity/model"

type ProfileResponse struct {
	ID    int     `json:"id"`
	Email string  `json:"email"`
	Name  *string `json:"name"`
}

func (ProfileResponse) Make(user model.User) ProfileResponse {
	var name *string
	if user.Profile != nil {
		name = &user.Profile.Name
	}

	return ProfileResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  name,
	}
}
