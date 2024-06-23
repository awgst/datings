package integrationtest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_api_integration_test_get_profile_endpoint(t *testing.T) {
	app, ts := runTestServer()
	db := app.DB.Gorm
	defer ts.Close()
	defer TruncateDatabase(db)

	t.Run("it should return 200 when get profile success", func(t *testing.T) {
		defer TruncateDatabase(db)
		accessToken := mockAccessToken(t, app, true)
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("%s/v1/profile", ts.URL),
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
			fmt.Sprintf("%s/v1/profile", ts.URL),
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
