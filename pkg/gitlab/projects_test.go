package gitlab

import (
	"testing"
)

var project = Project{
	Path:        "worldpeace",
	Description: "promotes world peace",
}

func TestDeleteProjectByPath(t *testing.T) {
	_, err := gitlab.CreateProject(
		&Project{
			Path: "root/dummy",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	_, err = gitlab.DeleteProjectByPath("root/dummy")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateProject(t *testing.T) {

}
