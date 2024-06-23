package usecase

import (
	"github.com/awgst/datings/internal/customerror"
	"github.com/awgst/datings/internal/entity/model"
	"github.com/awgst/datings/internal/usecase/repo"
)

type UserUsecase interface {
	FindByID(id int) (model.User, error)
}

type userUsecase struct {
	userFinder repo.UserFinder
}

func NewUserUsecase(uc userUsecase) UserUsecase {
	return uc
}

func (u userUsecase) FindByID(id int) (model.User, error) {
	user, err := u.userFinder.FindByID(id)
	if err != nil {
		return model.User{}, err
	}

	if user.ID == 0 {
		return model.User{}, customerror.Error{
			Code: customerror.ErrorCodeNotFound,
			Err:  nil,
		}
	}

	return user, nil
}
