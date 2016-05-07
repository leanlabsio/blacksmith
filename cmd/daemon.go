package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
	"github.com/go-macaron/bindata"
	"github.com/leanlabsio/blacksmith/api"
	"github.com/leanlabsio/blacksmith/templates"
	"github.com/leanlabsio/blacksmith/web"
	"github.com/leanlabsio/sockets"
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
	m.Use(macaron.Static("web", macaron.StaticOptions{
		FileSystem: bindata.Static(bindata.Options{
			Asset:      web.Asset,
			AssetDir:   web.AssetDir,
			AssetNames: web.AssetNames,
			AssetInfo:  web.AssetInfo,
			Prefix:     "web",
		}),
	}))

	m.Use(macaron.Renderer(
		macaron.RenderOptions{
			Directory: "templates",
			TemplateFileSystem: bindata.Templates(bindata.Options{
				Asset:      templates.Asset,
				AssetDir:   templates.AssetDir,
				AssetNames: templates.AssetNames,
				AssetInfo:  templates.AssetInfo,
				Prefix:     "",
			}),
		},
	))

	m.Map(redisClient)
	m.Map(dockerClient)

	m.Post("/api/trigger", api.PostTrigger()...)

	m.Put("/api/projects/:host/:namespace/:name", api.PutProject()...)
	m.Get("/api/projects", api.ListProject()...)
	m.Get("/api/projects/:host/:namespace/:name", api.GetProject()...)

	m.Get("/api/builds/*", api.ListBuild()...)
	m.Get("/api/logs/*", api.GetBuild()...)

	m.Post("/api/auth/github", api.PostGitHubAuth(c.String("github-client-id"), c.String("github-client-secret"))...)
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
		ctx.HTML(200, "templates/index")
	})

	err := http.ListenAndServe("0.0.0.0:80", m)

	if err != nil {
		log.Fatalf("Failed to start %s", err)
	}
}
