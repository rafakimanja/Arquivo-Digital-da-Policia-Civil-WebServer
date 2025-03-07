package controllers

import (
	"adpc-webserver/src/database"
	"adpc-webserver/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTelaLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", nil)
}

func LoginAcess(c *gin.Context) {

	type LoginUser struct {
		RG    string
		Senha string
	}

	var userLogin LoginUser
	var user models.Usuario

	if err := c.ShouldBind(&userLogin); err != nil {
		c.HTML(http.StatusBadRequest, "login", gin.H{
			"Message": err.Error(),
			"Status":  false,
		})
		return
	}

	database.DB.First(&user, userLogin)

	if user.ID == 0 {
		c.HTML(http.StatusNotFound, "login.html", gin.H{
			"Message": "Usuario inválido!",
			"Status":  false,
		})
		return
	}

	c.Redirect(http.StatusFound, "/index")
}

func ExibeTelaIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
