package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/hytkgami/hrmos-closing-validator/domain"
)

const apiRoot = "https://ieyasu.co/api"

func host() string {
	company := os.Getenv("COMPANY_ID")
	return fmt.Sprintf("%s/%s/v1/", apiRoot, company)
}

func httpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func get(ctx context.Context, path string, params map[string]string, additionalHeader map[string]string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, host()+path, nil)
	if err != nil {
		return nil, err
	}
	queries := request.URL.Query()
	for key, value := range params {
		queries.Set(key, value)
	}
	request.URL.RawQuery = queries.Encode()
	addHeaders(request, additionalHeader)
	response, err := httpClient().Do(request)
	if err != nil {
		return nil, err
	}
	return extract(response)
}

func delete(ctx context.Context, path string, data []byte, additionalHeader map[string]string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodDelete, host()+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	addHeaders(request, additionalHeader)
	response, err := httpClient().Do(request)
	if err != nil {
		return nil, err
	}
	return extract(response)
}

func addHeaders(request *http.Request, headers map[string]string) {
	request.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		request.Header.Add(key, value)
	}
}

func extract(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if (response.StatusCode / 100) != 2 {
		var errBody *domain.Error
		if err := json.Unmarshal(body, &errBody); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("HTTP Status %d, code: %d, message: %s", response.StatusCode, errBody.Code, errBody.Message)
	}
	return body, nil
}
