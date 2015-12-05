package route

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/vasiliy-t/blacksmith/job"
	"github.com/vasiliy-t/blacksmith/webhook"
	"gopkg.in/macaron.v1"
	"log"
)

//PostPush is an API endpoint to trigger builds
//POST Content-Type: application/json {GitLab/GitHub json webhook payload}
func PostPush() []macaron.Handler {
	return []macaron.Handler{
		webhook.Resolve(),
		func(job *job.Job, client *docker.Client) string {
			log.Printf("%+v", job)
			gitURL := fmt.Sprintf("REPOSITORY_GIT_HTTP_URL=%s", job.Repository.URL)
			ref := fmt.Sprintf("REF=%s", job.Ref)
			commit := fmt.Sprintf("AFTER=%s", job.Commit)
			reponame := fmt.Sprintf("REPOSITORY_NAME=%s", job.Repository.Name)

			config := &docker.Config{
				Image: "leanlabs/blacksmith-docker-runner",
				Volumes: map[string]struct{}{
					"/home":                {},
					"/var/run/docker.sock": {},
				},
				WorkingDir: "/home",
				Env: []string{
					gitURL,
					commit,
					reponame,
					ref,
				},
			}

			hostConfig := &docker.HostConfig{
				Binds: []string{
					"/home:/home",
					"/var/run/docker.sock:/var/run/docker.sock",
				},
			}

			options := docker.CreateContainerOptions{
				Config:     config,
				HostConfig: hostConfig,
			}

			container, err := client.CreateContainer(options)

			if err == docker.ErrNoSuchImage {
				client.PullImage(
					docker.PullImageOptions{
						Repository: "leanlabs/blacksmith-docker-runner",
						Tag:        "latest",
					},
					docker.AuthConfiguration{},
				)

				container, err = client.CreateContainer(options)

				if err != nil {
					log.Fatalf("Docker error: %s", err)
				}
			}

			err = client.StartContainer(container.ID, nil)

			if err != nil {
				log.Fatalf("Docker error: %s", err)
			}

			log.Printf("Job executed %+v", job)

			return "QWERTY"
		},
	}
}
