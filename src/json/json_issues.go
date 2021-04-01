package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const tableDescription = `<h2> {{.}} </h2>`

const dataTemplHTML = `
<table>
	<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
	</tr>
	{{range .}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
</table>`

type IssuesSearchResult struct {
	Items []*Item
}

type JSONError struct {
	Message string
}

type Item struct {
	Number    int
	HTMLURL   string `json:"html_url`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url`
}

type ComplexIssuesByTime struct {
	BeforeMonth []Item
	BeforeYear  []Item
	AfterYear   []Item
}

type HTMLOutput struct {
	descriptTempl *template.Template
	tableTempl    *template.Template
	out           io.Writer
}

var issues *IssuesSearchResult

func init() {
	issues = formatJSONAnswer(doGetHTTPRequest(getURLRequest(os.Args[1:])))
	fmt.Print()
}

func newHTMLOutputForIssues(o io.Writer) *HTMLOutput {
	var result HTMLOutput = HTMLOutput{out: o}
	result.descriptTempl = template.Must(template.New("descriptionOfIssuesList").Parse(tableDescription))
	result.tableTempl = template.Must(template.New("issuesList").Parse(dataTemplHTML))
	return &result
}

const APIURL = "https://api.github.com/search/issues"

func formatJSONAnswer(resp *http.Response) *IssuesSearchResult {
	catchErrorRequest(resp)
	defer resp.Body.Close()

	return decodingJSONAnswer(resp)

}

func decodingJSONAnswer(resp *http.Response) *IssuesSearchResult {
	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatalln("decodingJSON:", err)
	}

	return &result
}

func catchErrorRequest(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		errorMessage := decodingErrorAnswer(resp)
		log.Fatalln("Request Error:", errorMessage.Message)
	}

}

func decodingErrorAnswer(resp *http.Response) JSONError {
	var errorMessage JSONError
	if err := json.NewDecoder(resp.Body).Decode(&errorMessage); err != nil {
		resp.Body.Close()
		log.Fatalln("JSON Decoding:", err)
	}
	return errorMessage
}

func doGetHTTPRequest(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("doHttpRequest:", err)
	}

	return resp
}

func getURLRequest(params []string) string {
	q := url.QueryEscape(strings.Join(params, " "))
	return APIURL + "?q=" + q
}

func sortingByName(issueses []*Item) *ComplexIssuesByTime {
	var result ComplexIssuesByTime
	for _, isuses := range issueses {
		if isuses.CreatedAt.Month() == time.Now().Month() {
			result.BeforeMonth = append(result.BeforeMonth, *isuses)
		} else if isuses.CreatedAt.Year() == time.Now().Year() {
			result.BeforeYear = append(result.BeforeYear, *isuses)
		} else {
			result.AfterYear = append(result.AfterYear, *isuses)
		}
	}
	return &result
}

func OutputIssuesInHtml(out io.Writer, issues *ComplexIssuesByTime) {
	templs := newHTMLOutputForIssues(out)
	outputMounthIssuesInHTML(templs, issues)
	outputYearIssuesInHTML(templs, issues)
	outputAfterYearIssuesInHtml(templs, issues)
}

func outputByIssusesByTime(issues *ComplexIssuesByTime) {
	outputMounthIssues(issues)
	outputYearIssues(issues)
	outputAfterYearIssues(issues)
}

func outputMounthIssuesInHTML(templs *HTMLOutput, issues *ComplexIssuesByTime) {
	if err := templs.descriptTempl.Execute(templs.out, "Before Month"); err != nil {
		log.Fatalln(err)
	}
	if err := templs.tableTempl.Execute(templs.out, issues.BeforeMonth); err != nil {
		log.Fatalln(err)
	}
}

func outputYearIssuesInHTML(templs *HTMLOutput, issues *ComplexIssuesByTime) {
	if err := templs.descriptTempl.Execute(templs.out, "Before Year"); err != nil {
		log.Fatalln(err)
	}
	if err := templs.tableTempl.Execute(templs.out, issues.BeforeYear); err != nil {
		log.Fatalln(err)
	}
}

func outputAfterYearIssuesInHtml(templs *HTMLOutput, issues *ComplexIssuesByTime) {
	if err := templs.descriptTempl.Execute(templs.out, "After Year"); err != nil {
		log.Fatalln(err)
	}
	if err := templs.tableTempl.Execute(templs.out, issues.AfterYear); err != nil {
		log.Fatalln(err)
	}
}

func outputMounthIssues(issues *ComplexIssuesByTime) {
	fmt.Println("Issues in this month: ")
	if err := outputIssues(issues.BeforeMonth); err != nil {
		log.Fatalln("month issues:", err)
	}
}

func outputYearIssues(issues *ComplexIssuesByTime) {
	fmt.Println("Issues in this year: ")
	if err := outputIssues(issues.BeforeYear); err != nil {
		log.Fatalln("month issues:", err)
	}
}

func outputAfterYearIssues(issues *ComplexIssuesByTime) {
	fmt.Println("Issues more year: ")
	if err := outputIssues(issues.AfterYear); err != nil {
		log.Fatalln("month issues:", err)
	}
}

func outputIssues(issuses []Item) error {
	if issuses == nil {
		return errors.New("null issues")
	}

	for _, item := range issuses {
		if item.User == nil {
			return errors.New("nill Users")
		}
		fmt.Printf("\t#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	return nil
}

func ServerIssues(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	OutputIssuesInHtml(w, sortingByName(issues.Items))
}

func main() {

	isServer := flag.Bool("s", false, "Use like server")
	flag.Parse()

	if *isServer {
		http.HandleFunc("/", ServerIssues)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	} else {
		outputByIssusesByTime(sortingByName(issues.Items))
	}
}
