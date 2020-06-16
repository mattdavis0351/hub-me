package github

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	token = os.Getenv("HUB_ME_TOKEN")
	ctx   = context.Background()
	ts    = oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc     = oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
)
