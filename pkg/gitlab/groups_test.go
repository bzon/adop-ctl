package gitlab

import (
	"testing"
)

var group = Group{
	Name: "Ex Batallion",
	Path: "exb",
}

func TestCreateGroup(t *testing.T) {
	// Test SetUp
	// Ensure group is deleted
	_, err := gitlab.DeleteGroupByPath(group.Path)
	if err != nil {
		t.Logf("not deleting %s", group.Path)
	}
	_, _, err = gitlab.CreateGroup(group)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchGroup(t *testing.T) {
	_, group, err := gitlab.SearchGroup(group.Path)
	if err != nil {
		t.Fatal(err)
	}
	if group.ID < 1 {
		t.Fatalf("group %s not found!", group.Path)
	}
}
