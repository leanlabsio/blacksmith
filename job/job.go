package job

//Job represents single CI job to execute
type Job struct {
	Commit     string
	Ref        string
	Repository Repository
}

//Repository represents CI job repository to act on
type Repository struct {
	Name string
	URL  string
}
