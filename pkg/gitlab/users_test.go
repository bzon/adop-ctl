package gitlab

import (
	"testing"
	"time"
)

var user = User{
	Name:        "John Smith",
	Username:    "john",
	Password:    "12345678",
	AccessLevel: OwnerLevel,
	Email:       "john@example.com",
}

func TestSearchUserByEmailOrUserName(t *testing.T) {
	// Search for root user which is by default has a user id of 1 and is an admin
	_, user, err := gitlab.SearchUserByEmailOrUserName("root")
	if err != nil {
		t.Fatal(err)
	}
	if user.ID != 1 && !user.IsAdmin {
		t.Fatal("User search failed")
	}
}

func TestCreateUser(t *testing.T) {
	// Test SetUp
	// Ensure user is deleted
	_, err := gitlab.DeleteUserByUsername(user.Username)
	// If err is nil, then the user exists, ensure that the test user is deleted
	if err == nil {
		// Wait for 10 seconds for the user to be deleted
		time.Sleep(10 * time.Second)
		_, deletedUser, _ := gitlab.SearchUserByEmailOrUserName(user.Username)
		if deletedUser != nil {
			t.Fatal("User is still not deleted after 10 seconds")
		}
	}
	_, createdUser, err := gitlab.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	if createdUser.State != "active" {
		t.Fatalf("Failed creating user %s\n", user.Username)
	}
}
