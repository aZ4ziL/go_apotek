package handlers

import (
	"net/http"
	"time"

	"github.com/aZ4zil/go_apotek/auth"
	"github.com/aZ4zil/go_apotek/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Login
// handler for login user
func Login(ctx *gin.Context) {
	session := sessions.Default(ctx)

	if user := session.Get("user"); user != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}

	if ctx.Request.Method == "GET" {
		defer flasher.Del()

		ctx.HTML(http.StatusOK, "login", gin.H{
			"flasher": flasher,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")

		user, err := models.GetUserByEmail(email)
		if err != nil {
			flasher.Set("danger", "Email atau katasandi yang Anda masukkan salah. 1")
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		if !auth.DecryptionPassword(user.Password, password) {
			flasher.Set("danger", "Email atau katasandi yang Anda masukkan salah. 2")
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		user.LastLogin = time.Now()
		models.Connection().Save(&user)

		userSession := map[string]interface{}{
			"id":          user.ID,
			"full_name":   user.FirstName + " " + user.LastName,
			"email":       user.Email,
			"username":    user.Username,
			"last_login":  user.LastLogin,
			"date_joined": user.DateJoined,
			"is_admin":    user.IsAdmin,
			"is_active":   user.IsActive,
		}

		session.Set("user", userSession)
		if err := session.Save(); err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx.Redirect(http.StatusFound, "/")
	}
}

// Register
// handler for register user
func Register(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.Redirect(http.StatusOK, "/")
		return
	}

	if ctx.Request.Method == "GET" {
		defer flasher.Del()

		ctx.HTML(http.StatusOK, "register", gin.H{
			"flasher": flasher,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		firstName := ctx.PostForm("first_name")
		lastName := ctx.PostForm("last_name")
		username := ctx.PostForm("username")
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")

		user := models.User{
			FirstName: firstName,
			LastName:  lastName,
			Username:  username,
			Email:     email,
			Password:  password,
		}

		if err := models.CreateNewUser(&user); err != nil {
			flasher.Set("danger", err.Error())
			ctx.Redirect(http.StatusFound, "/register")
			return
		}

		flasher.Set("success", "Selamat akun Anda telah berhasil didaftarkan.")
		ctx.Redirect(http.StatusFound, "/login")
		return
	}
}
