package api

import (
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"log"
)

type Auth struct {
	ClientID string `json:"clientId"`
	Code     string `json:"code"`
	RedirectUri string `json:"redirectUri"`
}

func PostGitHubAuth(ghcid, ghcs string) []macaron.Handler {
	return []macaron.Handler{
		binding.Json(Auth{}),
		func(payload Auth) string {
			conf := &oauth2.Config{
				ClientID: ghcid,
				ClientSecret: ghcs,
				Endpoint: github.Endpoint,
			}

			token, err := conf.Exchange(oauth2.NoContext, payload.Code)
			if err != nil {
				log.Printf("AUTH error %s", err)
			}

			log.Printf("TOKEN: %+v", token)

			return "success"
		},
	}
}
