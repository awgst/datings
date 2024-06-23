package integrationtest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/awgst/datings/internal/entity/model"
	"github.com/stretchr/testify/assert"
)

func Test_api_integration_test_feed_recommendation_endpoint(t *testing.T) {
	app, ts := runTestServer()
	db := app.DB.Gorm
	defer ts.Close()
	defer TruncateDatabase(db)

	t.Run("it should return 200 when get feed recommendation success", func(t *testing.T) {
		defer TruncateDatabase(db)
		if err := db.Create(&model.User{
			ID:           2,
			Email:        "XjK9A@example.com",
			PasswordHash: "password_hash",
			Profile: &model.Profile{
				ID:     2,
				UserID: 2,
			},
		}).Error; err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		accessToken := mockAccessToken(t, app, true)
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("%s/v1/feed/", ts.URL),
			bytes.NewBufferString(""),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Contains(t, string(body), "id")
	})

	t.Run("it should return 401 when get profile without Authorization header", func(t *testing.T) {
		defer TruncateDatabase(db)
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("%s/v1/feed", ts.URL),
			bytes.NewBufferString(""),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

}
