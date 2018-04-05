package gitlab

import (
	"net/http"
	"strconv"
)

// AddMemberToGroup adds a user to a group. It also returns the uid of the added member and the uid of the group
//
// API doc: https://docs.gitlab.com/ce/api/members.html#add-a-member-to-a-group-or-project
func (gitlab *API) AddMemberToGroup(member User, groupPath string) (*http.Response, int, int, error) {
	// Verify user exists and get user id
	_, user, err := gitlab.SearchUserByEmailOrUserName(member.Username)
	if err != nil {
		return nil, -1, -1, err
	}
	// Verify group exists and get group id
	_, group, err := gitlab.SearchGroup(groupPath)
	if err != nil {
		return nil, -1, -1, err
	}
	// Add user to group
	resp, err := gitlab.NewRequest("POST",
		"/groups/"+strconv.Itoa(group.ID)+"/members?access_level="+strconv.Itoa(member.AccessLevel)+"&user_id="+strconv.Itoa(user.ID),
		nil,
		http.StatusCreated,
	)
	if err != nil {
		return nil, -1, -1, err
	}
	return resp, user.ID, group.ID, nil
}

// RemoveMemberFromGroup removes a user from a group
//
// API doc: https://docs.gitlab.com/ce/api/members.html#remove-a-member-from-a-group-or-project
func (gitlab *API) RemoveMemberFromGroup(username string, groupPath string) (*http.Response, int, int, error) {
	// Verify user exists and get user id
	_, user, err := gitlab.SearchUserByEmailOrUserName(username)
	if err != nil {
		return nil, -1, -1, err
	}
	// Verify group exists and get group id
	_, group, err := gitlab.SearchGroup(groupPath)
	if err != nil {
		return nil, -1, -1, err
	}
	// Remove user from group
	resp, err := gitlab.NewRequest("DELETE",
		"/groups/"+strconv.Itoa(group.ID)+"/members/"+strconv.Itoa(user.ID),
		nil,
		http.StatusNoContent,
	)
	if err != nil {
		return nil, -1, -1, err
	}
	return resp, user.ID, group.ID, nil
}
