package nexus3

import (
	"testing"
)

var nexus = API{
	HostURL:  "http://localhost:8081",
	Username: "admin",
	Password: "admin123",
}

func TestNewRequest(t *testing.T) {
	resp, err := nexus.NewRequest("GET", "maven-releases", "assets", nil, 200)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("want status code 200 got %d", resp.StatusCode)
	}
}
