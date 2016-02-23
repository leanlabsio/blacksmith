package api

import (
	"encoding/json"
	"github.com/go-macaron/binding"
	"github.com/google/go-github/github"
	"github.com/leanlabsio/blacksmith/middleware"
	"github.com/leanlabsio/blacksmith/model"
	"golang.org/x/oauth2"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
	"strings"
)

//PostJob is an API endpoint to store jobs configuration
func PutJob() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		binding.Json(model.Job{}),
		func(ctx *macaron.Context, j model.Job, redis *redis.Client, user *model.User) {
			v, _ := json.Marshal(j)
			_, err := redis.Set(j.Repository, v, 0).Result()

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

			if strings.Contains(j.FullName, user.Login) {
				_, _, err = client.Repositories.CreateHook(user.Login, j.Name, &hook)
			} else {
				org := strings.Split(j.FullName, "/")[0]
				_, _, err = client.Repositories.CreateHook(org, j.Name, &hook)
			}

			if err != nil {
				log.Printf("ERROR %s", err)
			}

			ctx.JSON(200, j)
		},
	}
}

func ListJob() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *model.User, redis *redis.Client) {
			token := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: user.AccessToken},
			)

			tc := oauth2.NewClient(oauth2.NoContext, token)
			client := github.NewClient(tc)
			opts := &github.RepositoryListOptions{
				Type:        "all",
				ListOptions: github.ListOptions{PerPage: 100, Page: 1},
			}
			repos, _, _ := client.Repositories.List("", opts)

			var resp []*model.Job
			for _, repo := range repos {
				record, _ := redis.Get(*repo.CloneURL).Result()
				if len(record) != 0 {
					var j *model.Job
					json.Unmarshal([]byte(record), &j)
					resp = append(resp, j)
				} else {
					j := &model.Job{Repository: *repo.CloneURL, Name: *repo.Name, FullName: *repo.FullName, Enabled: false}
					resp = append(resp, j)
				}
			}

			ctx.JSON(200, resp)
		},
	}
}

func GetJob() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *model.User, redis *redis.Client) {
			data, err := redis.Get(ctx.Params("*")).Result()
			if err != nil {
			}
			var j *model.Job
			json.Unmarshal([]byte(data), &j)

			ctx.JSON(200, j)
		},
	}
}
