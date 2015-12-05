package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
	"github.com/vasiliy-t/blacksmith/job"
	"github.com/vasiliy-t/blacksmith/route"
	"github.com/vasiliy-t/blacksmith/webhook"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
	"net/http"
)

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
		cli.BoolFlag{
			Name:   "docker-tls-verify",
			EnvVar: "DOCKER_TLS_VERIFY",
		},
		cli.StringFlag{
			Name:   "redis-addr",
			EnvVar: "REDIS_ADDR",
		},
	},
	Action: daemon,
}

func daemon(c *cli.Context) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.String("redis-addr"),
		Password: "",
		DB:       0,
	})

	var client *docker.Client

	if c.Bool("docker-tls-verify") {
		certPath := c.String("docker-cert-path")
		client, _ = docker.NewTLSClient(
			c.String("docker-host"),
			fmt.Sprintf("%s/cert.pem", certPath),
			fmt.Sprintf("%s/key.pem", certPath),
			fmt.Sprintf("%s/ca.pem", certPath),
		)
	} else {
		client, _ = docker.NewClient(c.String("docker-host"))
	}

	m := macaron.New()
	m.Use(macaron.Recovery())
	m.Use(macaron.Logger())
	m.Map(redisClient)
	m.Map(client)
	m.Post("/push", webhook.Resolve(), func(job *job.Job, client *docker.Client) string {
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
	})

	m.Post("/env", route.PostEnv()...)

	err := http.ListenAndServe("0.0.0.0:9000", m)

	if err != nil {
		log.Fatalf("Failed to start %s", err)
	}
}
