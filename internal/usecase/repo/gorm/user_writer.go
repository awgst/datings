package gorm

import (
	"github.com/awgst/datings/internal/entity/model"
	"github.com/awgst/datings/internal/usecase/repo"
	"gorm.io/gorm"
)

type userWriter struct {
	*userRepo
}

func NewGormUserWriter(db *gorm.DB) repo.UserWriter {
	return &userWriter{
		userRepo: &userRepo{
			db: db,
		},
	}
}

func (u *userWriter) Create(user *model.User) error {
	return u.db.Create(user).Error
}
