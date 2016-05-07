package trigger

import (
	"github.com/leanlabsio/blacksmith/repo"
)

type TriggerRepository struct {
	hosting repo.Hosting
}

func NewRepository(c repo.Hosting) *TriggerRepository {
	return &TriggerRepository{
		hosting: c,
	}
}

func (t *TriggerRepository) CreateTrigger(r repo.Repository) {
	t.hosting.CreateWebhook(r.Owner.Name, r.Name)
}
