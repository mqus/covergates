package scm

import (
	"errors"
	"net/http"

	"github.com/drone/go-scm/scm"

	"github.com/covergates/covergates/config"
	"github.com/covergates/covergates/core"
)

var errWebhookNotSuport = errors.New("webhook not support")

type webhookService struct {
	config *config.Config
	client *scm.Client
	scm    core.SCMProvider
}

func (service *webhookService) Parse(req *http.Request) (core.HookEvent, error) {
	cfg := service.config
	hook, err := service.client.Webhooks.Parse(req, func(webhook scm.Webhook) (string, error) {
		return cfg.Server.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if event, ok := hook.(*scm.PullRequestHook); ok {
		if event.Action == scm.ActionMerge {
			return &core.PullRequestHook{
				Number: event.PullRequest.Number,
				Merged: true,
				Commit: event.PullRequest.Sha,
				Source: event.PullRequest.Source,
				Target: event.PullRequest.Target,
			}, nil
		} else if service.scm == core.Gitea && event.Action == scm.ActionClose {
			return &core.PullRequestHook{
				Number: event.PullRequest.Number,
				Merged: true,
				Commit: event.PullRequest.Sha,
				Source: event.PullRequest.Source,
				Target: event.PullRequest.Target,
			}, nil
		}
	}

	return nil, errWebhookNotSuport
}

func (service *webhookService) IsWebhookNotSupport(err error) bool {
	return err == errWebhookNotSuport
}
