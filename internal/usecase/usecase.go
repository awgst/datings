package usecase

import (
	"github.com/awgst/datings/internal/usecase/repo/gorm"
	"github.com/awgst/datings/pkg/app"
	"github.com/awgst/datings/pkg/password"
	"github.com/awgst/datings/pkg/token"
)

type Usecase struct {
	App     *app.App
	Auth    AuthUsecase
	User    UserUsecase
	Premium PremiumUsecase
	Feed    FeedUsecase
}

func New(app *app.App) *Usecase {
	gormUserFinder := gorm.NewGormUserFinder(app.DB.Gorm)
	gormUserWriter := gorm.NewGormUserWriter(app.DB.Gorm)
	gormPremiumWriter := gorm.NewGormPremiumWriter(app.DB.Gorm)
	gormSwipeWriter := gorm.NewGormSwipeWriter(app.DB.Gorm)

	return &Usecase{
		App: app,
		Auth: NewAuthUsecase(authUsecase{
			cfg:        app.Config,
			userFinder: gormUserFinder,
			userWriter: gormUserWriter,
			token:      token.NewToken(),
			password:   password.NewPassword(),
		}),
		User: NewUserUsecase(userUsecase{
			userFinder: gormUserFinder,
			userWriter: gormUserWriter,
		}),
		Premium: NewPremiumUsecase(premiumUsecase{
			premiumWriter: gormPremiumWriter,
		}),
		Feed: NewFeedUsecase(feedUsecase{
			swipeWriter: gormSwipeWriter,
			userFinder:  gormUserFinder,
		}),
	}
}
