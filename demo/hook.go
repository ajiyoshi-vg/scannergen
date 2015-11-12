package demo

//go:generate ../scannergen -file hook.go -pkg demo -o hook.scanner.go

type Hook struct {
	ID         int64 `sql:"pk: true, auto: true"`
	Sha        string
	After      string
	Before     string
	Created    bool
	Deleted    bool
	Forced     bool
	HeadCommit *Commit `sql:"name: head"`
}

type Commit struct {
	ID        string
	Message   string
	Timestamp string
	Author    *Author
	Committer *Author
}

type Author struct {
	Name     string
	Email    string
	Username string
}
