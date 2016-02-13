package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
	"github.com/leanlabsio/sockets"
	"github.com/vasiliy-t/blacksmith/api"
	"github.com/vasiliy-t/ws"
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
		cli.StringFlag{
			Name:   "token-sign",
			EnvVar: "TOKEN_SIGN",
		},
		cli.StringFlag{
			Name:   "github-client-id",
			EnvVar: "GITHUB_CLIENT_ID",
		},
		cli.StringFlag{
			Name:   "github-client-secret",
			EnvVar: "GITHUB_CLIENT_SECRET",
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

	var dockerClient *docker.Client

	if c.Bool("docker-tls-verify") {
		certPath := c.String("docker-cert-path")
		dockerClient, _ = docker.NewTLSClient(
			c.String("docker-host"),
			fmt.Sprintf("%s/cert.pem", certPath),
			fmt.Sprintf("%s/key.pem", certPath),
			fmt.Sprintf("%s/ca.pem", certPath),
		)
	} else {
		dockerClient, _ = docker.NewClient(c.String("docker-host"))
	}

	m := macaron.New()
	m.Use(macaron.Recovery())
	m.Use(macaron.Logger())
	m.Use(macaron.Renderer())
	m.Use(macaron.Static("web"))

	m.Map(redisClient)
	m.Map(dockerClient)
	m.Get("/repo", api.ListRepo()...)
	m.Post("/push", api.PostPush()...)
	m.Put("/jobs", api.PutJob()...)
	m.Get("/jobs", api.ListJob()...)
	m.Get("/jobs/*", api.GetJob()...)
	m.Post("/auth/github", api.PostGitHubAuth(c.String("github-client-id"), c.String("github-client-secret"))...)
	m.Get("/ws/*", sockets.Messages(), ws.ListenAndServe)

	m.Get("/*", func(ctx *macaron.Context) {
		ctx.Data["BSConfig"] = map[string]interface{}{
			"version": "1.0.0",
			"github": map[string]interface{}{
				"oauth": map[string]string{
					"clientid": c.String("github-client-id"),
				},
			},
		}
		ctx.HTML(200, "index")
	})

	err := http.ListenAndServe("0.0.0.0:80", m)

	if err != nil {
		log.Fatalf("Failed to start %s", err)
	}
}
