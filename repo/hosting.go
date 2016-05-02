package repo

type Hosting interface {
	//	GetRepository() Repository
	ListRepositories() []Repository
	//	CreateWebhook()
	//	RemoveWebhook()
}
