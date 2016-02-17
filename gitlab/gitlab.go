package gitlab

import (
	"github.com/vasiliy-t/blacksmith/job"
)

//WebHook is a basic struct representing any gitlab webhook payload
type WebHook struct {
	ObjectKind string `json:"object_kind"`
}

//Repository represents gitlab repo info from webhook payload
type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
}

//Commit represents gitlab commit info from webhook payload
type Commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    User   `json:"author"`
}

//User represents gitlab user info from webhook payload
type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

//Push represents gitlab push notification payload
type Push struct {
	WebHook
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	UserID            int        `json:"user_id"`
	UserName          string     `json:"user_name"`
	UserEmail         string     `json:"user_email"`
	UserAvatar        string     `json:"user_avatar"`
	ProjectID         int        `json:"project_id"`
	Commits           []Commit   `json:"commits"`
	Repository        Repository `json:"repository"`
	TotalCommitsCount int        `json:"total_commits_count"`
	Added             []string   `json:"added"`
	Modified          []string   `json:"modified"`
	Removed           []string   `json:"removed"`
}

//TagPush represents gitlab tag push notification payload
type TagPush struct {
	WebHook
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	UserID            int        `json:"user_id"`
	UserName          string     `json:"user_name"`
	ProjectID         int        `json:"project_id"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
	TotalCommitsCount int        `json:"total_commit_count"`
}

//MapToJob maps webhook payload to executable job
func (p *Push) MapToJob() *job.Job {
	j := &job.Job{
		Commit: p.After,
		Ref:    p.Ref,
		Repository: job.Repository{
			Name: p.Repository.Name,
			URL:  p.Repository.GitHTTPURL,
		},
	}
	return j
}
