package github

//Push represents github push webhook payload
type Push struct {
	Ref          string   `json:"ref"`
	Head         string   `json:"head"`
	Before       string   `json:"before"`
	Size         int64    `json:"size"`
	DistinctSize int64    `json:"distinct_size"`
	Commits      []Commit `json:"commits"`
}

//Commit represents github webhook payload commit info
type Commit struct {
	Sha       string `json:"sha"`
	Message   string `json:"message"`
	Author    User   `json:"author"`
	URL       string `json:"url"`
	Distrinct bool   `json:"distinct"`
}

//User represents github webhook payload user info
type User struct{}
