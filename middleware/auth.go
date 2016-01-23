package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/vasiliy-t/blacksmith/model"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
)

func Auth() macaron.Handler {
	return func(ctx *macaron.Context, redis *redis.Client) {
		header := ctx.Req.Header.Get("Authorization")
		if len(header) == 0 {
		}
		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte("qwerty"), nil
		})

		if err != nil {
		}
		if !token.Valid {
		}

		name, _ := token.Claims["name"]

		data, err := redis.HGetAllMap(name.(string)).Result()

		if err != nil {
		}
		if len(data) == 0 {
		}

		user := &model.User{
			ID:          data["id"],
			Name:        data["name"],
			AccessToken: data["access_token"],
			AvatarURL:   data["avatar_url"],
		}

		ctx.Map(user)
	}
}
