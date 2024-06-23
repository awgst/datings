package integrationtest

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/awgst/datings/internal/entity/model"
	"github.com/stretchr/testify/assert"
)

func Test_api_integration_test_feed_swipe_endpoint(t *testing.T) {
	app, ts := runTestServer()
	db := app.DB.Gorm
	defer ts.Close()
	defer TruncateDatabase(db)

	t.Run("it should return 200 when feed swipe success", func(t *testing.T) {
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
			"POST",
			fmt.Sprintf("%s/v1/feed/swipe", ts.URL),
			bytes.NewBufferString(`{
				"profile_id": 2,
				"type": "like"
			}`),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("it should return 200 when premium user swipe more than 10 times", func(t *testing.T) {
		defer TruncateDatabase(db)
		users := []*model.User{}
		for i := 2; i < 13; i++ {
			users = append(users, &model.User{
				ID:           i,
				Email:        fmt.Sprintf("test%d@test.com", i),
				PasswordHash: "password_hash",
				Profile: &model.Profile{
					ID:     i,
					UserID: i,
				},
			})
		}
		if err := db.Create(&users).Error; err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		swipes := []*model.Swipe{}
		for _, u := range users[:10] {
			swipes = append(swipes, &model.Swipe{
				UserID:    mockUser.ID,
				ProfileID: u.Profile.ID,
				Type:      model.SwipeTypeLike,
				CreatedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
			})
		}
		if err := db.Create(&swipes).Error; err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		accessToken := mockAccessToken(t, app, true)
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v1/feed/swipe", ts.URL),
			bytes.NewBufferString(`{
				"profile_id": 12,
				"type": "like"
			}`),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("it should return 400 when user swipe profile that already swiped", func(t *testing.T) {
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

		if err := db.Create(&model.Swipe{
			ProfileID: 2,
			UserID:    1,
			Type:      model.SwipeTypeLike,
			CreatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}).Error; err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		accessToken := mockAccessToken(t, app, true)
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v1/feed/swipe", ts.URL),
			bytes.NewBufferString(`{
				"profile_id": 2,
				"type": "pass"
			}`),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("it should return 400 when non premium user swipe more than 10 times", func(t *testing.T) {
		defer TruncateDatabase(db)
		users := []*model.User{}
		for i := 2; i < 12; i++ {
			users = append(users, &model.User{
				ID:           i,
				Email:        fmt.Sprintf("test%d@test.com", i),
				PasswordHash: "password_hash",
				Profile: &model.Profile{
					ID:     i,
					UserID: i,
				},
			})
		}
		if err := db.Create(&users).Error; err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		swipes := []*model.Swipe{}
		for _, u := range users {
			swipes = append(swipes, &model.Swipe{
				UserID:    mockUser.ID,
				ProfileID: u.Profile.ID,
				Type:      model.SwipeTypeLike,
				CreatedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
			})
		}
		if err := db.Create(&swipes).Error; err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		accessToken := mockAccessToken(t, app, false)
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v1/feed/swipe", ts.URL),
			bytes.NewBufferString(`{
				"profile_id": 11,
				"type": "like"
			}`),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
