package openldap

import (
	"github.com/bzon/adop-ctl/pkg/gitlab"
)

// Group is an LDAP Group struct
type Group struct {
}

func SearchGroup(name string) (*Group, error) {

	return nil, nil
}

func SyncGroup(ldapGroup Group, gitlabGroup gitlab.Group) error {
	return nil
}
