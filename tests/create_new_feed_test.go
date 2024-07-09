package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateFeed(t *testing.T) {
	httpServer := httptest.NewServer(srv)
	defer httpServer.Close()

	user := createUser(t)

	t.Run("Should return 401 if request is not authenticated", func(t *testing.T) {
		r, err := NewRequest(httpServer.URL + "/feeds").
			WithMethod(http.MethodPost).
			WithBody(`{"name":"Jon","url":"http://example.com"}`).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		if r.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected status code to be %v; got %v", http.StatusUnauthorized, r.StatusCode)
		}
	})

	t.Run("Should return 400 if request body is empty", func(t *testing.T) {
		r, err := NewRequest(httpServer.URL+"/feeds").
			WithMethod(http.MethodPost).
			WithHeader("X-Api-Key", user.ApiKey).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		if r.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code to be %v; got %v", http.StatusBadRequest, r.StatusCode)
		}
	})

	t.Run("Should return 400 if name is empty", func(t *testing.T) {
		r, err := NewRequest(httpServer.URL+"/feeds").
			WithMethod(http.MethodPost).
			WithHeader("X-Api-Key", user.ApiKey).
			WithBody(`{"name":"","url":"http://example.com"}`).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		if r.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code to be %v; got %v", http.StatusBadRequest, r.StatusCode)
		}
	})

	t.Run("Should return 400 if url is empty", func(t *testing.T) {
		r, err := NewRequest(httpServer.URL+"/feeds").
			WithMethod(http.MethodPost).
			WithHeader("X-Api-Key", user.ApiKey).
			WithBody(`{"name":"Jon","url":""}`).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		if r.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code to be %v; got %v", http.StatusBadRequest, r.StatusCode)
		}
	})

	t.Run("Should return 201 if request is valid", func(t *testing.T) {
		r, err := NewRequest(httpServer.URL+"/feeds").
			WithMethod(http.MethodPost).
			WithHeader("X-Api-Key", user.ApiKey).
			WithBody(`{"name":"Jon","url":"http://example.com"}`).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		if r.StatusCode != http.StatusCreated {
			t.Errorf("expected status code to be %v; got %v", http.StatusCreated, r.StatusCode)
		}
	})
}
