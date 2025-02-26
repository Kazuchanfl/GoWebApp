package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
SetupRouter でルーティングの定義とHTTPサーバーの起動
*/
func SetupRouter() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// ユーザー関連のリクエストのルーティング
	SetupUserRouter(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
