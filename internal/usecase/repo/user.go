package repo

import (
	"github.com/awgst/datings/internal/entity/model"
	"github.com/awgst/datings/pkg/pagination"
)

type UserFinder interface {
	FindByEmail(email string) (model.User, error)
	FindByID(id int) (model.User, error)
	FindAllProfile(user model.User, paging *pagination.Paginator) ([]model.Profile, error)
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
