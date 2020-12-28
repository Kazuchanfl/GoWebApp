package router

import (
	"net/http"

	"github.com/Kazuchanfl/GoWebApp/controllers"
	"github.com/gin-gonic/gin"
)

/*
SetupUserRouter でユーザー関連のリクエストのルーティング
*/
func SetupUserRouter(r *gin.Engine) {
	// GET /users
	r.GET("/users", func(c *gin.Context) {
		users := controllers.GetAllUsers()
		c.JSON(http.StatusOK, users)
	})
}
