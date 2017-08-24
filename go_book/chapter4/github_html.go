package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	gh "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var issueListTemplate = template.Must(template.New("issueList").Parse(`
<h1>{{.Issues | len}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Issues}}
<tr>
	<td><a href='{{.URL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.URL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.URL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var issueTemplate = template.Must(template.New("issue").Parse(`
<h1>{{.Title}}</h1>
<dl>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</dl>
<p>{{.Body}}</p>
`))

const IssuesURL = "https://api.github.com/search/issues"
const APIURL = "https://api.github.com"

type IssueCache struct {
	Issues         []gh.Issue
	IssuesByNumber map[int]gh.Issue
}

func NewIssueCache(issues []*gh.Issue) (ic IssueCache) {

	ic.IssuesByNumber = make(map[int]gh.Issue, len(issues))
	for _, issue := range issues {
		ic.IssuesByNumber[*issue.Number] = *issue
		ic.Issues = append(ic.Issues, *issue)
	}
	return
}
func logNonNil(v interface{}) {
	if v != nil {
		log.Print(v)
	}
}

func (ic IssueCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	if len(pathParts) < 3 || pathParts[2] == "" {
		log.Printf("issueList Used")
		logNonNil(issueListTemplate.Execute(w, ic))
		return
	}
	numStr := pathParts[2]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("Issue number isn't a number: '%s'", numStr)))
		if err != nil {
			log.Printf("Error writing response for %s: %s", r, err)
		}
		return
	}
	issue, ok := ic.IssuesByNumber[num]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(fmt.Sprintf("No issue '%d'", num)))
		if err != nil {
			log.Printf("Error writing response for %s: %s", r, err)
		}
		return
	}
	log.Printf("issueTemplate Used")
	logNonNil(issueTemplate.Execute(w, issue))
}

func main() {
	github := os.Getenv("GITHUB_AUTH")
	//AUTH
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: github},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := gh.NewClient(tokenClient)
	//ISSUES BY REPO
	opt := &gh.IssueListByRepoOptions{
		Assignee: "arobson73",
	}
	iss, r, err := client.Issues.ListByRepo(context, "arobson73", "golang", opt)

	if err != nil {
		fmt.Printf("Issues.ListByRepo returned error: %v\n", err)
		return
	}
	if len(iss) == 0 {
		fmt.Printf("Issues.ListByRepo has no issues\n")
		return

	}

	fmt.Printf("%+v\n", iss)
	fmt.Printf("%+v\n", r)

	ic := NewIssueCache(iss)

	for k, v := range ic.IssuesByNumber {
		fmt.Printf("issue %d\n", k)
		fmt.Printf("%+v\n", v)

	}

	http.Handle("/", ic)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
