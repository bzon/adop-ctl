package openldap

import (
	"fmt"
	"testing"

	gitlab "github.com/xanzy/go-gitlab"
)

var basehttpURL = "http://localhost:10080"
var gitlabUsername = "root"
var gitlabPassword = "123qwe123"

var ldapGroup = Group{
	CN:           "one.direction",
	uniqueMember: []string{"cn=admin,dc=ldap,dc=adop,dc=com", "cn=john.smith,ou=people,dc=ldap,dc=adop,dc=com"},
}

func TestGetGroup(t *testing.T) {
	// get nx-admin and administrators group
	group, err := openldap.GetGroup(ldapDomain, "administrators", "nx-admin")
	if err != nil {
		t.Fatal(err)
	}

	// print results
	for i, g := range group {
		fmt.Printf("(%d) Group: %s\n", i+1, g.CN)
		for _, u := range g.uniqueMember {
			fmt.Printf("- %s\n", u)
		}
	}
}

func TestGetGroupList(t *testing.T) {
	// test get group list
	groups, err := openldap.GetGroupList(ldapDomain)
	if err != nil {
		t.Fatal(err)
	}

	// print results
	for i, v := range groups {
		fmt.Printf("(%d) %s\n", i+1, v)
	}
}

func TestSyncGroup(t *testing.T) {
	git, err := gitlab.NewBasicAuthClient(nil,
		basehttpURL,
		gitlabUsername,
		gitlabPassword,
	)
	if err != nil {
		t.Fatal(err)
	}

	// get group list
	groupList, err := openldap.GetGroupList(ldapDomain)
	if err != nil {
		t.Fatal(err)
	}

	// get groups
	groups, err := openldap.GetGroup(ldapDomain, groupList...)
	if err != nil {
		t.Fatal(err)
	}

	// sync groups
	if err := openldap.SyncGitlabGroup(groups, git); err != nil {
		t.Fatal(err)
	}
}

func TestCreateGroup(t *testing.T) {

	// Run Delete function dont fetch errors
	if err := openldap.DeleteGroup(ldapDomain, ldapGroup); err != nil {
		t.Fatal(err)
	}

	// Run Create function
	err := openldap.CreateGroup(ldapDomain, ldapGroup)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDeleteGroup(t *testing.T) {

	// Delete User
	err := openldap.DeleteGroup(ldapDomain, ldapGroup)
	if err != nil {
		t.Fatal(err)
	}

}
