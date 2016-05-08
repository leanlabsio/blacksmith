package project

import (
	"fmt"
	"github.com/leanlabsio/blacksmith/repo"
	"github.com/leanlabsio/blacksmith/trigger"
)

type Executor struct {
	Image   Image `json:"image"`
	EnvVars []Env `json:"env"`
}

// Env represents any additional confugration parameters
// to be passed to Builder, in key - value format
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Image represents actual docker image to be used
// for build
type Image struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

// Project represents single project to build
type Project struct {
	Repository repo.Repository `json:"repository"`
	Trigger    trigger.Trigger `json:"trigger"`
	Executor   Executor        `json:"executor"`

	triggerRepo *trigger.TriggerRepository
	projectRepo *ProjectRepository
}

func (p *Project) ToggleTrigger() {
	if !p.Trigger.Active {
		p.triggerRepo.CreateTrigger(p.Repository)
		p.Trigger.Active = true
	} else {
		p.Trigger.Active = false
		//p.triggerRepo.RemoveTrigger(p.Repository)
	}
}

func (p *Project) Name() string {
	return fmt.Sprintf("%s:%s:%s", p.projectRepo.hosting.Host(), p.Repository.Owner.Name, p.Repository.Name)
}
