package repo

import (
	"github.com/gin-gonic/gin"

	"github.com/covergates/covergates/core"
)

// HandleHookCreate for the repository
// @Summary create repository webhook
// @Tags Repository
// @Param scm path string true "SCM"
// @Param namespace path string true "Namespace"
// @Param name path string true "name"
// @Success 200 {object} string ok
// @Router /repos/{scm}/{namespace}/{name}/hook/create [post]
func HandleHookCreate(service core.HookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo, _ := c.MustGet(keyRepo).(*core.Repo)
		ctx := c.Request.Context()
		if err := service.Create(ctx, repo); err != nil {
			c.String(500, err.Error())
			return
		}
		c.String(200, "ok")
	}
}

// HandleHook event
// @Summary handle webhook event
// @Tags Repository
// @Param scm path string true "SCM"
// @Param namespace path string true "Namespace"
// @Param name path string true "name"
// @Success 200 {object} string ok
// @Router /repos/{scm}/{namespace}/{name}/hook [post]
func HandleHook(scm core.SCMService, service core.HookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := c.MustGet(keyRepo).(*core.Repo)
		ctx := c.Request.Context()
		client, err := scm.Client(repo.SCM)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		hook, err := client.Webhooks().Parse(c.Request)
		if err != nil && client.Webhooks().IsWebhookNotSupport(err) {
			c.String(200, "ok")
			return
		} else if err != nil {
			c.String(500, err.Error())
			return
		}
		if err := service.Resolve(ctx, repo, hook); err != nil {
			c.String(500, err.Error())
			return
		}
		c.String(200, "ok")
	}
}
