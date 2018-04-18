package openldap

import (
	"fmt"
	"os"
	"testing"

	"github.com/bzon/adop-ctl/pkg/gitlab"
)

var openldap = &Client{
	Host:   "localhost",
	Scheme: "tcp",
	Port:   389,
}

var GitlabAPI = gitlab.API{
	HostURL: "http://localhost:10080/api/v4",
	Token:   os.Getenv("GITLAB_PRIVATE_TOKEN"),
}

func TestGetGroup(t *testing.T) {

	// get nx-admin and administrators group
	group, err := openldap.GetGroup("dc=ldap,dc=adop,dc=com", "administrators", "nx-admin")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// print results
	for i := range group {
		for j := 0; j < len(group[i].uniqueMembers); j++ {
			fmt.Printf("Group: %s, Member: %s\n", group[i].cn, group[i].uniqueMembers[j])
		}
	}

}
func TestGetGroupList(t *testing.T) {

	// test get group list
	groups, err := openldap.GetGroupList("dc=ldap,dc=adop,dc=com")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// print results
	for j := 0; j < len(groups); j++ {
		fmt.Println("Groups:" + groups[j])
	}

}
func TestSyncGroup(t *testing.T) {

	// get group list
	groupList, err := openldap.GetGroupList("dc=ldap,dc=adop,dc=com")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// get groups
	groups, err := openldap.GetGroup("dc=ldap,dc=adop,dc=com", groupList...)

	// sync groups
	openldap.SyncGitlabGroup(groups, GitlabAPI)
}
