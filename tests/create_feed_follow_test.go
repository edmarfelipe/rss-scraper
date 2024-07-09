package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateFeedFollow(t *testing.T) {
	httpServer := httptest.NewServer(srv)
	defer httpServer.Close()

	user := createUser(t)
	feed := createFeed(t)

	t.Run("Should return 401 if request is not authenticated", func(t *testing.T) {
		r, err := NewRequest(httpServer.URL + "/feeds/follow").
			WithMethod(http.MethodPost).
			WithBody(`{"feed_id":""}`).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		if r.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected status code to be %v; got %v", http.StatusUnauthorized, r.StatusCode)
		}
	})

	t.Run("Should return 400 if request body is empty", func(t *testing.T) {
		r, err := NewRequest(httpServer.URL+"/feeds/follow").
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

	t.Run("Should return 400 if feed_id is empty", func(t *testing.T) {
		r, err := NewRequest(httpServer.URL+"/feeds/follow").
			WithMethod(http.MethodPost).
			WithHeader("X-Api-Key", user.ApiKey).
			WithBody(`{"feed_id":""}`).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		if r.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code to be %v; got %v", http.StatusBadRequest, r.StatusCode)
		}
	})

	t.Run("Should return 200 if request is valid", func(t *testing.T) {
		r, err := NewRequest(httpServer.URL+"/feeds/follow").
			WithMethod(http.MethodPost).
			WithHeader("X-Api-Key", user.ApiKey).
			WithBody(fmt.Sprintf(`{"feed_id":"%s"}`, feed.ID)).
			Send()
		if err != nil {
			t.Fatalf("error making request to server: %v", err)
		}
		if r.StatusCode != http.StatusCreated {
			t.Errorf("expected status code to be %v; got %v", http.StatusCreated, r.StatusCode)
		}
	})
}
