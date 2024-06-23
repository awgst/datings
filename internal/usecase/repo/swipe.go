package repo

import "github.com/awgst/datings/internal/entity/model"

type SwipeWriter interface {
	Create(user model.User, swipe *model.Swipe) error
}
