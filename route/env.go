package route

import (
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"log"
)

type Env struct {
	EnvName string `json:"envname"`
	EnvVal  string `json:"envval"`
}

//PostEnv is an API endpoint to store user defined environment
//variables in database
//POST Content-Type: application/json [{"ENVNAME:ENVVAL"}]
func PostEnv() []macaron.Handler {
	return []macaron.Handler{
		binding.Json(Env{}),
		func(e Env) string {
			log.Printf("%+v", e)
			return "Sucess"
		},
	}
}
