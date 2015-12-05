package webhook

import (
	"encoding/json"
	"github.com/vasiliy-t/blacksmith/github"
	"github.com/vasiliy-t/blacksmith/gitlab"
	"gopkg.in/macaron.v1"
	"log"
)

//Resolve resolves webhook initiator, GitLab and GitHub is supported for now
func Resolve() macaron.Handler {
	return func(ctx *macaron.Context) {
		defer ctx.Req.Request.Body.Close()
		h := ctx.Req.Header.Get("X-Github-Event")
		if len(h) > 0 {
			var message github.Push
			err := json.NewDecoder(ctx.Req.Request.Body).Decode(&message)
			if err != nil {
				ctx.Map(err)
			}
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
			ctx.Map(message.MapToJob())
			return
		}

		ctx.Resp.WriteHeader(400)
		ctx.Resp.Write([]byte("Provider support is not implemented"))
	}
}
