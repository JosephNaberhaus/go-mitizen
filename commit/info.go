package commit

type info struct {
	CommitType string
	Scope string
	subject string
	body []string
	breakingChanges string
	issueReference string
}
