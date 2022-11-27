package handlers

import (
	"net/http"

	"github.com/aZ4zil/go_apotek/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetALlDrug(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := session.Get("user")

	if user == nil {
		flasher.Set("info", "Mohon login terlebih dahulu sebelum Anda mengakses halaman ini.")
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	if ctx.Request.Method == "GET" {
		defer flasher.Del()

		drugs := models.GetAllGrugs()
		ctx.HTML(http.StatusOK, "get_all_drugs", gin.H{
			"user":  user,
			"drugs": drugs,
		})
	}
}
