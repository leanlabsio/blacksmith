package project

import (
	"github.com/leanlabsio/blacksmith/executor"
	"github.com/leanlabsio/blacksmith/repo"
	"github.com/leanlabsio/blacksmith/trigger"
)

// Project represents single project to build
type Project struct {
	Repository repo.Repository         `json:"repository"`
	Trigger    trigger.Trigger         `json:"trigger"`
	Executor   executor.DockerExecutor `json:"executor"`

	triggerRepo *trigger.TriggerRepository
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
