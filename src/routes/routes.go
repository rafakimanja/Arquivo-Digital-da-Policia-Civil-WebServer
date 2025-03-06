package routes

import (
	"adpc-webserver/src/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", controllers.ExibeTelaLogin)
	r.POST("/login", controllers.LoginAcess)
	indexGroup := r.Group("/index")
	{
		indexGroup.GET("/", controllers.ExibeTelaIndex)
		indexGroup.GET("/documentos", controllers.ExibeTodosDocumentos)
		indexGroup.GET("/documentos/form", controllers.ExibeFormDocumentos)
		indexGroup.GET("/documentos/:id", controllers.BuscaArquivo)
		indexGroup.GET("/documentos/donwload/:id", controllers.BaixaArquivo)
		indexGroup.POST("/documentos", controllers.CriaNovoArquivo)
		indexGroup.PUT("/documentos/:id", controllers.AtualizaArquivo)
		indexGroup.DELETE("/documentos/:id", controllers.DeletaArquivo)
		indexGroup.GET("/config", controllers.ExibeConfSistema)
		indexGroup.POST("/config", controllers.SalvaConfSistema)
		indexGroup.GET("/usuarios", controllers.ExibeTodosUsuarios)
		indexGroup.GET("/usuarios/:id", controllers.BuscaUsuario)
		indexGroup.GET("/usuarios/form", controllers.ExibeFormUsuario)
		indexGroup.POST("/usuarios/form", controllers.CriaNovoUsuario)
		indexGroup.PATCH("usuarios/:id", controllers.AtualizaUsuario)
		indexGroup.DELETE("/usuarios/:id", controllers.DeletaUsuario)
	}
	r.Run(":5000")
}
