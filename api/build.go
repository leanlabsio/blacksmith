package api

import (
	"fmt"
	"github.com/leanlabsio/blacksmith/middleware"
	"github.com/leanlabsio/blacksmith/model"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
)

func ListBuild() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *model.User, r *redis.Client) {
			repo := ctx.Params("*")
			log.Printf("REPO %s", repo)
			data, _ := r.ZRevRangeByScoreWithScores(repo+":builds", redis.ZRangeByScore{Min: "-inf", Max: "+inf", Offset: 0, Count: 10}).Result()
			var builds []model.Build
			for _, item := range data {
				build, _ := r.HGetAllMap(item.Member.(string)).Result()
				builds = append(builds, model.Build{UserName: build["user_name"], Commit: build["commit"], Status: model.BUILD_STATUS_FAILED})
			}
			ctx.JSON(200, builds)
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
