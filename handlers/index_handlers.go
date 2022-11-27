package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	defer flasher.Del()

	session := sessions.Default(ctx)

	user := session.Get("user")

	ctx.HTML(http.StatusOK, "index", gin.H{
		"user":    user,
		"flasher": flasher,
	})
}
