package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	<td><a href='{{.CacheURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.CacheURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var issueTemplate = template.Must(template.New("issue").Parse(`
<h1>{{.Title}}</h1>	
<d1>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</d1>
<p>{{.Body}}</p>
`))

// IssueCache issue缓存
type IssueCache struct {
	Issues         []Issue
	IssuesByNumber map[int]Issue
}

// NewIssueCatch 新的issue
func NewIssueCatch(owner, repo string) (ic IssueCache, err error) {
	issues, err := GetIssues(owner, repo)
	if err != nil {
		return
	}
	ic.Issues = issues
	ic.IssuesByNumber = make(map[int]Issue, len(issues))
	for _, issue := range issues {
		ic.IssuesByNumber[issue.Number] = issue
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
		logNonNil(issueListTemplate.Execute(w, ic))
		return
	}

	numStr := pathParts[2]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("Issue number isn't a number: '%s'", numStr)))
		if err != nil {
			log.Printf("Error writing response for %s: %s", r.RequestURI, err)
		}
		return
	}

	issue, ok := ic.IssuesByNumber[num]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(fmt.Sprintf("No issue '%d'", num)))
		if err != nil {
			log.Printf("Error writing response for %s: %s", r.RequestURI, err)
		}
		return
	}

	logNonNil(issueTemplate.Execute(w, issue))
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: githubserver OWNER REPO")
		os.Exit(1)
	}
	owner, repo := os.Args[1], os.Args[2]
	issueCache, err := NewIssueCatch(owner, repo)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", issueCache)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
