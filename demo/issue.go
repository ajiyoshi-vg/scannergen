package demo

//go:generate ../scannergen -file issue.go -pkg demo -o issue.scanner.go

type Issue struct {
	ID       int64 `sql:"pk: true, auto: true"`
	Number   int
	Title    string   `sql:"size: 512"`
	Body     string   `sql:"size: 2048"`
	Assignee string   `sql:"index: issue_assignee"`
	State    string   `sql:"size: 50"`
	Labels   []string `sql:"encode: json"`

	locked bool `sql:"-"`
}
