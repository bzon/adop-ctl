package gitlab

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// API contains the Gitlab HostURL and Token strings
type API struct {
	HostURL string
	Token   string
}

// NewRequest creates an http API request to a gitlab instance
func (gitlab *API) NewRequest(method, requestURI string, body io.Reader, expectedStatus int) (*http.Response, error) {
	req, err := http.NewRequest(method, gitlab.HostURL+"/"+requestURI, body)
	if err != nil {
		return nil, fmt.Errorf("failed creating a new http request: %v", err)
	}

	// Add gitlab request headers
	req.Header.Add("PRIVATE-TOKEN", gitlab.Token)
	req.Header.Add("Content-Type", "application/json")

	// Execute the query
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed Doing client: %v", err)
	}

	// Debug mode
	log.Printf("method=%s request_uri=%s status_code=%d expected_status_code=%d\n", method, requestURI, resp.StatusCode, expectedStatus)

	if resp.StatusCode != expectedStatus {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body +%v", err)
		}
		return nil, fmt.Errorf("reponse=%s", string(respBody))
	}
	return resp, nil
}
