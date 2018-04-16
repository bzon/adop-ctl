package nexus3

import (
	"encoding/json"
	"net/http"
)

// ComponentItems is a slice of Component
//
// API doc: /swagger-ui/#!/components/getComponents
type ComponentItems struct {
	Items             []Component `json:"items"`
	ContinuationToken string      `json:"continuationToken"`
}

// Component is the components/:id model
//
// API doc: /swagger-ui/#!/components/getComponentById
type Component struct {
	ID         string  `json:"id"`
	Repository string  `json:"repository"`
	Format     string  `json:"format"`
	Group      string  `json:"group"`
	Name       string  `json:"name"`
	Version    string  `json:"version"`
	Assets     []Asset `json:"assets"`
}

func (nexus *API) GetComponents(repoID string) (*http.Response, *ComponentItems, error) {
	resp, err := nexus.NewRequest("GET", repoID, "components", nil, http.StatusOK)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	var components ComponentItems
	if err := json.NewDecoder(resp.Body).Decode(&components); err != nil {
		return nil, nil, err
	}
	return resp, &components, nil
}
