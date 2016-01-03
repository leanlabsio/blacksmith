package api

import (
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
)

//Job represents single API payload entry
type Job struct {
	Repository    string   `json:"repository"`
	EnvReferences []string `json:"env"`
}

//PostJob is an API endpoint to store jobs configuration
func PostJob() []macaron.Handler {
	return []macaron.Handler{
		binding.Json(Job{}),
		func(j Job, redis *redis.Client) {
			i, err := redis.SAdd(j.Repository, j.EnvReferences...).Result()
			log.Printf("POST Job: %+v", i)
			if err != nil {
				log.Printf("POST Job error: %s", err)
			}
		},
	}
}
