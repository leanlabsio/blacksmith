package repo

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHub struct {
	client *github.Client
}

func convertToRepository(r github.Repository) Repository {
	return Repository{
		CloneURL:    *r.CloneURL,
		FullName:    *r.FullName,
		Name:        *r.Name,
		Description: *r.Description,
	}
}

func NewGithub(token oauth2.TokenSource) GitHub {
	tc := oauth2.NewClient(oauth2.NoContext, token)
	client := github.NewClient(tc)
	return GitHub{
		client: client,
	}
}

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
