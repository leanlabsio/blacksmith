package executor

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/leanlabsio/blacksmith/logger"
	"log"
)

// DockerExecutor represents docker task executor
type DockerExecutor struct {
	Image   Image `json:"image"`
	EnvVars []Env `json:"env"`
	client  *docker.Client
	logger  *logger.Writer
}

// Env represents any additional confugration parameters
// to be passed to Builder, in key - value format
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Image represents actual docker image to be used
// for build
type Image struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func (e *DockerExecutor) SetClient(client *docker.Client) {
	e.client = client
}

func (e *DockerExecutor) SetLogger(logger *logger.Writer) {
	e.logger = logger
}

func (e *DockerExecutor) WithData(data string) {
	env := Env{
		Name:  "EVENT_PAYLOAD",
		Value: data,
	}
	e.EnvVars = append(e.EnvVars, env)
}

func (e *DockerExecutor) Execute() {
	var envs []string

	if (len(e.EnvVars)) > 0 {
		for _, e := range e.EnvVars {
			env := fmt.Sprintf("%s=%s", e.Name, e.Value)
			envs = append(envs, env)
		}
	}

	config := &docker.Config{
		Image: e.Image.Name + ":" + e.Image.Tag,
		Volumes: map[string]struct{}{
			"/var/run/docker.sock": {},
		},
		Env: envs,
	}

	hostConfig := &docker.HostConfig{
		Binds: []string{
			"/var/run/docker.sock:/var/run/docker.sock",
		},
	}

	options := docker.CreateContainerOptions{
		Config:     config,
		HostConfig: hostConfig,
	}

	container, err := e.client.CreateContainer(options)

	if err == docker.ErrNoSuchImage {
		e.client.PullImage(
			docker.PullImageOptions{
				Repository: e.Image.Name,
				Tag:        e.Image.Tag,
			},
			docker.AuthConfiguration{},
		)

		container, err = e.client.CreateContainer(options)

		if err != nil {
			log.Printf("Docker error: %s", err)
		}
	}

	err = e.client.StartContainer(container.ID, nil)
	/*	writer := &WsWriter{
		name:  fmt.Sprintf("%s:%s:log", job.Repository.URL, job.Commit),
		redis: r,
	}*/
	e.client.Logs(docker.LogsOptions{
		Container:    container.ID,
		OutputStream: e.logger,
		ErrorStream:  e.logger,
		Follow:       true,
		Stdout:       true,
		Stderr:       true,
	})

	if err != nil {
		msg := fmt.Sprintf("Docker error: %s", err)
		e.logger.Write([]byte(msg))
		log.Fatal(msg)
	}

	log.Printf("Job executed %+v", e)
}
