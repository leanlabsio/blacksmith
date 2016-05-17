package project

import (
	"encoding/json"
	"fmt"
	"github.com/leanlabsio/blacksmith/repo"
	"github.com/leanlabsio/blacksmith/trigger"
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

func (r *ProjectRepository) List(enabledOnly bool) []Project {
	repos := r.hosting.ListRepositories()

	var ret []Project

	for _, repo := range repos {
		key := fmt.Sprintf("%s:%s:%s", r.hosting.Host(), repo.Owner.Name, repo.Name)
		record, _ := r.db.Get(key).Result()

		if len(record) != 0 {
			var j Project
			json.Unmarshal([]byte(record), &j)
			if j.Trigger.Active == false && enabledOnly == true {
				continue
			}
			ret = append(ret, j)
		} else {
			if enabledOnly == false {
				j := Project{
					Repository: repo,
					Trigger: trigger.Trigger{
						Active: false,
					},
				}
				ret = append(ret, j)
			}
		}
	}

	return ret
}

func (r *ProjectRepository) Get(namespace, name string) *Project {
	repo := r.hosting.GetRepository(namespace, name)

	key := fmt.Sprintf("%s:%s:%s", r.hosting.Host(), repo.Owner.Name, repo.Name)

	record, _ := r.db.Get(key).Result()
	var ret Project

	if len(record) != 0 {
		json.Unmarshal([]byte(record), &ret)
	} else {
		ret = Project{
			Repository: repo,
			Trigger: trigger.Trigger{
				Active: false,
			},
		}
	}

	ret.triggerRepo = trigger.NewRepository(r.hosting)
	ret.projectRepo = r

	return &ret
}

func (r *ProjectRepository) Save(p *Project) *Project {
	v, _ := json.Marshal(p)
	key := fmt.Sprintf("%s:%s:%s", r.hosting.Host(), p.Repository.Owner.Name, p.Repository.Name)
	r.db.Set(key, v, 0).Result()
	return p
}
