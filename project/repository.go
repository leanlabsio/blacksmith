package project

import (
	"encoding/json"
	"github.com/leanlabsio/blacksmith/repo"
	"gopkg.in/redis.v3"
)

type ProjectRepository struct {
	hosting repo.Hosting
	db      *redis.Client
}

func FromContext() *ProjectRepository {
	return &ProjectRepository{}
}

func New(h repo.Hosting, db *redis.Client) *ProjectRepository {
	return &ProjectRepository{
		hosting: h,
		db:      db,
	}
}

func (r *ProjectRepository) List() []Project {
	repos := r.hosting.ListRepositories()

	var ret []Project

	for _, repo := range repos {
		record, _ := r.db.Get(repo.CloneURL).Result()
		if len(record) != 0 {
			var j Project
			json.Unmarshal([]byte(record), &j)
			ret = append(ret, j)
		} else {
			j := Project{
				Repository: repo,
				Enabled:    false,
			}
			ret = append(ret, j)
		}
	}

	return ret
}

func (r *ProjectRepository) Get() {

}
