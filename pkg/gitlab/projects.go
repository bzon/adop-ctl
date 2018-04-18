package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Project JSON fields
type Project struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Path                string `json:"path"`
	Description         string `json:"description,omitempty"`
	DefaultBranch       string `json:"default_branch,omitempty"`
	Visibility          string `json:"visibility,omitempty"`
	IssuesEnabled       bool   `json:"issues_enabled,omitempty"`
	MergeRequestEnabled bool   `json:"merge_request_enabled,omitempty"`
	NamespaceID         int    `json:"namespace_id,omitempty"`
}

// ProjectHook JSON fields
type ProjectHook struct {
	ID                       int    `json:"id"`
	URL                      string `json:"url"`
	PushEvents               bool   `json:"push_events"`
	IssuesEvents             bool   `json:"issues_events"`
	ConfidentialIssuesEvents bool   `json:"confidential_issues_events"`
	MergeRequestsEvents      bool   `json:"merge_requests_events"`
	TagPushEvents            bool   `json:"tag_push_events"`
	NoteEvents               bool   `json:"note_events"`
	JobEvents                bool   `json:"job_events"`
	PipelineEvents           bool   `json:"pipeline_events"`
	WikiPageEvents           bool   `json:"wiki_page_events"`
	EnableSSLVerification    bool   `json:"enable_ssl_verification"`
	Token                    string `json:"token"`
}

// CreateProject creates a gitlab project using Project and a path where to create the project
//
// API doc: https://docs.gitlab.com/ce/api/projects.html#create-project
func (gitlab *API) CreateProject(project *Project) (*http.Response, error) {
	projectBytes, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}
	projectBuffer := bytes.NewBuffer(projectBytes)
	resp, err := gitlab.NewRequest("POST", "projects", projectBuffer, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(resp.Body).Decode(project); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSingleProject returns the project id
//
// API doc: https://docs.gitlab.com/ce/api/projects.html#get-single-project
func (gitlab *API) GetSingleProject(projectID int) (*http.Response, *Project, error) {
	resp, err := gitlab.NewRequest("GET", "projects/"+strconv.Itoa(projectID), nil, http.StatusOK)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting single project: %+v", err)
	}
	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, nil, fmt.Errorf("failed decoding: %+v", err)
	}
	return resp, &project, nil
}

// GetProjectByPath returns the project id using the project path
//
// API doc: https://docs.gitlab.com/ce/api/projects.html#list-all-projects
func (gitlab *API) GetProjectByPath(projectPath string) (*http.Response, *Project, error) {
	encodedPath := url.QueryEscape(projectPath)
	resp, err := gitlab.NewRequest("GET", "projects/"+encodedPath, nil, http.StatusOK)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting project by path: %+v", err)
	}
	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, nil, fmt.Errorf("failed decoding project %+v", err)
	}
	return resp, &project, nil
}

// DeleteProject deletes a single project
//
// API doc: https://docs.gitlab.com/ce/api/projects.html#remove-project
func (gitlab *API) DeleteProject(projectID int) (*http.Response, error) {
	resp, err := gitlab.NewRequest("DELETE", "projects/"+strconv.Itoa(projectID), nil, http.StatusAccepted)
	if err != nil {
		return nil, fmt.Errorf("failed deleting project: %+v", err)
	}
	return resp, nil
}

// DeleteProjectByPath deletes a project using project path
func (gitlab *API) DeleteProjectByPath(projectPath string) (*http.Response, error) {
	_, project, err := gitlab.GetProjectByPath(projectPath)
	if err != nil {
		return nil, err
	}
	resp, err := gitlab.DeleteProject(project.ID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListProjectHooks returns webhooks configured for the project
//
// API doc: https://docs.gitlab.com/ce/api/projects.html#list-project-hooks
func (gitlab *API) ListProjectHooks(projectPath string) {
}

// DeleteProjectHooks deletes webhooks in a project
//
// API doc: https://docs.gitlab.com/ce/api/projects.html#delete-project-hook
func (gitlab *API) DeleteProjectHooks() {
}

// CreateProjectHook adds a webhook in a project
//
// API doc: https://docs.gitlab.com/ce/api/projects.html#add-project-hook
func (gitlab *API) CreateProjectHook(hook ProjectHook) {
}
