package gitlab

import (
	"testing"
)

func TestAddMember(t *testing.T) {
	_, uid, gid, err := gitlab.AddMemberToGroup(user, group.Path)
	if err != nil {
		t.Fatal(err)
	}
	group.ID = gid
	if found, err := gitlab.MemberExistsInGroup(uid, gid); !found || err != nil {
		t.Fatalf("Member %d not found in Group %d", uid, gid)
	}
}

func TestRemoveMemberFromGroup(t *testing.T) {
	_, uid, gid, err := gitlab.RemoveMemberFromGroup(user.Username, group.Path)
	if err != nil {
		t.Fatal(err)
	}
	group.ID = gid
	if found, err := gitlab.MemberExistsInGroup(uid, gid); found || err != nil {
		t.Fatalf("Member %d is still found in Group %d", uid, gid)
	}
}
