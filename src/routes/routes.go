package routes

import (
	"adpc-webserver/src/controllers"
	"adpc-webserver/src/middlewares"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", controllers.ExibeTelaLogin)
	r.GET("/cadastro", controllers.ExibeTelaCadastro)
	r.GET("/recSenha", controllers.ExibeTelaRecSenha)
	r.POST("/login", controllers.LoginAcess)
	indexGroup := r.Group("/index", middlewares.SessionMiddleware)
	{
		indexGroup.GET("/", controllers.ExibeTelaIndex)
		indexGroup.GET("/documentos", controllers.ExibeTelaDocumentos)
		indexGroup.GET("/documentos/json", controllers.ExibeTodosDocumentos)
		indexGroup.GET("/documentos/form", controllers.ExibeFormDocumentos)
		indexGroup.GET("/documentos/:id", controllers.BuscaArquivo)
		indexGroup.GET("/documentos/download/:id", controllers.BaixaArquivo)
		indexGroup.POST("/documentos", controllers.CriaNovoArquivo)
		indexGroup.PUT("/documentos/:id", controllers.AtualizaArquivo)
		indexGroup.DELETE("/documentos/:id", controllers.DeletaArquivo)
		indexGroup.GET("/config", controllers.ExibeConfSistema)
		indexGroup.POST("/config", controllers.SalvaConfSistema)
		indexGroup.GET("/usuarios", controllers.ExibeTelaUsuarios)
		indexGroup.GET("/usuarios/json", controllers.ExibeTodosUsuarios)
		indexGroup.GET("/usuarios/:id", controllers.BuscaUsuario)
		indexGroup.GET("/usuarios/form", controllers.ExibeFormUsuario)
		indexGroup.POST("/usuarios", controllers.CriaNovoUsuario)
		indexGroup.PATCH("usuarios/:id", controllers.AtualizaUsuario)
		indexGroup.DELETE("/usuarios/:id", controllers.DeletaUsuario)
		indexGroup.GET("/ajuda", controllers.ExibeTelaAjuda)
		indexGroup.POST("/logout", controllers.LogoutSession)
	}
	r.NoRoute(controllers.ExibeTelaNotFound)
	r.Run(":5000")
}
