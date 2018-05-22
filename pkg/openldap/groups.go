package openldap

import (
	"fmt"
	"log"
	"strings"

	"github.com/xanzy/go-gitlab"
	ldap "gopkg.in/ldap.v2"
)

// Group is an LDAP Group struct
type Group struct {
	CN           string
	uniqueMember []string
}

// GetGroup gets specific group(s) and return an array of openldap.Group
func (openldap *Client) GetGroup(baseDN string, groupName ...string) ([]Group, error) {

	// Initialize Slice depending of number of groupnames
	var groups = make([]Group, len(groupName))

	// Loop through Groupnames
	for j := range groups {

		// Get Matched CNs
		cn, err := openldap.NewSearch("ou=groups,"+baseDN, "(&(objectClass=groupOfUniqueNames)(cn="+groupName[j]+"))", "cn")
		if err != nil {
			return nil, err
		}
		// Loop through Matched CNs
		for i, value := range cn {

			// Assign Group with exact match only
			if value == groupName[j] {
				uniqueMembers, err := openldap.NewSearch("cn="+cn[i]+",ou=groups,"+baseDN, "(&(objectClass=groupOfUniqueNames))", "uniqueMember")
				if err != nil {
					return nil, err
				}

				// Initialize only the group with the exact value
				groups[j] = Group{
					CN:           cn[i],
					uniqueMember: uniqueMembers,
				}
			}

		}
	}

	return groups, nil
}

// GetGroupList gets List of groups under ou=groups
func (openldap *Client) GetGroupList(baseDN string) ([]string, error) {

	// Get Group List
	groups, err := openldap.NewSearch("ou=groups,"+baseDN, "(&(objectClass=groupOfUniqueNames))", "cn")
	if err != nil {
		return nil, fmt.Errorf("Failed to get groups. %s", err)
	}
	return groups, nil

}

// SyncGitlabGroup accepts list of groups
func (openldap *Client) SyncGitlabGroup(ldapGroup []Group, GitlabAPI *gitlab.Client) error {

	// loop through groups
	for j := 0; j < len(ldapGroup); j++ {
		groupPath := "ldap_" + ldapGroup[j].CN
		gitlabGroup := gitlab.CreateGroupOptions{
			Name:        gitlab.String(ldapGroup[j].CN),
			Path:        gitlab.String(groupPath),
			Description: gitlab.String("Auto generated group \"" + ldapGroup[j].CN + "\" by ldap sync"),
			Visibility:  gitlab.Visibility(gitlab.PrivateVisibility),
		}

		// Delete group
		GitlabAPI.Groups.DeleteGroup(groupPath)

		// Create Group
		group, _, err := GitlabAPI.Groups.CreateGroup(&gitlabGroup)
		if err != nil {
			return err
		}

		// Loop through group
		for i := 0; i < len(ldapGroup[j].uniqueMember); i++ {
			// concatinate cn to get username
			name := strings.Split(strings.Split(ldapGroup[j].uniqueMember[i], ",")[0], "=")[1]

			// get user id
			opts := &gitlab.ListUsersOptions{
				Username: &name,
			}

			userlist, _, _ := GitlabAPI.Users.ListUsers(opts)
			// if more than 1 result is returned or no entry is found
			if len(userlist) > 2 || len(userlist) == 0 {
				return err
			}

			add := &gitlab.AddGroupMemberOptions{
				UserID:      gitlab.Int(userlist[0].ID),
				AccessLevel: gitlab.AccessLevel(30),
			}

			// add member to group
			_, _, err := GitlabAPI.GroupMembers.AddGroupMember(groupPath, add)
			if err != nil {
				return err
			}
			log.Printf("ldap_client: username=%s (gitlab user id %d) has been added to group=%s (gitlab group id %d) with access_level=%v\n", name, userlist[0].ID, groupPath, group.ID, add.AccessLevel)

		}
	}
	return nil
}

// CreateGroup Ka allergy yung warnings
func (openldap *Client) CreateGroup(baseDN string, ldapGroup Group) error {
	addRequest := ldap.NewAddRequest("cn=" + ldapGroup.CN + ",ou=groups," + baseDN)

	// default attributes for adop ldap group
	addRequest.Attribute("objectClass", []string{"groupOfUniqueNames", "top"})

	// assign values
	addRequest.Attribute("cn", []string{ldapGroup.CN})
	addRequest.Attribute("uniqueMember", ldapGroup.uniqueMember)

	// Add group
	return openldap.AddEntry(addRequest)

}

// DeleteGroup Ka allergy yung wanrnings
func (openldap *Client) DeleteGroup(baseDN string, ldapGroup Group) error {

	// Create Delete Request
	deleteRequest := ldap.NewDelRequest("cn="+ldapGroup.CN+",ou=groups,"+baseDN, nil)

	// Delete Group
	return openldap.DeleteEntry(deleteRequest)

}
