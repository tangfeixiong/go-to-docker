/*
  Refer to:
  - https://github.com/google/go-github
*/
package gogithub

import (
	"context"
	"fmt"

	"github.com/google/go-github/v18/github"

	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("List repositories: %v", repos)
}
