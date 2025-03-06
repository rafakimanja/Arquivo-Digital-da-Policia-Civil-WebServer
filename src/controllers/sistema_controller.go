package controllers

import (
	"adpc-webserver/src/database"
	"adpc-webserver/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeConfSistema(c *gin.Context) {
	c.HTML(http.StatusOK, "configuracoes.html", nil)
}

func SalvaConfSistema(c *gin.Context) {
	var configuracoes models.Sistema

	if err := c.ShouldBindJSON(&configuracoes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&configuracoes)
	c.JSON(http.StatusOK, configuracoes)
}
