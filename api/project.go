package api

import (
	"encoding/json"
	"github.com/go-macaron/binding"
	"github.com/google/go-github/github"
	"github.com/leanlabsio/blacksmith/middleware"
	"github.com/leanlabsio/blacksmith/model"
	"github.com/leanlabsio/blacksmith/project"
	"github.com/leanlabsio/blacksmith/repo"
	"golang.org/x/oauth2"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
	"strings"
)

//PostJob is an API endpoint to store jobs configuration
func PutProject() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		binding.Json(project.Project{}),
		func(ctx *macaron.Context, j project.Project, redis *redis.Client, user *model.User) {
			v, _ := json.Marshal(j)
			_, err := redis.Set(j.Repository.CloneURL, v, 0).Result()

			token := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: user.AccessToken},
			)

			tc := oauth2.NewClient(oauth2.NoContext, token)
			client := github.NewClient(tc)
			hook := github.Hook{
				Name:   github.String("web"),
				Active: github.Bool(true),
				Events: []string{"push", "pull_request"},
				Config: map[string]interface{}{
					"url":          "http://" + ctx.Req.Host + "/push",
					"content_type": "json",
				},
			}

			if strings.Contains(j.Repository.FullName, user.Login) {
				_, _, err = client.Repositories.CreateHook(user.Login, j.Repository.Name, &hook)
			} else {
				org := strings.Split(j.Repository.FullName, "/")[0]
				_, _, err = client.Repositories.CreateHook(org, j.Repository.Name, &hook)
			}

			if err != nil {
				log.Printf("ERROR %s", err)
			}

			ctx.JSON(200, j)
		},
	}
}

func ListProject() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *model.User, redis *redis.Client) {

			token := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: user.AccessToken},
			)
			hosting := repo.NewGithub(token)
			repository := project.New(&hosting, redis)

			repos := repository.List()

			ctx.JSON(200, repos)
		},
	}
}

func GetProject() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *model.User, redis *redis.Client) {
			data, err := redis.Get(ctx.Params("*")).Result()
			if err != nil {
				log.Printf("REDIS ERR %s", err)
			}
			var j *project.Project
			json.Unmarshal([]byte(data), &j)

			ctx.JSON(200, j)
		},
	}
}
