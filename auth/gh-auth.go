package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	gh "golang.org/x/oauth2/github"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     getEnvVariable("CLIENT_ID"),
		ClientSecret: getEnvVariable("CLIENT_SECRET"),
		Scopes:       []string{"user:email", "repo"},
		Endpoint:     gh.Endpoint,
	}
	// TODO Make this string random!
	oauthString = "randstring45454"
	oauthToken  *oauth2.Token

	// UserLogin needs a comment
	UserLogin = false
)

// Login

// HandleGitHubLogin needs comment
func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthString, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Callback from GitHub

// HandleGitHubCallback needs comment
func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthString, state)
		return

	}
	fmt.Printf("Request URL is:\n%v\n", r.URL)
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Failure to parse form: %v\n", err)
	}

	code := r.FormValue("code")
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
	}

	oauthToken = token
	// TODO remove client call here, this is great for testing
	c := CreateGitHubClient()
	user, _, err := c.Users.Get(context.Background(), "")
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	UserLogin = true
	fmt.Printf("UserLogin is now: %v\n", UserLogin)
	fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)
}

// CreateGitHubClient needs comment
func CreateGitHubClient() *github.Client {

	oauthClient := oauthConf.Client(context.Background(), oauthToken)
	client := github.NewClient(oauthClient)

	return client

}
