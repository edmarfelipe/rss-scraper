package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {
	httpServer := httptest.NewServer(srv)
	defer httpServer.Close()

	t.Run("Should return 400 if request body is empty", func(t *testing.T) {
		resp, err := NewRequest(httpServer.URL + "/users").
			WithBody(`{}`).
			WithMethod(http.MethodPost).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code to be %v; got %v", http.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("Should return 400 if name is empty", func(t *testing.T) {
		resp, err := NewRequest(httpServer.URL + "/users").
			WithBody(`{"name":""}`).
			WithMethod(http.MethodPost).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code to be %v; got %v", http.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("Should create a new user", func(t *testing.T) {
		resp, err := NewRequest(httpServer.URL + "/users").
			WithBody(`{"name":"Jon"}`).
			WithMethod(http.MethodPost).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status code to be %v; got %v", http.StatusCreated, resp.StatusCode)
		}

		type response struct {
			ID        uuid.UUID `json:"id"`
			Name      string    `json:"name"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		}

		body, err := decode[response](resp)
		if err != nil {
			t.Fatalf("error decoding response body: %v", err)
		}

		if body.Name != "Jon" {
			t.Errorf("expected name to be %q; got %q", "Jon", body.Name)
		}

		if body.ID == uuid.Nil {
			t.Errorf("expected ID to be a valid UUID; got %v", body.ID)
		}

		if body.CreatedAt.IsZero() {
			t.Errorf("expected CreatedAt to be a valid time; got %v", body.CreatedAt)
		}

		if body.UpdatedAt.IsZero() {
			t.Errorf("expected UpdatedAt to be a valid time; got %v", body.UpdatedAt)
		}
	})
}
