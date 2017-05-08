package api

import (
	"github.com/go-macaron/binding"
	"github.com/leanlabsio/blacksmith/user"
	"github.com/leanlabsio/blacksmith/middleware"
	"github.com/leanlabsio/blacksmith/project"
	"github.com/leanlabsio/blacksmith/repo"
	"golang.org/x/oauth2"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
)

// PutProject is an API endpoint to store project configuration
func PutProject(hostname string) []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		binding.Json(project.Project{}),
		func(ctx *macaron.Context, j project.Project, redis *redis.Client, user *user.User) {
			token := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: user.AccessToken},
			)

			hosting := repo.NewGithub(token)
			hosting.SetSelfHost(hostname)
			repository := project.New(hosting, redis)

			project := repository.Get(ctx.Params(":namespace"), ctx.Params(":name"))
			project.Executor = j.Executor

			if project.Trigger.Active != j.Trigger.Active {
				project.ToggleTrigger()
			}

			repository.Save(project)

			ctx.JSON(200, j)
		},
	}
}

// ListProject is an API endpoint to get all projects
func ListProject() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *user.User, redis *redis.Client) {
			enabledOnly := ctx.QueryInt("enabled")
			token := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: user.AccessToken},
			)
			hosting := repo.NewGithub(token)
			repository := project.New(hosting, redis)

			repos := repository.List(enabledOnly != 0)

			ctx.JSON(200, repos)
		},
	}
}

// GetProject is an API endpoint to get single project
func GetProject() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *user.User, redis *redis.Client) {
			token := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: user.AccessToken},
			)
			hosting := repo.NewGithub(token)
			repository := project.New(hosting, redis)

			project := repository.Get(ctx.Params(":namespace"), ctx.Params(":name"))

			ctx.JSON(200, project)
		},
	}
}
