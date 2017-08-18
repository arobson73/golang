package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

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
		&oauth2.Token{AccessToken: ""}, //get your own access token
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

	//now edit a github issue
	v := "you little git"

	input := &github.IssueRequest{Title: &v}
	//input := &IssueRequest{Title: &v}

	issnew, _, err := client.Issues.Edit(context, "arobson73", "golang", 1, input)
	if err != nil {
		fmt.Printf("Issues.Edit returned error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Issue is now called %s\n", *issnew.Title)

	//now do this editing via vim
	//first get the issue
	issnew1, _, err := client.Issues.Get(context, "arobson73", "golang", 1)
	if err != nil {
		fmt.Printf("Problem getting issue golang number 1")
		os.Exit(1)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}
	tempfile, err := ioutil.TempFile("", "issue_crud")
	if err != nil {
		log.Fatal(err)
	}
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	encoder := json.NewEncoder(tempfile)
	err = encoder.Encode(map[string]string{
		"title": *issnew1.Title,
		"state": *issnew1.State,
		"body":  *issnew1.Body,
	})
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	tempfile.Seek(0, 0)
	fields := new(github.IssueRequest)
	if err = json.NewDecoder(tempfile).Decode(&fields); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Title:%s\n", *fields.Title)
	fmt.Printf("State:%s\n", *fields.State)
	fmt.Printf("Body:%s\n", *fields.Body)

	issnew2, _, err := client.Issues.Edit(context, "arobson73", "golang", 1, fields)
	if err != nil {
		fmt.Printf("Issues.Edit returned error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("edited title is %s \n", issnew2.GetTitle())

}
