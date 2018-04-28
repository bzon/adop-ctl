package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// User is a gitlab user struct
// Gitlab API doc: https://docs.gitlab.com/ce/api/users.html#user
type User struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Password      string `json:"password,omitempty"`
	ResetPassword bool   `json:"reset_password"`
	State         string `json:"state,omitempty"`
	Email         string `json:"email"`
	AvatarURL     string `json:"avatar_url,omitempty"`
	WebURL        string `json:"web_url,omitempty"`
	AccessLevel   int    `json:"access_level"`
	ExpiresAt     string `json:"expires_at,omitempty"`
	IsAdmin       bool   `json:"is_admin"`
}

const (
	_ = iota
	// GuestLevel access level starts with 10
	GuestLevel int = iota * 10
	// ReporterLevel is 20
	ReporterLevel
	// DeveloperLevel is 30
	DeveloperLevel
	// MasterLevel is 40
	MasterLevel
	// OwnerLevel is 50
	OwnerLevel
)

// CreateUser creates a gitlab user by accepting a struct of type User.
// It returns the client http response and a struct of type *User.
//
// API doc: https://docs.gitlab.com/ce/api/users.html#user-creation
func (gitlab *API) CreateUser(user User) (*http.Response, *User, error) {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return nil, nil, err
	}
	userBuffer := bytes.NewBuffer(userBytes)
	resp, err := gitlab.NewRequest("POST", "/users", userBuffer, http.StatusCreated)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	// Fill up the User struct with whatever is returned from the response Body
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, nil, err
	}
	return resp, &user, nil
}

// DeleteUser deletes a registered user by accepting a user id and returns the client http response.
//
// API doc: https://docs.gitlab.com/ce/api/users.html#user-deletion
func (gitlab *API) DeleteUser(userID int) (*http.Response, error) {
	resp, err := gitlab.NewRequest("DELETE", "/users/"+strconv.Itoa(userID), nil, http.StatusNoContent)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteUserByUsername deletes a user by accepting a username and returns the client http response.
// This is meant to be used over DeleteUser if the caller doesn't know the user id.
func (gitlab *API) DeleteUserByUsername(userName string) (*http.Response, error) {
	resp, user, err := gitlab.SearchUserByEmailOrUserName(userName)
	if err != nil {
		return nil, err
	}
	resp, err = gitlab.DeleteUser(user.ID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchUserByEmailOrUserName searches for a user by accepting a username or an email.
// It returns the client http response, and a struct of type *User that was found.
//
// API doc: https://docs.gitlab.com/ce/api/users.html#single-user
func (gitlab *API) SearchUserByEmailOrUserName(userToSearch string) (*http.Response, *User, error) {
	// You can search for users by email or username with: /users?search=John
	var searchParam string
	if strings.Contains(userToSearch, "@") {
		searchParam = "email"
	} else {
		searchParam = "username"
	}
	resp, err := gitlab.NewRequest("GET", "/users?"+searchParam+"="+userToSearch, nil, http.StatusOK)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, nil, err
	}
	if len(users) > 0 {
		return resp, &users[0], nil
	}
	return resp, nil, fmt.Errorf("No %s with %s is found", searchParam, userToSearch)
}
