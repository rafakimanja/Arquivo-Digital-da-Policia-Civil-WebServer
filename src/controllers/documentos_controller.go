package controllers

import (
	"adpc-webserver/src/database"
	"adpc-webserver/src/models"
	"adpc-webserver/src/services"
	"fmt"
	"net/http"
	"os"

	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExibeTelaDocumentos(c *gin.Context){
	usuario, _ := c.Get("Usuario")
	c.HTML(http.StatusOK, "documentos.html", gin.H{
		"Usuario": usuario,
	})
}

func ExibeTodosDocumentos(c *gin.Context){
	var documentos []models.Documento
	database.DB.Find(&documentos)
	c.JSON(http.StatusOK, gin.H{
		"documentos": documentos,
	})
}

func ExibeFormDocumentos(c *gin.Context) {
	usuario, _ := c.Get("Usuario")
	c.HTML(http.StatusOK, "form-documento.html", gin.H{
		"Usuario": usuario,
	})
}

func CriaNovoArquivo(c *gin.Context) {

	nome := c.PostForm("nome")
	categoria := c.PostForm("categoria")
	ano := c.PostForm("ano")
	documento, err := c.FormFile("arquivo")
	if err != nil {
		c.HTML(http.StatusBadRequest, "erro.html",gin.H{
			"code": http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	anoInt, _ := strconv.Atoi(ano)
	arquivo := models.Documento{Nome: nome, Ano: anoInt, Categoria: categoria, Arquivo: nome + filepath.Ext(documento.Filename)}

	gerenciador := services.Construtor()
	_, path := gerenciador.SalvaArquivo(arquivo)

	pathDocumento := path + "/" + arquivo.Arquivo
	err = c.SaveUploadedFile(documento, pathDocumento)
	database.DB.Create(&arquivo)

	if err != nil {
		c.HTML(http.StatusBadRequest, "erro.html",gin.H{
			"code": http.StatusBadRequest,
			"message": "erro no upload do arquivo!",
		})
		return
	}

	c.Status(http.StatusCreated)
}

func BuscaArquivo(c *gin.Context) {
	var documento models.Documento

	id := c.Params.ByName("id")
	database.DB.First(&documento, id)

	usuario, _ := c.Get("Usuario")

	if documento.ID == 0 {
		c.HTML(http.StatusNotFound, "erro.html", gin.H{
			"code": http.StatusNotFound,
			"message": "Arquivo nao encontrado!",
		})
	} else {
		c.HTML(http.StatusOK, "form-documento.html", gin.H{
			"Update": true,
			"Documento": documento,
			"Usuario": usuario,
		})
	}
}

func BaixaArquivo(c *gin.Context) {

	var documento models.Documento

	idArq := c.Params.ByName("id")
	database.DB.First(&documento, idArq)
	
	if documento.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	path := "./arquivos/" + fmt.Sprint(documento.Ano) + "/" + documento.Categoria + "/" + documento.Arquivo
	file, err := os.Open(path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code": http.StatusInternalServerError,
			"message": "Erro ao abrir arquivo",
		})
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	extraHeaders := map[string]string{
		"Content-Disposition": "attachment; filename="+documento.Arquivo,
	}

	c.DataFromReader(http.StatusOK, fileInfo.Size(), "application/pdf", file, extraHeaders)
}

func DeletaArquivo(c *gin.Context) {
	var documento models.Documento
	id := c.Params.ByName("id")
	database.DB.First(&documento, id)

	if documento.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	gerenciador := services.Construtor()

	if !gerenciador.DeletaArquivo(documento, false) {
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code": http.StatusInternalServerError,
			"message": "Erro ao deletar arquivo!",
		})
		return
	}

	result := database.DB.Delete(&documento, id)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code": http.StatusInternalServerError,
			"message": "Erro ao deletar arquivo!",
		})
		return
	}

	if !gerenciador.DeletaArquivo(documento, true) {
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code": http.StatusInternalServerError,
			"message": "Erro ao deletar arquivo!",
		})
		return
	}

	c.Status(http.StatusOK)
}

func AtualizaArquivo(c *gin.Context) {
	var documento models.Documento
	id := c.Params.ByName("id")
	database.DB.First(&documento, id)

	if documento.ID == 0 {
		c.HTML(http.StatusNotFound, "erro.html", gin.H{
			"code": http.StatusNotFound,
			"message": "Arquivo nao encontrado!"})
		return
	}

	nome := c.PostForm("nome")
	categoria := c.PostForm("categoria")
	ano := c.PostForm("ano")
	arquivo, err := c.FormFile("arquivo")

	anoInt, _ := strconv.Atoi(ano)
	documentoAtt := models.Documento{Nome: nome, Ano: int(anoInt), Categoria: categoria}

	gerenciador := services.Construtor()

	if err != nil {
		documentoAtt.Arquivo = nome+filepath.Ext(documento.Arquivo)

		if !gerenciador.AtualizarArquivo(documento, documentoAtt) {
			c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
				"code": http.StatusInternalServerError,
				"message": "Erro ao atualizar arquivo!"})
			return
		}

	} else {

		documentoAtt.Arquivo = nome+filepath.Ext(arquivo.Filename)

		_, path := gerenciador.SalvaArquivo(documentoAtt)
		pathDoc := path+"/"+documentoAtt.Arquivo

		err = c.SaveUploadedFile(arquivo, pathDoc)

		if err != nil {
			c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
				"code": http.StatusInternalServerError,
				"message": "Erro no upload do arquivo!"})
			return
		}

		if !gerenciador.DeletaArquivo(documento, false) || !gerenciador.DeletaArquivo(documento, true){
			c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
				"code": http.StatusInternalServerError,
				"message": "Erro ao atualizar arquivo!"})
			return
		}
	}

	if err := database.DB.Model(&documento).Updates(documentoAtt).Error; err != nil{
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code": http.StatusInternalServerError,
			"message": "Erro ao atualizar arquivo!"})
		return
	}

	c.Status(http.StatusOK)
}
