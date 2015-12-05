package github

import (
	"github.com/vasiliy-t/blacksmith/job"
)

//Push represents github push webhook payload
type Push struct {
	Ref          string     `json:"ref"`
	After        string     `json:"after"`
	Before       string     `json:"before"`
	Size         int64      `json:"size"`
	DistinctSize int64      `json:"distinct_size"`
	Commits      []Commit   `json:"commits"`
	Repository   Repository `json:"repository"`
}

//Commit represents github webhook payload commit info
type Commit struct {
	Sha       string `json:"sha"`
	Message   string `json:"message"`
	Author    User   `json:"author"`
	URL       string `json:"url"`
	Distrinct bool   `json:"distinct"`
}

//User represents github webhook payload user info
type User struct{}

//Repository represents github webhook payload repo info
type Repository struct {
	Name     string `json:"name"`
	CloneURL string `json:"clone_url"`
}

//MapToJob maps webhook payload to executable job
func (p *Push) MapToJob() *job.Job {
	j := &job.Job{
		Commit: p.After,
		Ref:    p.Ref,
		Repository: job.Repository{
			Name: p.Repository.Name,
			URL:  p.Repository.CloneURL,
		},
	}
	return j
}
