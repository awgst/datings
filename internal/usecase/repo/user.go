package repo

import "github.com/awgst/datings/internal/entity/model"

type UserFinder interface {
	FindByEmail(email string) (model.User, error)
	FindByID(id int) (model.User, error)
}

type UserUpdateParams struct {
	ID                int
	UpdateData        map[string]interface{}
	UpdateProfileData map[string]interface{}
}

type UserWriter interface {
	Create(user *model.User) error
	Update(param UserUpdateParams) error
}
