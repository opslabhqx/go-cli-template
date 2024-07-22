package auth

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

var (
	Client      *github.Client
	GithubToken string
)

func LoginGh() {
	InitGhSession()
}

func InitGhSession() {
	GithubToken = os.Getenv("GITHUB_TOKEN")

	if GithubToken == "" {
		log.Fatal("Missing environment variable: GITHUB_TOKEN")
	}

	if Client == nil {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: GithubToken})
		tc := oauth2.NewClient(ctx, ts)
		Client = github.NewClient(tc)
	}
}
