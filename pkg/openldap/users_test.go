package openldap

import (
	"testing"
)

var ldapUser = User{
	CN:           "zayn.malik",
	displayName:  "ZAYN",
	givenName:    "Zayn Malik",
	mail:         "zayn.malika@adop.ldap.com",
	SN:           "User",
	UID:          "zayn.malik",
	userPassword: "123qwe123",
}

func TestCreateUser(t *testing.T) {

	// Run Delete function dont fetch errors
	openldap.DeleteUser(ldapDomain, ldapUser)

	// Create User
	err := openldap.CreateUser(ldapDomain, ldapUser)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
}

func TestDeleteUser(t *testing.T) {

	// Delete User
	err := openldap.DeleteUser(ldapDomain, ldapUser)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

}

func TestAddUserToGroup(t *testing.T) {

	// Create Group to make sure test runs sucessfully
	checker := openldap.CreateGroup(ldapDomain, ldapGroup)

	// Add user to group
	err := openldap.AddUserToGroup(ldapDomain, ldapUser, ldapGroup)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	// If checker is empty means group was only created in this function
	if checker == nil {
		openldap.DeleteGroup(ldapDomain, ldapGroup)

	}

}
