package job

import (
	"encoding/json"
	"fmt"
	"github.com/leanlabsio/blacksmith/model"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
	"log"
)

//Job represents singleCI job to execute
type Job struct {
	Builder    Builder
	Commit     string
	Ref        string
	Repository Repository
	EnvVars    []string
}

type Builder struct {
	Name string
	Tag  string
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
		j := model.Job{}
		data, err := redis.Get(job.Repository.URL).Result()
		if err != nil {
			log.Printf("REDIS ERROR %s", err)
		}

		json.Unmarshal([]byte(data), &j)

		if len(j.EnvVars) > 0 {
			for _, e := range j.EnvVars {
				env := fmt.Sprintf("%s=%s", e.Name, e.Value)
				job.EnvVars = append(job.EnvVars, env)
			}
		}
		job.Builder = Builder{
			Name: j.Builder.Name,
			Tag:  j.Builder.Tag,
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
