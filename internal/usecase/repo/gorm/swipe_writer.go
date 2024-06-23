package gorm

import (
	"fmt"
	"time"

	"github.com/awgst/datings/internal/customerror"
	"github.com/awgst/datings/internal/entity/model"
	"github.com/awgst/datings/internal/usecase/repo"
	"gorm.io/gorm"
)

type swipeRepo struct {
	*base
}

type swipeWriter struct {
	*swipeRepo
}

func NewGormSwipeWriter(db *gorm.DB) repo.SwipeWriter {
	return &swipeWriter{
		swipeRepo: &swipeRepo{
			base: &base{
				db: db,
			},
		},
	}
}

func (s *swipeWriter) Create(user model.User, swipe *model.Swipe) error {
	return s.DBTransaction(func(tx *gorm.DB) error {
		if !user.HasUnlimitedSwipe() {
			var count int64
			err := tx.Table(swipe.TableName()).
				Where(
					"user_id = ? AND DATE(created_at) = DATE(?)",
					swipe.UserID,
					time.Now(),
				).
				Count(&count).
				Error
			if err != nil {
				return err
			}
			if count >= 10 {
				return customerror.Error{
					Code: customerror.ErrorCodeInvalidRequest,
					Err:  fmt.Sprintf("user with id %d has already swiped on 10 profiles", swipe.UserID),
				}
			}
		}

		var count int64
		err := tx.Table(swipe.TableName()).
			Where(
				"user_id = ? AND profile_id = ? AND DATE(created_at) = DATE(?)",
				swipe.UserID,
				swipe.ProfileID,
				time.Now(),
			).
			Count(&count).
			Error
		if err != nil {
			return err
		}
		if count > 0 {
			return customerror.Error{
				Code: customerror.ErrorCodeInvalidRequest,
				Err:  fmt.Sprintf("user with id %d has already swiped on profile with id %d for today", swipe.UserID, swipe.ProfileID),
			}
		}

		var profile model.Profile
		err = tx.Select("id", "user_id").Where("id = ?", swipe.ProfileID).First(&profile).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if profile.UserID == swipe.UserID {
			return customerror.Error{
				Code: customerror.ErrorCodeInvalidRequest,
				Err:  fmt.Sprintf("user with id %d cannot swipe on itself", swipe.UserID),
			}
		}

		return tx.Create(swipe).Error
	})
}
