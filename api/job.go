package api

import (
	"encoding/json"
	"github.com/go-macaron/binding"
	"github.com/google/go-github/github"
	"github.com/vasiliy-t/blacksmith/middleware"
	"github.com/vasiliy-t/blacksmith/model"
	"golang.org/x/oauth2"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
	"strconv"
)

//Job represents single API payload entry
type Job struct {
	Repository string `json:"repository"`
	EnvVars    []Env  `json:"env"`
	Enabled    bool   `json:"enabled"`
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//PostJob is an API endpoint to store jobs configuration
func PutJob() []macaron.Handler {
	return []macaron.Handler{
		binding.Json(Job{}),
		func(ctx *macaron.Context, j Job, redis *redis.Client) {
			log.Printf("JOB %+v", j)
			v, _ := json.Marshal(j)
			_, err := redis.Set(j.Repository, v, 0).Result()

			if err != nil {

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
				Type: "owner",
			}
			repos, _, _ := client.Repositories.List(user.Login, opts)

			var resp []*Job
			for _, repo := range repos {
				j := &Job{Repository: *repo.CloneURL, Enabled: false}
				job, _ := redis.HGetAllMap(*repo.CloneURL).Result()

				if len(job) != 0 {
					e, _ := strconv.ParseBool(job["enabled"])
					j.Enabled = e
				}
				resp = append(resp, j)
			}

			ctx.JSON(200, resp)
		},
	}
}

func GetJob() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *model.User, redis *redis.Client) {
			data, err := redis.HGetAllMap(ctx.Params("*")).Result()
			if err != nil {
			}

			e, _ := strconv.ParseBool(data["enabled"])

			job := &Job{
				Repository: ctx.Params("*"),
				Enabled:    e,
			}
			log.Printf("REPO %s", ctx.Params("*"))

			ctx.JSON(200, job)
		},
	}
}
