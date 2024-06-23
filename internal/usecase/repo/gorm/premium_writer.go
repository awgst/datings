package gorm

import (
	"fmt"

	"github.com/awgst/datings/internal/customerror"
	"github.com/awgst/datings/internal/entity/model"
	"github.com/awgst/datings/internal/usecase/repo"
	"gorm.io/gorm"
)

type premiumRepo struct {
	*base
}

type premiumWriter struct {
	*premiumRepo
}

func NewGormPremiumWriter(db *gorm.DB) repo.PremiumWriter {
	return &premiumWriter{
		premiumRepo: &premiumRepo{
			base: &base{
				db: db,
			},
		},
	}
}

func (u *premiumWriter) Create(createPremium *model.Premium) error {
	return u.DBTransaction(func(tx *gorm.DB) error {
		var premium model.Premium
		err := tx.Where("user_id = ? AND feature = ?", createPremium.UserID, createPremium.Feature).First(&premium).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if premium.ID != 0 {
			return customerror.Error{
				Code: customerror.ErrorCodeAlreadyExists,
				Err:  fmt.Sprintf("premium with user_id %d and feature %s already exists", premium.UserID, premium.Feature),
			}
		}

		return tx.Create(createPremium).Error
	})
}
