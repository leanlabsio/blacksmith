package project

import (
	"github.com/leanlabsio/blacksmith/repo"
)

//Project represents single project to build
type Project struct {
	Builder    Builder         `json:"builder,omitempty"`
	Repository repo.Repository `json:"repository"`
	EnvVars    []Env           `json:"env"`
	Enabled    bool            `json:"enabled"`
}

// Builder describes task executor reference
type Builder struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

// Env represents any additional confugration parameters
// to be passed to Builder, in key - value format
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
