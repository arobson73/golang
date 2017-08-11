package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Model
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func main() {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},//get your own access :token
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	repo, _, err := client.Repositories.Get(context, "arobson73", "golang")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &Package{
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:  *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}

	fmt.Printf("%+v\n", pack)

	commitInfo, _, err := client.Repositories.ListCommits(context, "arobson73", "golang", nil)

	if err != nil {
		fmt.Printf("Problem in commit information %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", commitInfo[0]) // Last commit information
	opt := &github.IssueListByRepoOptions{
		Assignee: "arobson73",
	}
	iss, r, err := client.Issues.ListByRepo(context, "arobson73", "golang", opt)
	//	iss, r, err := client.Issues.ListByRepo(context, "arobson73", "golang", nil)

	if err != nil {
		fmt.Printf("Issues.ListByRepo returned error: %v\n", err)
	}
	fmt.Printf("Values is Issues\n")
	fmt.Printf("%+v %s\n", iss, iss)

	fmt.Printf("Values is rep\n")
	fmt.Printf("%+v\n", r)

}
