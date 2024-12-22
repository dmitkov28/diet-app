package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIClient struct {
	HTTPClient *http.Client
}

func NewAPIClient(httpClient *http.Client)*APIClient{
	return &APIClient{HTTPClient: httpClient}
}

func (c *APIClient) NewRequest(method, endpoint string, body interface{}) (*http.Request, error) {

	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)

		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, endpoint, bodyReader)

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil
}

func (c *APIClient) Do(req *http.Request, response interface{}) error {
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	if resp != nil {
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}
	return nil
}
