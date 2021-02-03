package main

import (
	"github.com/google/wire"

	"github.com/covergates/covergates/config"
	"github.com/covergates/covergates/core"
	"github.com/covergates/covergates/modules/charts"
	"github.com/covergates/covergates/modules/git"
	"github.com/covergates/covergates/modules/hook"
	"github.com/covergates/covergates/modules/oauth"
	"github.com/covergates/covergates/modules/repo"
	"github.com/covergates/covergates/modules/report"
	"github.com/covergates/covergates/modules/scm"
	"github.com/covergates/covergates/modules/session"
	"github.com/covergates/covergates/service/coverage"
)

// nolint:deadcode,varcheck,unused
var serviceSet = wire.NewSet(
	provideSCMService,
	provideSession,
	provideCoverageService,
	provideChartService,
	provideGit,
	provideReportService,
	provideHookService,
	provideOAuthService,
	provideRepoService,
)

func provideSCMService(
	config *config.Config,
	userStore core.UserStore,
	git core.Git,
) core.SCMService {
	return &scm.Service{
		Config:    config,
		UserStore: userStore,
		Git:       git,
	}
}

func provideSession() core.Session {
	return &session.Session{}
}

func provideCoverageService() core.CoverageService {
	return &coverage.Service{}
}

func provideChartService() core.ChartService {
	return &charts.ChartService{}
}

func provideGit() core.Git {
	return &git.Service{}
}

func provideReportService(
	config *config.Config,
	repoStore core.RepoStore,
) core.ReportService {
	return &report.Service{
		Config:    config,
		RepoStore: repoStore,
	}
}

func provideHookService(
	scm core.SCMService,
	repoStore core.RepoStore,
	reportStore core.ReportStore,
	reportService core.ReportService,
) core.HookService {
	return &hook.Service{
		SCM:           scm,
		RepoStore:     repoStore,
		ReportService: reportService,
		ReportStore:   reportStore,
	}
}

func provideOAuthService(
	config *config.Config,
	oAuthStore core.OAuthStore,
	userStore core.UserStore,
) core.OAuthService {
	return oauth.NewService(config, oAuthStore, userStore)
}

func provideRepoService(
	config *config.Config,
	scmService core.SCMService,
	userStore core.UserStore,
	repoStore core.RepoStore,
) core.RepoService {
	return repo.NewService(config, scmService, userStore, repoStore)
}
