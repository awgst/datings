package integrationtest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/awgst/datings/pkg/password"
	"github.com/stretchr/testify/assert"
)

func Test_api_integration_test_login_endpoint(t *testing.T) {
	app, ts := runTestServer()
	db := app.DB.Gorm
	defer ts.Close()
	defer TruncateDatabase(db)

	t.Run("it should return 200 when login success", func(t *testing.T) {
		pass := password.NewPassword().Hash("password")
		if err := db.Exec("INSERT INTO users (email, password_hash) VALUES ('XjK9A@example.com', ?)", pass).Error; err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		resp, err := http.Post(
			fmt.Sprintf("%s/v1/auth/login", ts.URL),
			"application/json",
			bytes.NewBufferString(`{"email": "XjK9A@example.com", "password": "password"}`),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Contains(t, string(body), "token")
	})

	t.Run("it should return 400 when login failed", func(t *testing.T) {
		resp, err := http.Post(
			fmt.Sprintf("%s/v1/auth/login", ts.URL),
			"application/json",
			bytes.NewBufferString(`{"email": "unknown@mail.com", "password": "password"}`),
		)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("it should return 422 when request miss required parameters", func(t *testing.T) {
		resp, err := http.Post(fmt.Sprintf("%s/v1/auth/login", ts.URL), "application/json", bytes.NewBufferString("{}"))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	})

}
