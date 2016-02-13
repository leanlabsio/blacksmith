package model

//Job represents single API payload entry
type Job struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
	EnvVars    []Env  `json:"env"`
	Enabled    bool   `json:"enabled"`
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
