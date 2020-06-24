package utils

import (
	"context"

	"github.com/mattdavis0351/hub-me/auth"
)

var (
	// token = os.Getenv("HUB_ME_TOKEN")
	ctx = context.Background()
	// ts    = oauth2.StaticTokenSource(
	// 	&oauth2.Token{AccessToken: token},
	// )
	// tc     = oauth2.NewClient(ctx, ts)
	client = auth.CreateGitHubClient()
)
