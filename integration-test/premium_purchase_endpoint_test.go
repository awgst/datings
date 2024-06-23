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

func Test_api_integration_test_purchase_premium_endpoint(t *testing.T) {
	app, ts := runTestServer()
	db := app.DB.Gorm
	defer ts.Close()

	t.Run("it should return 200 when purchase premium no swipe quota success", func(t *testing.T) {
		defer TruncateDatabase(db)
		accessToken := mockAccessToken(t, app, true)
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v1/premium", ts.URL),
			bytes.NewBufferString(`{
				"premium_feature": "no_swipe_quota"
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

	t.Run("it should return 400 when purchase premium already purchased", func(t *testing.T) {
		defer TruncateDatabase(db)
		if err := db.Create(&model.Premium{
			UserID:  mockUser.ID,
			Feature: model.PremiumFeatureNoSwipeQuota,
		}).Error; err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		accessToken := mockAccessToken(t, app, true)
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v1/premium", ts.URL),
			bytes.NewBufferString(`{"premium_feature": "no_swipe_quota"}`),
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

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Contains(t, string(body), fmt.Sprintf("premium with user_id %d and feature no_swipe_quota already exists", mockUser.ID))
	})

	t.Run("it should return 422 when request miss required parameters", func(t *testing.T) {
		defer TruncateDatabase(db)
		accessToken := mockAccessToken(t, app, true)
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v1/premium", ts.URL),
			bytes.NewBufferString(`{}`),
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

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
		assert.Contains(t, string(body), "premium_feature")
	})
}
