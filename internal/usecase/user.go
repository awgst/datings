package usecase

import (
	"github.com/awgst/datings/internal/customerror"
	"github.com/awgst/datings/internal/entity/model"
	profilerequest "github.com/awgst/datings/internal/entity/request/profile"
	"github.com/awgst/datings/internal/usecase/repo"
)

type UserUsecase interface {
	FindByID(id int) (model.User, error)
	Update(id int, req profilerequest.UpdateProfileRequest) error
}

type userUsecase struct {
	userFinder repo.UserFinder
	userWriter repo.UserWriter
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

func (u userUsecase) Update(id int, req profilerequest.UpdateProfileRequest) error {
	updateData := map[string]interface{}{}
	if req.Email != "" {
		updateData["email"] = req.Email
	}

	updateProfileData := map[string]interface{}{}
	if req.Name != "" {
		updateProfileData["name"] = req.Name
	}

	return u.userWriter.Update(repo.UserUpdateParams{
		ID:                id,
		UpdateData:        updateData,
		UpdateProfileData: updateProfileData,
	})
}
