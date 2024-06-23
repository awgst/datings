package integrationtest

import (
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/awgst/datings/config"
	v1 "github.com/awgst/datings/internal/controller/http/v1"
	"github.com/awgst/datings/internal/entity/model"
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/app"
	"github.com/awgst/datings/pkg/database"
	"github.com/awgst/datings/pkg/logger"
	"github.com/awgst/datings/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var mockUser = &model.User{
	ID:           1,
	Email:        "XjK9A@example.com",
	PasswordHash: "password",
	Profile: &model.Profile{
		UserID: 1,
		Name:   "name",
	},
	Premium: &model.Premium{
		UserID:  1,
		Feature: model.PremiumFeatureNoSwipeQuota,
	},
}

func setupServer() (*app.App, *gin.Engine) {
	cfg, err := config.NewConfigForTest()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	errorLogger := logger.New()

	// Database
	db := database.New(database.ConnectionConfig{
		DbUrl: cfg.Database.URL,
	})

	// App
	app := &app.App{
		Config: cfg,
		DB:     db,
		Logger: errorLogger,
	}

	// Usecase
	uc := usecase.New(app)

	// HTTP Server
	gin.SetMode(gin.TestMode)
	handler := gin.New()
	v1.NewRouter(handler, uc)

	return app, handler
}

func runTestServer() (*app.App, *httptest.Server) {
	app, server := setupServer()
	TruncateDatabase(app.DB.Gorm)
	return app, httptest.NewServer(server)
}

func mockAccessToken(t *testing.T, app *app.App, isPremium bool) string {
	if !isPremium {
		mockUser.Premium = nil
	}
	if err := app.DB.Gorm.Create(mockUser).Error; err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	premium := "none"
	if mockUser.Premium != nil {
		premium = string(mockUser.Premium.Feature)
	}
	accessToken, err := token.NewToken().JwtToken(app.Config.JWT.Secret, jwt.MapClaims{
		"user_id": mockUser.ID,
		"email":   mockUser.Email,
		"premium": premium,
		"exp":     time.Now().Add(time.Minute * time.Duration(app.Config.JWT.ExpireInMinutes)).Unix(),
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	return accessToken
}
