package controllers

import (
	"adpc-webserver/src/database"
	"adpc-webserver/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTelaUsuarios(c *gin.Context){
	usuario, _ := c.Get("Usuario")
	c.HTML(http.StatusOK, "usuarios.html", gin.H{
		"Usuario": usuario,
	})
}

func ExibeTodosUsuarios(c *gin.Context) {
	var usuarios []models.Usuario
	database.DB.Find(&usuarios)
	c.JSON(http.StatusOK, gin.H{
		"usuarios": usuarios,
	})
}

func ExibeFormUsuario(c *gin.Context) {
	usuario, _ := c.Get("Usuario")
	c.HTML(http.StatusOK, "form-usuarios.html", gin.H{
		"Usuario": usuario,
	})
}

func CriaNovoUsuario(c *gin.Context) {
	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&usuario)
	c.Status(http.StatusCreated)
}

func BuscaUsuario(c *gin.Context) {
	var usuario models.Usuario

	id := c.Params.ByName("id")
	database.DB.First(&usuario, id)

	if usuario.ID == 0 {
		c.HTML(http.StatusNotFound, "erro.html", gin.H{
			"code": http.StatusNotFound,
			"error": "Usuario nao encontrado!",
		})
	} else {
		c.HTML(http.StatusOK, "form-usuarios.html", gin.H{
			"Update": true,
			"Usuario": usuario,
		})
	}
}

func DeletaUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	result := database.DB.Delete(&usuario, id)

	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code": http.StatusInternalServerError,
			"error": "Erro ao deletar usuario!"})
	} else {
		c.Status(http.StatusOK)
	}		
}

func AtualizaUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.First(&usuario, id)

	if usuario.ID == 0 {
		c.HTML(http.StatusNotFound, "erro.html", gin.H{
			"code": http.StatusNotFound,
			"message": "Usuario nao encontrado!",
		})
		return
	}

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.HTML(http.StatusBadRequest, "erro.html", gin.H{
			"code": http.StatusBadRequest,
			"error": "Erro ao atualizar usuario!",
		})
		return
	}
	
	if err := database.DB.Model(&usuario).Updates(usuario); err != nil {
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code": http.StatusInternalServerError,
			"error": "Erro ao atualizar usuario!",
		})
		return
	}
	
	c.Status(http.StatusOK)
}
