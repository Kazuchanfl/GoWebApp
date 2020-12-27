package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
SetupUserRouter でユーザー関連のリクエストのルーティング
*/
func SetupUserRouter(r *gin.Engine) {
	// Ping test
	r.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "Kazuaki")
	})
}
