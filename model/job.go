package model

//Project represents single API payload entry
type Project struct {
	Name        string  `json:"name"`
	Builder     Builder `json:"builder,omitempty"`
	FullName    string  `json:"full_name"`
	Repository  string  `json:"clone_url"`
	EnvVars     []Env   `json:"env"`
	Enabled     bool    `json:"enabled"`
	Avatar      string  `json:"avatar"`
	Description string  `json:"description"`
}

type Builder struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
