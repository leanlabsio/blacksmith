package job

import (
	"encoding/json"
	"fmt"
	"github.com/leanlabsio/blacksmith/model"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
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
		var j *model.Job
		data, _ := redis.Get(job.Repository.URL).Result()

		json.Unmarshal([]byte(data), &j)

		if len(j.EnvVars) > 0 {
			for _, e := range j.EnvVars {
				env := fmt.Sprintf("%s=%s", e.Name, e.Value)
				job.EnvVars = append(job.EnvVars, env)
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
