package web

import (
	"github.com/gin-gonic/gin"

	"github.com/covergates/covergates/core"
)

// TokenFrom context
func TokenFrom(c *gin.Context) *core.Token {
	return &core.Token{
		Token:   c.GetString(keyAccess),
		Refresh: c.GetString(keyRefresh),
		Expires: c.GetTime(keyExpires),
	}
}
