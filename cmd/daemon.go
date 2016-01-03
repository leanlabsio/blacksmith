package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
	"github.com/vasiliy-t/blacksmith/route"
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
	m.Post("/push", route.PostPush()...)
	m.Post("/env", route.PostEnv()...)
	m.Post("/job", route.PostJob()...)

	m.Get("/*", func(ctx *macaron.Context){
		ctx.HTML(200, "index")
	})

	err := http.ListenAndServe("0.0.0.0:80", m)

	if err != nil {
		log.Fatalf("Failed to start %s", err)
	}
}
