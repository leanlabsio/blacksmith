package api

import (
	"github.com/google/go-github/github"
	"github.com/leanlabsio/blacksmith/middleware"
	"github.com/leanlabsio/blacksmith/model"
	"golang.org/x/oauth2"
	"gopkg.in/macaron.v1"
)

func ListRepo() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *model.User) {
			token := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: user.AccessToken},
			)

			tc := oauth2.NewClient(oauth2.NoContext, token)
			client := github.NewClient(tc)
			opts := &github.RepositoryListOptions{
				Type: "owner",
			}
			repos, _, _ := client.Repositories.List(user.Login, opts)

			ctx.JSON(200, repos)
		},
	}
}
