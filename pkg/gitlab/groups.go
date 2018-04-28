package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Group is a gitlab group
//
// API doc: https://docs.gitlab.com/ce/api/groups.html#details-of-a-group
type Group struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Path                 string `json:"path"`
	Description          string `json:"description"`
	Visibility           string `json:"visibility"`
	LFSEnabled           bool   `json:"lfs_enabled"`
	AvatarURL            string `json:"avatar_url,omitempty"`
	WebURL               string `json:"web_url"`
	RequestAccessEnabled bool   `json:"request_access_enabled"`
	FullName             string `json:"full_name"`
	FullPath             string `json:"full_path"`
	ParentID             int    `json:"parent_id,omitempty"`
}

// CreateGroup creates a group and returns the client http response and a struct of type *Group.
//
// API doc: https://docs.gitlab.com/ce/api/groups.html#new-group
func (gitlab *API) CreateGroup(g Group) (*http.Response, *Group, error) {
	// use QueryEscape because group names and paths may have whitespaces
	name := url.QueryEscape(g.Name)
	path := url.QueryEscape(g.Path)
	resp, err := gitlab.NewRequest("POST", "/groups?&name="+name+"&path="+path, nil, http.StatusCreated)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	var group Group
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, nil, err
	}
	return resp, &group, nil
}

// DeleteGroup deletes a group using group id and returns the client http response.
//
// API doc: https://docs.gitlab.com/ce/api/groups.html#remove-group
func (gitlab *API) DeleteGroup(groupID int) (*http.Response, error) {
	group := strconv.Itoa(groupID)
	resp, err := gitlab.NewRequest("DELETE", "/groups/"+group, nil, http.StatusNoContent)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

// DeleteGroupByPath deletes a group using group path and returns the client http response.
// It calls SearchGroup to get find the group id and use it in DeleteGroup.
func (gitlab *API) DeleteGroupByPath(path string) (*http.Response, error) {
	_, group, err := gitlab.SearchGroup(path)
	if err != nil {
		return nil, err
	}
	if group.ID > 0 {
		resp, err := gitlab.DeleteGroup(group.ID)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	return nil, fmt.Errorf("no group id found for group path %s", path)
}

// SearchGroup searches for a group given the group path and returns the client http response and a struct of type *Group.
//
// API doc: https://docs.gitlab.com/ce/api/groups.html#search-for-group
func (gitlab *API) SearchGroup(path string) (*http.Response, *Group, error) {
	resp, err := gitlab.NewRequest("GET", "/groups?search="+path, nil, http.StatusOK)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var groups []Group
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return nil, nil, err
	}
	for _, group := range groups {
		if group.Path == path {
			return resp, &group, nil
		}
	}
	return nil, nil, fmt.Errorf("%s group not found", path)
}

// MemberExistsInGroup searches for a gitlab member in a group given a userID and groupID and returns a boolean value.
//
// API doc: https://docs.gitlab.com/ce/api/members.html#list-all-members-of-a-group-or-project
func (gitlab *API) MemberExistsInGroup(userID, groupID int) (bool, error) {
	resp, err := gitlab.NewRequest(
		"GET",
		"/groups/"+strconv.Itoa(groupID)+"/members",
		nil,
		http.StatusOK,
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var members []User
	if err := json.NewDecoder(resp.Body).Decode(&members); err != nil {
		return false, err
	}
	for _, member := range members {
		if member.ID == userID {
			return true, nil
		}
	}
	return false, nil
}
