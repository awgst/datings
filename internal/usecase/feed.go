package usecase

import (
	"github.com/awgst/datings/internal/entity/model"
	feedrequest "github.com/awgst/datings/internal/entity/request/feed"
	"github.com/awgst/datings/internal/usecase/repo"
	"github.com/awgst/datings/pkg/pagination"
)

type FeedUsecase interface {
	Swipe(user model.User, req feedrequest.SwipeRequest) error
	Recommendation(user model.User, paging *pagination.Paginator) ([]model.Profile, error)
}

type feedUsecase struct {
	swipeWriter repo.SwipeWriter
	userFinder  repo.UserFinder
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

func (u feedUsecase) Recommendation(user model.User, paging *pagination.Paginator) ([]model.Profile, error) {
	return u.userFinder.FindAllProfile(user, paging)
}
