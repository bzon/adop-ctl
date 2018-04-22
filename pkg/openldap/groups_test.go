package openldap

import (
	"fmt"
	"os"
	"testing"

	"github.com/bzon/adop-ctl/pkg/gitlab"
)

var GitlabAPI = gitlab.API{
	HostURL: "http://localhost:10080/api/v4",
	Token:   os.Getenv("GITLAB_PRIVATE_TOKEN"),
}

var ldapGroup = Group{
	CN:           "one.direction",
	uniqueMember: []string{"cn=admin,dc=ldap,dc=adop,dc=com", "cn=john.smith,ou=people,dc=ldap,dc=adop,dc=com"},
}

func TestGetGroup(t *testing.T) {
	// get nx-admin and administrators group
	group, err := openldap.GetGroup(ldapDomain, "administrators", "nx-admin")
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	// print results
	for i := range group {
		for j := 0; j < len(group[i].uniqueMember); j++ {
			fmt.Printf("Group: %s, Member: %s\n", group[i].CN, group[i].uniqueMember[j])
		}
	}
}

func TestGetGroupList(t *testing.T) {
	// test get group list
	groups, err := openldap.GetGroupList(ldapDomain)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	// print results
	for j := 0; j < len(groups); j++ {
		fmt.Println("Groups:" + groups[j])
	}
}

func TestSyncGroup(t *testing.T) {
	// get group list
	groupList, err := openldap.GetGroupList(ldapDomain)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	// get groups
	groups, err := openldap.GetGroup(ldapDomain, groupList...)

	// sync groups
	openldap.SyncGitlabGroup(groups, GitlabAPI)
}

func TestCreateGroup(t *testing.T) {

	// Run Delete function dont fetch errors
	openldap.DeleteGroup(ldapDomain, ldapGroup)

	// Run Create function
	err := openldap.CreateGroup(ldapDomain, ldapGroup)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

}

func TestDeleteGroup(t *testing.T) {

	// Delete User
	err := openldap.DeleteGroup(ldapDomain, ldapGroup)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

}
