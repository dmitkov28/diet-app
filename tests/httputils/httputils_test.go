package httputils_test

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/dmitkov28/dietapp/internal/utils"
)

func TestNewRequest(t *testing.T) {
	httpClient := http.Client{}
	apiClient := utils.NewAPIClient(&httpClient)

	testUrl := "https://example.com"

	type testBody struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	tests := []struct {
		name          string
		method        string
		url           string
		body          interface{}
		wantErr       bool
		compareReq    *http.Request
		expectedBody  string
		expectedError string
	}{
		{
			name:    "Valid GET request",
			method:  http.MethodGet,
			url:     testUrl,
			wantErr: false,
			compareReq: &http.Request{
				Method: http.MethodGet,
				URL:    mustParseURL(testUrl),
				Header: http.Header{},
			},
		},
		{
			name:    "Invalid GET request (invalid URL)",
			method:  http.MethodGet,
			url:     "invalid url",
			wantErr: true,
			compareReq: &http.Request{
				Method: http.MethodGet,
				URL:    mustParseURL(testUrl),
				Header: http.Header{},
			},
			expectedError: "invalid URL: invalid url",
		},
		{
			name:    "Valid POST request",
			method:  http.MethodPost,
			url:     testUrl,
			body:    testBody{Name: "Test", Value: "Test"},
			wantErr: false,
			compareReq: &http.Request{
				Method: http.MethodPost,
				URL:    mustParseURL(testUrl),
				Header: http.Header{},
			},
			expectedBody: `{"name":"Test","value":"Test"}`,
		},
		{
			name:    "Invalid POST request (invalid body)",
			method:  http.MethodPost,
			url:     testUrl,
			body:    make(chan int),
			wantErr: true,
			compareReq: &http.Request{
				Method: http.MethodPost,
				URL:    mustParseURL(testUrl),
				Header: http.Header{},
			},
			expectedError: "failed to marshal body: json: unsupported type: chan int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := apiClient.NewRequest(tt.method, tt.url, tt.body)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error containing %q, got %q", tt.expectedError, err.Error())
				}
				return
			}

			if req.Method != tt.compareReq.Method {
				t.Errorf("Expected method: %s, got: %s", tt.compareReq.Method, req.Method)
			}

			if req.URL.String() != tt.compareReq.URL.String() {
				t.Errorf("Expected URL: %s, got: %s", tt.compareReq.URL, req.URL)
			}

			if tt.body != nil {
				bodyBytes, _ := io.ReadAll(req.Body)
				defer req.Body.Close()
				if string(bodyBytes) != tt.expectedBody {
					t.Errorf("Expected Body: %s, got: %s", tt.expectedBody, string(bodyBytes))
				}
			}
		})
	}
}

func mustParseURL(u string) *url.URL {
	parsed, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	return parsed
}
