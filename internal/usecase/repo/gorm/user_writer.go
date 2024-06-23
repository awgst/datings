package gorm

import (
	"database/sql"
	"time"

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
			base: &base{
				db: db,
			},
		},
	}
}

func (u *userWriter) Create(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *userWriter) Update(param repo.UserUpdateParams) error {
	return u.DBTransaction(func(tx *gorm.DB) error {
		if len(param.UpdateData) > 0 {
			param.UpdateData["updated_at"] = time.Now()
			err := tx.Model(&model.User{}).Where("id = ?", param.ID).Updates(param.UpdateData).Error
			if err != nil {
				return err
			}
		}

		if len(param.UpdateProfileData) > 0 {
			var profile model.Profile
			err := tx.Select("id").Where("user_id = ?", param.ID).First(&profile).Error
			if err == gorm.ErrRecordNotFound {
				return tx.Create(&model.Profile{
					UserID: param.ID,
					Name:   param.UpdateProfileData["name"].(string),
					CreatedAt: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				}).Error
			}
			if err != nil {
				return err
			}

			param.UpdateProfileData["updated_at"] = time.Now()
			err = tx.Model(&model.Profile{}).Where("user_id = ?", param.ID).Updates(param.UpdateProfileData).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}
