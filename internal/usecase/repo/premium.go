package repo

import "github.com/awgst/datings/internal/entity/model"

type PremiumWriter interface {
	Create(createPremium *model.Premium) error
}
