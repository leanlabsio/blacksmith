package api

import (
	"encoding/json"
	"log"

	"github.com/fsouza/go-dockerclient"
	"github.com/google/go-github/github"
	"github.com/leanlabsio/blacksmith/runner"
	"github.com/leanlabsio/blacksmith/logger"
	"github.com/leanlabsio/blacksmith/project"
	"github.com/leanlabsio/blacksmith/repo"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
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

			task := runner.Task{
				Name: pr.Name(),
				Builder: runner.Builder{
					Name: pr.Executor.Image.Name,
					Tag:  pr.Executor.Image.Tag,
				},
				Vars: runner.VarCollection{
					0: runner.Var{
						Name:  "EVENT_PAYLOAD",
						Value: data,
					},
				},
			}

			for _, v := range pr.Executor.EnvVars {
				task.Vars = append(task.Vars, runner.Var{Name: v.Name, Value: v.Value})
			}

			var event github.PushEvent
			json.Unmarshal([]byte(data), &event)

			logEntry := l.NewEntry(event, eventType, pr.Name())
			l.CreateEntry(logEntry)

			logEntry.Start()
			e := runner.New(client, logEntry)
			go e.Execute(task)

			log.Printf("%+v", pr)
			return "ok"
		},
	}
}
