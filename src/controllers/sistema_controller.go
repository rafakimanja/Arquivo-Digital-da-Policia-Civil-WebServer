package controllers

import (
	"adpc-webserver/src/database"
	"adpc-webserver/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeConfSistema(c *gin.Context) {
	var configuracoes models.Sistema
	resp := database.DB.Last(&configuracoes)
	if resp.RowsAffected == 0 {
		c.HTML(http.StatusOK, "configuracoes.html", gin.H{
			"configuracoes": nil,
		})
	} else {
		usuario, _ := c.Get("Usuario")
		c.HTML(http.StatusOK, "configuracoes.html", gin.H{
			"configuracoes": configuracoes,
			"Usuario": usuario,
		})
	}
}

func SalvaConfSistema(c *gin.Context) {
	var configuracoes models.Sistema

	upload_only := c.PostForm("only")
	categoria := c.PostForm("category")

	if upload_only == "on" {configuracoes.UploadUnico=true} else {configuracoes.UploadUnico=false}

	configuracoes.TipoArq = categoria

	database.DB.Create(&configuracoes)
	c.Redirect(http.StatusSeeOther, "/index/config")
}
