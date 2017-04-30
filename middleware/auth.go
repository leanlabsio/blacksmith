package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/leanlabsio/blacksmith/auth"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
)

type Claims struct {
	Name string
	*jwt.StandardClaims
}

func Auth() macaron.Handler {
	return func(ctx *macaron.Context, redis *redis.Client) {
		header := ctx.Req.Header.Get("Authorization")
		if len(header) == 0 {
		}
		t := header[7:]

		token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("qwerty"), nil
		})

		if err != nil {

		}
		if !token.Valid {
		}
		claims := token.Claims.(*Claims)

		data, err := redis.HGetAllMap(claims.Name).Result()

		if err != nil {
		}
		if len(data) == 0 {
		}

		user := &auth.User{
			ID:          data["id"],
			Name:        data["name"],
			AccessToken: data["access_token"],
			AvatarURL:   data["avatar_url"],
			Login:       data["login"],
		}

		ctx.Map(user)
	}
}
