package api

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/leanlabsio/blacksmith/executor"
	"github.com/leanlabsio/blacksmith/logger"
	"github.com/leanlabsio/blacksmith/project"
	"github.com/leanlabsio/blacksmith/repo"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
)

func PostTrigger() []macaron.Handler {
	return []macaron.Handler{
		func(ctx *macaron.Context, rc *redis.Client, client *docker.Client, l *logger.Logger) {
			data, _ := ctx.Req.Body().String()

			host := ctx.Query("host")
			namespace := ctx.Query("namespace")
			name := ctx.Query("name")
			hosting := repo.NewDummy(host)

			repository := project.New(hosting, rc)

			pr := repository.Get(namespace, name)

			task := executor.Task{
				Name: pr.Name(),
				Builder: executor.Builder{
					Name: pr.Executor.Image.Name,
					Tag:  pr.Executor.Image.Tag,
				},
				Vars: executor.VarCollection{
					0: executor.Var{
						Name:  "EVENT_PAYLOAD",
						Value: data,
					},
				},
			}

			for _, v := range pr.Executor.EnvVars {
				task.Vars = append(task.Vars, executor.Var{Name: v.Name, Value: v.Value})
			}

			logEntry := l.CreateEntry(pr.Name())
			e := executor.New(client, logEntry)
			go e.Execute(task)

			log.Printf("%+v", pr)
		},
	}
}
