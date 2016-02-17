package webhook

import (
	"encoding/json"
	"fmt"
	"github.com/vasiliy-t/blacksmith/github"
	"github.com/vasiliy-t/blacksmith/gitlab"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
	"time"
)

//Resolve resolves webhook initiator, GitLab and GitHub is supported for now
func Resolve() macaron.Handler {
	return func(ctx *macaron.Context, r *redis.Client) {
		defer ctx.Req.Request.Body.Close()
		h := ctx.Req.Header.Get("X-GitHub-Event")
		if len(h) > 0 {
			var message github.Push
			err := json.NewDecoder(ctx.Req.Request.Body).Decode(&message)
			if err != nil {
				ctx.Map(err)
			}
			log.Printf("%s:%s", message.Repository.CloneURL, message.Ref)
			ctx.Map(message.MapToJob())
			return
		}

		h = ctx.Req.Header.Get("X-Gitlab-Event")
		if len(h) > 0 {
			var message gitlab.Push
			err := json.NewDecoder(ctx.Req.Request.Body).Decode(&message)
			if err != nil {
				ctx.Map(err)
				log.Fatalf("%+v", err)
			}
			key := fmt.Sprintf("%s:%s", message.Repository.GitHTTPURL, message.After)
			r.HMSet(key, "user_name", message.UserName, "commit", message.After).Result()

			r.ZAdd(message.Repository.GitHTTPURL+":builds", redis.Z{Score: float64(time.Now().Unix()), Member: key}).Result()
			ctx.Map(message.MapToJob())
			return
		}

		ctx.Resp.WriteHeader(400)
		ctx.Resp.Write([]byte("Provider support is not implemented"))
	}
}
