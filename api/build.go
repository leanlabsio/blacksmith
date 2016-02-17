package api

import (
	"fmt"
	"github.com/vasiliy-t/blacksmith/middleware"
	"github.com/vasiliy-t/blacksmith/model"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
)

type Build struct {
	Commit   string `json:"commit"`
	UserName string `json:"username"`
	Log      string `json:"log,omitempty"`
}

func ListBuild() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(ctx *macaron.Context, user *model.User, r *redis.Client) {
			repo := ctx.Params("*")
			log.Printf("REPO %s", repo)
			data, _ := r.ZRevRangeByScoreWithScores(repo+":builds", redis.ZRangeByScore{Min: "-inf", Max: "+inf", Offset: 0, Count: 10}).Result()
			var builds []Build
			for _, item := range data {
				build, _ := r.HGetAllMap(item.Member.(string)).Result()
				builds = append(builds, Build{UserName: build["user_name"], Commit: build["commit"]})
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
			ctx.JSON(200, Build{Log: data, UserName: build["user_name"], Commit: build["commit"]})
		},
	}
}
