package nexus3

import (
	"encoding/json"
	"net/http"
)

// AssetItems is a slice of Asset
//
// API doc: /swagger-ui/#!/assets/getAssets
type AssetItems struct {
	Items             []Asset `json:"items"`
	ContinuationToken string  `json:"continuationToken"`
}

// Asset is the assets/:id model
//
// API doc: /swagger-ui/#!/assets/getAssetById
type Asset struct {
	DownloadURL string `json:"download_url"`
	Path        string `json:"path"`
	ID          string `json:"id"`
	Repository  string `json:"repository"`
	Format      string `json:"format"`
	Checksum    `json:"checksum"`
}

// Checksum is part of Asset that has SHA1 and MD5
type Checksum struct {
	SHA1 string `json:"sha1"`
	MD5  string `json:"md5"`
}

// GetAssets list all assets of a repository
func (nexus *API) GetAssets(repoID string) (*http.Response, *AssetItems, error) {
	resp, err := nexus.NewRequest("GET", repoID, "assets", nil, 200)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	var assets AssetItems
	if err := json.NewDecoder(resp.Body).Decode(&assets); err != nil {
		return nil, nil, err
	}
	return resp, &assets, nil
}
