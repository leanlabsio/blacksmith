package api

import (
	"encoding/json"
	"github.com/fsouza/go-dockerclient"
	"github.com/google/go-github/github"
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
		func(ctx *macaron.Context, rc *redis.Client, client *docker.Client, l *logger.Logger) string {
			data, _ := ctx.Req.Body().String()
			eventType := ctx.Req.Header.Get("X-GitHub-Event")

			if eventType == "ping" {
				return "ok"
			}

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

			var event github.PushEvent
			json.Unmarshal([]byte(data), &event)

			logEntry := l.NewEntry(event, eventType, pr.Name())
			l.CreateEntry(logEntry)

			logEntry.Start()
			e := executor.New(client, logEntry)
			go e.Execute(task)

			log.Printf("%+v", pr)
			return "ok"
		},
	}
}
