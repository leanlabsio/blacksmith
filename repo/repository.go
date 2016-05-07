package repo

type Repository struct {
	CloneURL    string `json:"clone_url"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Owner       Owner  `json:"owner"`
}

type Owner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
