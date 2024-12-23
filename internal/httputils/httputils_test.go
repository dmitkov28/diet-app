package httputils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRequest(t *testing.T) {
	client := NewAPIClient(nil)

	tests := []struct {
		name      string
		method    string
		endpoint  string
		body      interface{}
		wantError bool
	}{
		{"valid GET request", http.MethodGet, "https://example.com", nil, false},
		{"valid POST request with body", http.MethodPost, "https://example.com", map[string]string{"key": "value"}, false},
		{"invalid body", http.MethodPost, "https://example.com", func() {}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := client.NewRequest(tt.method, tt.endpoint, tt.body)
			if (err != nil) != tt.wantError {
				t.Errorf("NewRequest() error = %v, wantError %v", err, tt.wantError)
			}
			if err == nil && req.Method != tt.method {
				t.Errorf("Expected method %s, got %s", tt.method, req.Method)
			}
		})
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		name           string
		mockStatusCode int
		mockBody       interface{}
		responseTarget interface{}
		wantError      bool
	}{
		{
			"successful response",
			http.StatusOK,
			map[string]string{"message": "success"},
			&map[string]string{},
			false,
		},
		{
			"error status code",
			http.StatusBadRequest,
			"error occurred",
			nil,
			true,
		},
		{
			"invalid JSON response",
			http.StatusOK,
			"not JSON",
			&map[string]string{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockResponse := httptest.NewRecorder()
			if bodyBytes, err := json.Marshal(tt.mockBody); err == nil {
				mockResponse.Body = bytes.NewBuffer(bodyBytes)
			} else if strBody, ok := tt.mockBody.(string); ok {
				mockResponse.Body = bytes.NewBufferString(strBody)
			}
			mockResponse.Code = tt.mockStatusCode

			mockHTTPClient := &http.Client{
				Transport: roundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: mockResponse.Code,
						Body:       io.NopCloser(mockResponse.Body),
						Header:     make(http.Header),
					}
				}),
			}

			client := NewAPIClient(mockHTTPClient)

			req, _ := client.NewRequest(http.MethodGet, "https://example.com", nil)
			err := client.Do(req, tt.responseTarget)
			if (err != nil) != tt.wantError {
				t.Errorf("Do() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}
