package repo

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"net/url"
)

type GitHub struct {
	client *github.Client
	host   string
}

func convertToRepository(r github.Repository) Repository {
	return Repository{
		CloneURL:    *r.CloneURL,
		FullName:    *r.FullName,
		Name:        *r.Name,
		Description: *r.Description,
		Owner: Owner{
			Name: *r.Owner.Login,
			ID:   *r.Owner.ID,
		},
	}
}

// NewGithub returns new GitHub repository hosting instance
func NewGithub(token oauth2.TokenSource) *GitHub {
	tc := oauth2.NewClient(oauth2.NoContext, token)
	client := github.NewClient(tc)
	log.Printf("%s", client.BaseURL)
	return &GitHub{
		client: client,
		host:   "github.com",
	}
}

// Host returns repo hosting domain name
func (g *GitHub) Host() string {
	return g.host
}

// ListRepositories returns github repositories list
func (g *GitHub) ListRepositories() []Repository {
	opts := &github.RepositoryListOptions{
		Type:        "all",
		ListOptions: github.ListOptions{PerPage: 100, Page: 1},
	}

	repos, _, _ := g.client.Repositories.List("", opts)

	ret := []Repository{}

	for _, repo := range repos {
		ghr := convertToRepository(repo)
		ret = append(ret, ghr)
	}

	return ret
}

func (g *GitHub) GetRepository(namespace, name string) Repository {
	repo, _, _ := g.client.Repositories.Get(namespace, name)

	ret := convertToRepository(*repo)

	return ret
}

// CreateWebhook creates webhook to Repository
func (g *GitHub) CreateWebhook(namespace, name string) {
	wh := url.URL{
		Scheme: "http",
		Host:   "qwerty.com", // todo pass host from config2
		Path:   "trigger",
	}
	v := url.Values{}
	v.Add("host", "github.com")
	v.Add("namespace", namespace)
	v.Add("name", name)
	wh.RawQuery = v.Encode()
	hook := github.Hook{
		Name:   github.String("web"),
		Active: github.Bool(true),
		Events: []string{"push", "pull_request"},
		Config: map[string]interface{}{
			"url":          wh.String(),
			"content_type": "json",
		},
	}

	g.client.Repositories.CreateHook(namespace, name, &hook)
}
