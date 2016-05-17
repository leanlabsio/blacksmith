package executor

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"io"
	"log"
)

// DockerExecutor represents docker task executor
type DockerExecutor struct {
	docker *docker.Client
	logger io.WriteCloser
}

func New(d *docker.Client, l io.WriteCloser) *DockerExecutor {
	return &DockerExecutor{
		docker: d,
		logger: l,
	}
}

func (e *DockerExecutor) Execute(t Task) {
	config := &docker.Config{
		Image: t.Builder.Name + ":" + t.Builder.Tag,
		Volumes: map[string]struct{}{
			"/var/run/docker.sock": {},
		},
		Env: t.Vars.String(),
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

	container, err := e.docker.CreateContainer(options)

	if err == docker.ErrNoSuchImage {
		err = e.docker.PullImage(
			docker.PullImageOptions{
				Repository: t.Builder.Name,
				Tag:        t.Builder.Tag,
			},
			docker.AuthConfiguration{},
		)

		if err != nil {
			msg := fmt.Sprintf("Docker error: %s", err)
			e.logger.Write([]byte(msg))
			e.logger.Close()
			return
		}

		container, err = e.docker.CreateContainer(options)

		if err != nil {
			msg := fmt.Sprintf("Docker error: %s", err)
			e.logger.Write([]byte(msg))
			e.logger.Close()
			return
		}
	}

	if err != nil {
		msg := fmt.Sprintf("Docker error: %s", err)
		e.logger.Write([]byte(msg))
		e.logger.Close()
		return
	}

	err = e.docker.StartContainer(container.ID, nil)
	e.docker.Logs(docker.LogsOptions{
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
	}
	e.logger.Close()

	log.Printf("Job executed %+v", e)
}
