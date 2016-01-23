package api

import (
	"github.com/vasiliy-t/blacksmith/middleware"
	"github.com/vasiliy-t/blacksmith/model"
	"gopkg.in/macaron.v1"
	"log"
)

func ListRepo() []macaron.Handler {
	return []macaron.Handler{
		middleware.Auth(),
		func(user *model.User) {
			log.Printf("USER %+v", user)
		},
	}
}
