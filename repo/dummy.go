package repo

type DummyHosting struct {
	host string
}

func NewDummy(host string) *DummyHosting {
	return &DummyHosting{
		host: host,
	}
}

func (d *DummyHosting) Host() string {
	return d.host
}

func (d *DummyHosting) ListRepositories() []Repository {
	return []Repository{}
}

func (d *DummyHosting) GetRepository(namespace, name string) Repository {
	return Repository{
		Name: name,
		Owner: Owner{
			Name: namespace,
		},
	}
}

func (d *DummyHosting) CreateWebhook(namespace, name string) {}
