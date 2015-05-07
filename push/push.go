package push

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PushHandler struct {
}

type GitlabWebHook struct {
	ObjectKind string `json:"object_kind"`
}

type Repository struct {
	Name            string `json:"name"`
	Url             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHttpUrl      string `json:"git_http_url"`
	GitSshUrl       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
}

type Commit struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
	Author    User   `json:"author"`
}

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

type GitlabPushRequest struct {
	GitlabWebHook
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	UserId            int        `json:"user_id"`
	UserName          string     `json:"user_name"`
	UserEmail         string     `json:"user_email"`
	ProjectId         int        `json:"project_id"`
	TotalCommitsCount int        `json:"total_commits_count"`
	Commits           []Commit   `json:"commits"`
	Repository        Repository `json:"repository"`
}

func (ph *PushHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var data GitlabPushRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Called handle on GitlabPushRequest %+v", data)
}
