package api

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/vasiliy-t/blacksmith/job"
	"github.com/vasiliy-t/blacksmith/webhook"
	"github.com/vasiliy-t/ws"
	"gopkg.in/macaron.v1"
	"log"
)

type WsWriter struct {
	name string
}

func (w WsWriter) Write(p []byte) (int, error) {
	s := ws.Server(w.name)
	s.Broadcast(string(p))
	return len(p), nil
}

//PostPush is an API endpoint to trigger builds
//POST Content-Type: application/json {GitLab/GitHub json webhook payload}
func PostPush() []macaron.Handler {
	return []macaron.Handler{
		webhook.Resolve(),
		job.Resolve(),
		func(ctx *macaron.Context, job *job.Job, client *docker.Client) string {

			go func() {
				config := &docker.Config{
					Image: "leanlabs/blacksmith-docker-runner",
					Volumes: map[string]struct{}{
						"/home":                {},
						"/var/run/docker.sock": {},
					},
					WorkingDir: "/home",
					Env:        job.EnvVars,
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
				client.Logs(docker.LogsOptions{
					Container:    container.ID,
					OutputStream: &WsWriter{name: job.Repository.URL},
					ErrorStream:  &WsWriter{name: job.Repository.URL},
					Follow:       true,
					Stdout:       true,
					Stderr:       true,
				})

				if err != nil {
					log.Fatalf("Docker error: %s", err)
				}

				log.Printf("Job executed %+v", job)
			}()

			return "QWERTY"
		},
	}
}
