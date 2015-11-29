package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
)

//GitlabWebHook is a basic struct representing any gitlab webhook payload
type GitlabWebHook struct {
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

//GitlabPushRequest represents gitlab push notification payload
type GitlabPushRequest struct {
	GitlabWebHook
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	UserID            int        `json:"user_id"`
	UserName          string     `json:"user_name"`
	UserEmail         string     `json:"user_email"`
	ProjectID         int        `json:"project_id"`
	TotalCommitsCount int        `json:"total_commits_count"`
	Commits           []Commit   `json:"commits"`
	Repository        Repository `json:"repository"`
}

//DaemonCmd is a command to start server in daemon mode
var DaemonCmd = cli.Command{
	Name: "daemon",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "docker-host",
			EnvVar: "DOCKER_HOST",
		},
		cli.StringFlag{
			Name:   "docker-cert-path",
			EnvVar: "DOCKER_CERT_PATH",
		},
	},
	Action: daemon,
}

func daemon(c *cli.Context) {
	m := macaron.New()
	m.Use(macaron.Recovery())
	m.Use(macaron.Logger())
	m.Post("/push", binding.Json(GitlabPushRequest{}), func(gpr GitlabPushRequest) string {
		client, err := docker.NewTLSClient(
			c.String("docker-host"),
			fmt.Sprintf("%s/cert.pem", c.String("docker-cert-path")),
			fmt.Sprintf("%s/key.pem", c.String("docker-cert-path")),
			fmt.Sprintf("%s/ca.pem", c.String("docker-cert-path")),
		)

		config := &docker.Config{
			Image: "leanlabs/bsr",
			Volumes: map[string]struct{}{
				"/home":                {},
				"/var/run/docker.sock": {},
				"/var/run/docker.pid":  {},
			},
			WorkingDir: "/home",
			Env: []string{
				"REPOSITORY_GIT_HTTP_URL=https://github.com/leanlabsio/kanban.git",
				"AFTER=c143783000e9e4f89d699e294cb5ecb05099c16b",
				"REPOSITORY_NAME=kanban",
			},
		}

		hostConfig := &docker.HostConfig{
			Binds: []string{
				"/home:/home",
				"/var/run/docker.sock:/var/run/docker.sock",
				"/var/run/docker.pid:/var/run/docker.pid",
			},
		}

		options := docker.CreateContainerOptions{
			Config:     config,
			HostConfig: hostConfig,
		}

		container, err := client.CreateContainer(options)

		if err != nil {
			log.Fatalf("Docker error: %s", err)
		}

		err = client.StartContainer(container.ID, nil)

		if err != nil {
			log.Fatalf("Docker error: %s", err)
		}

		log.Printf("GITLAB PAYLOAD %+v", gpr.Repository.GitHTTPURL)
		return "QWERTY"
	})
	err := http.ListenAndServe("0.0.0.0:9000", m)

	if err != nil {
		log.Fatalf("Failed to start %s", err)
	}
}
