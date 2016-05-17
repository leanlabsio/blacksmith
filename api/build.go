package api

import (
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
		func(ctx *macaron.Context, l *logger.Logger) {
			host := ctx.Params(":host")
			namespace := ctx.Params(":namespace")
			name := ctx.Params(":name")
			commit := ctx.Params(":commit")
			timestamp := ctx.Params(":timestamp")

			log := l.GetLog(host, namespace, name, commit, timestamp)

			ctx.JSON(200, log)
		},
	}
}
