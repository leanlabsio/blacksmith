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
)

//Job represents single API payload entry
type Job struct {
	Name       string `json:"name"`
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
				record, _ := redis.Get(*repo.CloneURL).Result()
				if len(record) != 0 {
					var j *Job
					json.Unmarshal([]byte(record), &j)
					resp = append(resp, j)
				} else {
					j := &Job{Repository: *repo.CloneURL, Name: *repo.FullName, Enabled: false}
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
			var j *Job
			json.Unmarshal([]byte(data), &j)

			ctx.JSON(200, j)
		},
	}
}
