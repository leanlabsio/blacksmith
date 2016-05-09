package api

import (
	"fmt"
	"github.com/leanlabsio/blacksmith/logger"
	"github.com/leanlabsio/blacksmith/middleware"
	"github.com/leanlabsio/blacksmith/model"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
)

func ListBuild() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *model.User, r *redis.Client, l *logger.Logger) {
			host := ctx.Params(":host")
			namespace := ctx.Params(":namespace")
			name := ctx.Params(":name")

			entries := l.ListEntries(host, namespace, name)

			ctx.JSON(200, entries)
		},
	}
}

func GetBuild() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, r *redis.Client) {
			key := fmt.Sprintf("%s:%s", ctx.Params("*"), ctx.Query("commit"))
			logKey := fmt.Sprintf("%s:log", key)
			build, _ := r.HGetAllMap(key).Result()
			data, _ := r.Get(logKey).Result()
			ctx.JSON(200, model.Build{Log: data, UserName: build["user_name"], Commit: build["commit"]})
		},
	}
}
