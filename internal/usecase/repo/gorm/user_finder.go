package gorm

import (
	"github.com/awgst/datings/internal/entity/model"
	"github.com/awgst/datings/internal/usecase/repo"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

type userFinder struct {
	*userRepo
}

func NewGormUserFinder(db *gorm.DB) repo.UserFinder {
	return &userFinder{
		userRepo: &userRepo{
			db: db,
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
