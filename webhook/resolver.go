package webhook

import (
	"gopkg.in/macaron.v1"
)

//Resolve resolves webhook initiator, GitLab and GitHub is supported for now
func Resolve() macaron.Handler {
	return func(ctx *macaron.Context) {
		h := ctx.Req.Header.Get("X-Github-Event")
		if len(h) > 0 {
			return
		}

		h = ctx.Req.Header.Get("X-Gitlab-Event")
		if len(h) > 0 {
			return
		}

		ctx.Resp.WriteHeader(400)
		ctx.Resp.Write([]byte("Provider support is not implemented"))
	}
}
