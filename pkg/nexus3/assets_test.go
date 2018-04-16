package nexus3

import "testing"

func TestGetAssets(t *testing.T) {
	resp, _, err := nexus.GetAssets("maven-releases")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("expected status code 200 got %d\n", resp.StatusCode)
	}
}
