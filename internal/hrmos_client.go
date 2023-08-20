package internal

import (
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
	return fmt.Sprintf("%s/%s/", apiRoot, company)
}

func httpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func Get(ctx context.Context, path string, params map[string]string) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, host()+path, nil)
	if err != nil {
		return nil, err
	}
	queries := request.URL.Query()
	for key, value := range params {
		queries.Set(key, value)
	}
	request.URL.RawQuery = queries.Encode()
	response, err := httpClient().Do(request)
	if err != nil {
		return nil, err
	}
	return extract(response)
}

func extract(response *http.Response) (*http.Response, error) {
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
		return nil, fmt.Errorf("status code: %d, message: %s", errBody.Code, errBody.Message)
	}
	return response, nil
}
