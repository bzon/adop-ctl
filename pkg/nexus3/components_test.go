package nexus3

import (
	"testing"
)

func TestGetComponents(t *testing.T) {
	resp, _, err := nexus.GetComponents("maven-releases")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected status code 200 got %d\n", resp.StatusCode)
	}
}

func TestDeleteComponent(t *testing.T) {
	resp, err := nexus.DeleteComponentByGroup()
	if err != nil {
		t.Fatal(err)
	}
	if resrp.StatusCode != 204 {
		t.Fatalf("expected status code 204 got %d\n", resp.StatusCode)
	}
}
