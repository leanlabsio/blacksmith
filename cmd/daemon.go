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
			Image: "leanlabs/make-builder",
			Cmd:   []string{"make"},
		}

		options := docker.CreateContainerOptions{
			Config:     config,
			HostConfig: nil,
		}

		container, err := client.CreateContainer(options)

		if err != nil {
			log.Fatalf("Docker error: %s", err)
		}

		err = client.StartContainer(container.ID, nil)

		if err != nil {
			log.Fatalf("Docker error: %s", err)
		}

		log.Printf("GITLAB PAYLOAD %+v", gpr.Repository.GitHttpUrl)
		return "QWERTY"
	})
	err := http.ListenAndServe("0.0.0.0:9000", m)

	if err != nil {
		log.Fatalf("Failed to start %s", err)
	}
}
