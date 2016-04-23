package api

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/leanlabsio/blacksmith/job"
	"github.com/leanlabsio/blacksmith/webhook"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
)

type WsWriter struct {
	name  string
	redis *redis.Client
}

func (w WsWriter) Write(p []byte) (int, error) {
	w.redis.Append(w.name, string(p)).Result()
	return len(p), nil
}

//PostPush is an API endpoint to trigger builds
//POST Content-Type: application/json {GitLab/GitHub json webhook payload}
func PostPush() []macaron.Handler {
	return []macaron.Handler{
		webhook.Resolve(),
		job.Resolve(),
		func(ctx *macaron.Context, job *job.Job, client *docker.Client, r *redis.Client) string {

			go func() {
				config := &docker.Config{
					Image: job.Builder.Name + ":" + job.Builder.Tag,
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
							Repository: job.Builder.Name,
							Tag:        job.Builder.Tag,
						},
						docker.AuthConfiguration{},
					)

					container, err = client.CreateContainer(options)

					if err != nil {
						log.Fatalf("Docker error: %s", err)
					}
				}

				err = client.StartContainer(container.ID, nil)
				writer := &WsWriter{
					name:  fmt.Sprintf("%s:%s:log", job.Repository.URL, job.Commit),
					redis: r,
				}
				client.Logs(docker.LogsOptions{
					Container:    container.ID,
					OutputStream: writer,
					ErrorStream:  writer,
					Follow:       true,
					Stdout:       true,
					Stderr:       true,
				})

				if err != nil {
					msg := fmt.Sprintf("Docker error: %s", err)
					writer.Write([]byte(msg))
					log.Fatal(msg)
				}

				log.Printf("Job executed %+v", job)
			}()

			return "QWERTY"
		},
	}
}
