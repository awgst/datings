package usecase

import (
	"github.com/awgst/datings/internal/entity/model"
	feedrequest "github.com/awgst/datings/internal/entity/request/feed"
	"github.com/awgst/datings/internal/usecase/repo"
)

type FeedUsecase interface {
	Swipe(user model.User, req feedrequest.SwipeRequest) error
}

type feedUsecase struct {
	swipeWriter repo.SwipeWriter
}

func NewFeedUsecase(uc feedUsecase) FeedUsecase {
	return uc
}

func (u feedUsecase) Swipe(user model.User, req feedrequest.SwipeRequest) error {
	return u.swipeWriter.Create(user, &model.Swipe{
		UserID:    user.ID,
		ProfileID: req.ProfileID,
		Type:      model.SwipeType(req.Type),
	})
}
