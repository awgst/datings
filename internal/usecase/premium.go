package usecase

import (
	"github.com/awgst/datings/internal/entity/model"
	premiumrequest "github.com/awgst/datings/internal/entity/request/premium"
	"github.com/awgst/datings/internal/usecase/repo"
)

type PremiumUsecase interface {
	Create(userID int, req premiumrequest.CreatePremiumRequest) error
}

type premiumUsecase struct {
	premiumWriter repo.PremiumWriter
}

func NewPremiumUsecase(uc premiumUsecase) PremiumUsecase {
	return uc
}

func (u premiumUsecase) Create(userID int, req premiumrequest.CreatePremiumRequest) error {
	return u.premiumWriter.Create(&model.Premium{
		UserID:  userID,
		Feature: model.PremiumFeature(req.PremiumFeature),
	})
}
