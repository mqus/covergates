package git

import (
	"context"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/covergates/covergates/core"
)

// Service provides git operations
type Service struct{}

// Clone repository in memory
func (s *Service) Clone(ctx context.Context, url, token string) (core.GitRepository, error) {
	store := memory.NewStorage()
	repo, err := git.CloneContext(ctx, store, nil, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: token,
			Password: "x-oauth-basic",
		},
	})
	if err != nil {
		return nil, err
	}
	return &repository{
		gitRepository: repo,
	}, nil
}

// PlainOpen a repository
func (s *Service) PlainOpen(ctx context.Context, path string) (core.GitRepository, error) {
	repo, err := git.PlainOpenWithOptions(path, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return nil, err
	}
	return &repository{
		gitRepository: repo,
	}, nil
}
