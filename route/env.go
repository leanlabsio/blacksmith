package route

import (
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
)

//Env represents single API payload entry
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//PostEnv is an API endpoint to store user defined environment
//variables in database
//POST Content-Type: application/json [{"ENVNAME:ENVVAL"}]
func PostEnv() []macaron.Handler {
	return []macaron.Handler{
		binding.Json([]Env{}),
		func(payload []Env, redis *redis.Client) string {
			for _, v := range payload {
				status := redis.Set(v.Name, v.Value, 0)
				log.Printf("%+v", status.Err())
			}
			return "Sucess"
		},
	}
}
