package gorm

import (
	"github.com/awgst/datings/internal/entity/model"
	"github.com/awgst/datings/internal/usecase/repo"
	"gorm.io/gorm"
)

type userRepo struct {
	*base
}

type userFinder struct {
	*userRepo
}

func NewGormUserFinder(db *gorm.DB) repo.UserFinder {
	return &userFinder{
		userRepo: &userRepo{
			base: &base{
				db: db,
			},
		},
	}
}

func (u *userFinder) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return model.User{}, nil
	}
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userFinder) FindByID(id int) (model.User, error) {
	var user model.User
	err := u.db.Where("id = ?", id).Preload("Profile").First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return model.User{}, nil
	}
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
