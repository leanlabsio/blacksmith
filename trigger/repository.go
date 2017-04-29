package trigger

import (
	"github.com/leanlabsio/blacksmith/repo"
)

type Repository struct {
	hosting repo.Hosting
}

func NewRepository(c repo.Hosting) *Repository {
	return &Repository{
		hosting: c,
	}
}

func (t *Repository) CreateTrigger(r repo.Repository) {
	t.hosting.CreateWebhook(r.Owner.Name, r.Name)
}
