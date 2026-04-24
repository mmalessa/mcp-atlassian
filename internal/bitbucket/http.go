package bitbucket

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const defaultAPIBase = "https://api.bitbucket.org/2.0"

func apiBase() string {
	if v := os.Getenv("BITBUCKET_API_URL"); v != "" {
		return v
	}
	return defaultAPIBase
}

// get performs a GET against the Bitbucket API with basic auth.
// It does not set an Accept header, so endpoints returning JSON or plain text
// (e.g. /diff) both work without content-negotiation 406s.
func get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(
		os.Getenv("BITBUCKET_EMAIL"),
		os.Getenv("BITBUCKET_API_TOKEN"),
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("bitbucket API %d: %s", resp.StatusCode, string(body))
	}
	return body, nil
}

// prURL builds the base URL for a pull request. Workspace comes from env,
// repo is passed in by the caller (provided per-tool-call by the LLM).
func prURL(repo string, id int) string {
	return fmt.Sprintf("%s/repositories/%s/%s/pullrequests/%d",
		apiBase(),
		os.Getenv("BITBUCKET_WORKSPACE"),
		repo,
		id,
	)
}
