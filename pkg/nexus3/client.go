package nexus3

import (
	"io"
	"log"
	"net/http"
)

// API contains the fields for authenticating to Nexus
type API struct {
	HostURL, Username, Password string
}

// NewRequest creates a new REST request to Nexus and executes it
//
// API doc: /swagger-ui/
func (nexus *API) NewRequest(method, repoID string, requestURI string, body io.Reader, expectedStatus int) (*http.Response, error) {
	req, err := http.NewRequest(method, nexus.HostURL+"/service/rest/beta/"+requestURI+"?repository="+repoID, body)
	if err != nil {
		return nil, err
	}

	// Basic Authentication
	req.SetBasicAuth(nexus.Username, nexus.Password)
	// Add request headers
	req.Header.Add("Content-Type", "application/json")

	// Create and execute the client
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Debug mode
	log.Printf("method=%s request_uri=%s status_code=%d expected_status_code=%d repository=%s\n", method, requestURI, resp.StatusCode, expectedStatus, repoID)
	return resp, nil
}
