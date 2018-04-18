package gitlab

import "os"

var gitlab = &API{
	HostURL: "http://localhost:10080/api/v4",
	Token:   os.Getenv("GITLAB_PRIVATE_TOKEN"),
}
