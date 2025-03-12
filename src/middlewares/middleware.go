package middlewares

import (
	"adpc-webserver/src/database"
	"adpc-webserver/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SessionMiddleware(c *gin.Context){
	var session models.Session

	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		c.Abort()
		return
	}

	database.DB.Where("token = ?", sessionToken).First(&session)
	if session.ID == 0 || session.Expired.Before(time.Now()){
		database.DB.Delete(&session, session.ID)
		c.SetCookie("session_token", "", -1, "/", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/")
		c.Abort()
		return
	}
	c.Next()
}