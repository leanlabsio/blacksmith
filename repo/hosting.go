package repo

type Hosting interface {
	ListRepositories() []Repository
	GetRepository(namespace, name string) Repository
	CreateWebhook(namespace, name string)
	Host() string
	//	RemoveWebhook()
}
