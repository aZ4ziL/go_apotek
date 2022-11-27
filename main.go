package main

import (
	"encoding/gob"
	"time"

	"github.com/aZ4zil/go_apotek/controllers"
	"github.com/aZ4zil/go_apotek/models"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(time.Time{})
	gob.Register(map[string]interface{}{})
}

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Static("/static", "./static")

	router.HTMLRender = createMyRender()

	// Set Session
	store := gormsessions.NewStore(models.Connection(), true, []byte("mysecretkey"))
	router.Use(sessions.Sessions("ginSessionID", store))

	controllers.Controllers(router)

	router.Run(":8888")
}
