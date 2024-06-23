package repo

import "github.com/awgst/datings/internal/entity/model"

type UserFinder interface {
	FindByEmail(email string) (model.User, error)
	FindByID(id int) (model.User, error)
}

type UserWriter interface {
	Create(user *model.User) error
}
