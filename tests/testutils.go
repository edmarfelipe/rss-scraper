package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type requestBuilder struct {
	method  string
	body    []byte
	headers map[string]string
	url     string
	ctx     context.Context
}

func NewRequest(url string) *requestBuilder {
	return &requestBuilder{url: url, headers: make(map[string]string)}
}

func (r *requestBuilder) WithMethod(method string) *requestBuilder {
	r.method = method
	return r
}

func (r *requestBuilder) WithHeader(key, value string) *requestBuilder {
	r.headers[key] = value
	return r
}

func (r *requestBuilder) WithBody(body string) *requestBuilder {
	r.body = []byte(body)
	return r
}

func (r *requestBuilder) Send() (*http.Response, error) {
	req, err := http.NewRequest(r.method, r.url, bytes.NewReader(r.body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}
	return http.DefaultClient.Do(req)
}

func decode[T any](r *http.Response) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
