package job

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
)

//Job represents single CI job to execute
type Job struct {
	Commit     string
	Ref        string
	Repository Repository
	EnvVars    []string
}

//Repository represents CI job repository to act on
type Repository struct {
	Name string
	URL  string
}

//Resolve fills Job with additional ENV variables
//which should be passed to runner
func Resolve() macaron.Handler {
	return func(redis *redis.Client, job *Job) {
		varRefences, err := redis.SMembers(job.Repository.URL).Result()
		if err != nil {
			log.Printf("JOB error: %s", err)
		}

		if len(varRefences) > 0 {
			vars, err := redis.MGet(varRefences...).Result()

			if err != nil {
				log.Printf("JOB error: %s", err)
			}

			for _, v := range vars {
				if str, ok := v.(string); ok {
					job.EnvVars = append(job.EnvVars, str)
				}
			}
		}

		job.EnvVars = append(
			job.EnvVars,
			fmt.Sprintf("REPOSITORY_GIT_HTTP_URL=%s", job.Repository.URL),
			fmt.Sprintf("REF=%s", job.Ref),
			fmt.Sprintf("COMMIT=%s", job.Commit),
			fmt.Sprintf("REPOSITORY_NAME=%s", job.Repository.Name),
		)
	}
}
