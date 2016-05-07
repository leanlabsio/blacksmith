package api

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/leanlabsio/blacksmith/logger"
	"github.com/leanlabsio/blacksmith/project"
	"github.com/leanlabsio/blacksmith/repo"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
)

func PostTrigger() []macaron.Handler {
	return []macaron.Handler{
		func(ctx *macaron.Context, rc *redis.Client, client *docker.Client) {
			data, _ := ctx.Req.Body().String()

			host := ctx.Query("host")
			namespace := ctx.Query("namespace")
			name := ctx.Query("name")
			hosting := repo.NewDummy(host)

			repository := project.New(hosting, rc)

			pr := repository.Get(namespace, name)
			lg := logger.New("qwerty", rc)

			pr.Executor.SetLogger(lg)
			pr.Executor.SetClient(client)
			pr.Executor.WithData(data)
			go pr.Executor.Execute()

			log.Printf("%+v", pr)
		},
	}
}
