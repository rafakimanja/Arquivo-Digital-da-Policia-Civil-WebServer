package middlewares

import (
	"adpc-webserver/src/database"
	"adpc-webserver/src/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SessionMiddleware(c *gin.Context) {
	var session models.Session
	var user models.Usuario

	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		c.Abort()
		return
	}

	database.DB.Where("token = ?", sessionToken).First(&session)
	if session.ID == 0 || session.Expired.Before(time.Now()) {
		database.DB.Delete(&session, session.ID)
		_ = godotenv.Load("/.env")
		ip_server := os.Getenv("IP_SERVER'")
		c.SetCookie("session_token", "", -1, "/", ip_server, false, true)
		c.Redirect(http.StatusSeeOther, "/")
		c.Abort()
		return
	}

	database.DB.First(&user, session.UserID)

	type userShort struct {
		Nome  string
		Admin bool
	}

	contextUser := userShort{Nome: user.Nome, Admin: user.Admin}

	c.Set("Usuario", contextUser)

	c.Next()
}
