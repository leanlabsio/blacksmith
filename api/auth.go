package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-macaron/binding"
	ghapi "github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
	"strconv"
	"time"
)

type Auth struct {
	ClientID    string `json:"clientId"`
	Code        string `json:"code"`
	RedirectUri string `json:"redirectUri"`
}

func PostGitHubAuth(ghcid, ghcs string) []macaron.Handler {
	return []macaron.Handler{
		binding.Json(Auth{}),
		func(ctx *macaron.Context, redis *redis.Client, payload Auth) {
			conf := &oauth2.Config{
				ClientID:     ghcid,
				ClientSecret: ghcs,
				Endpoint:     github.Endpoint,
			}

			token, err := conf.Exchange(oauth2.NoContext, payload.Code)

			if err != nil {
				log.Printf("AUTH error %s", err)
			}

			ghauth := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token.AccessToken},
			)

			tc := oauth2.NewClient(oauth2.NoContext, ghauth)
			client := ghapi.NewClient(tc)
			user, _, err := client.Users.Get("")

			if err != nil {

			}

			uid := strconv.Itoa(*user.ID)
			id := fmt.Sprintf("github:user:%s", uid)

			_, err = redis.HMSet(id, "id", uid, "name", *user.Name, "avatar", *user.AvatarURL, "access_token", token.AccessToken, "login", *user.Login).Result()

			if err != nil {
			}

			jwtToken := jwt.New(jwt.SigningMethodHS256)
			jwtToken.Claims["name"] = id
			jwtToken.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
			signed, _ := jwtToken.SignedString([]byte("qwerty"))

			type resp struct {
				Token string `json:"token"`
			}

			ctx.JSON(200, &resp{Token: signed})
		},
	}
}
