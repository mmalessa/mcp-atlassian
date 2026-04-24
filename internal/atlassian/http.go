package atlassian

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func atlassianRequest(method, url string, body []byte) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(
		os.Getenv("ATLASSIAN_EMAIL"),
		os.Getenv("ATLASSIAN_API_TOKEN"),
	)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("atlassian API %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
