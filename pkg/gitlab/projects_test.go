package gitlab

import (
	"testing"
)

var testProject = &Project{
	Name:        "worldpeace",
	Path:        "worldpeace",
	Description: "promotes world peace",
	NamespaceID: 1, // root user namespace
}

func TestCreateProject(t *testing.T) {
	_, err := gitlab.CreateProject(testProject)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteProjectByPath(t *testing.T) {
	_, err := gitlab.DeleteProjectByPath("root/" + testProject.Path)
	if err != nil {
		t.Fatal(err)
	}
}
