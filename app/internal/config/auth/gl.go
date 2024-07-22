package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

var (
	Git         *gitlab.Client
	GitlabToken string
	GitlabHost  string
)

func LoginGl() {
	InitGlSession()
}

func InitGlSession() {
	GitlabToken = os.Getenv("GITLAB_TOKEN")
	GitlabHost = os.Getenv("GITLAB_HOST")

	var missingVars []string

	if GitlabToken == "" {
		missingVars = append(missingVars, "GITLAB_TOKEN")
	}

	if GitlabHost == "" {
		missingVars = append(missingVars, "GITLAB_HOST")
	}

	if len(missingVars) > 0 {
		errorMessage := fmt.Sprintf("Missing environment variables: %v", missingVars)
		log.Fatal(errorMessage)
	}

	if Git == nil {
		var e error
		Git, e = gitlab.NewClient(GitlabToken, gitlab.WithBaseURL(GitlabHost))
		if e != nil {
			log.Fatal(e)
		}
	}
}
