package repo

import (
	"github.com/gin-gonic/gin"

	"github.com/covergates/covergates/core"
)

const keyRepo = "repository"

// WithRepo in context
func WithRepo(store core.RepoStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo, err := store.Find(&core.Repo{
			NameSpace: c.Param("namespace"),
			Name:      c.Param("name"),
			SCM:       core.SCMProvider(c.Param("scm")),
		})
		if err != nil {
			c.String(400, err.Error())
			c.Abort()
			return
		}
		c.Set(keyRepo, repo)
	}
}
