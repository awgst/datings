package usecase

import "github.com/awgst/datings/pkg/app"

type Usecase struct {
	App *app.App
}

func New(app *app.App) *Usecase {
	return &Usecase{
		App: app,
	}
}
