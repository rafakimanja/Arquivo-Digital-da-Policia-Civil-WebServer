package controllers

import (
	"adpc-webserver/src/database"
	"adpc-webserver/src/models"
	"adpc-webserver/src/services"
	"fmt"
	"net/http"
	"time"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"os"
)

func ExibeTelaLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", nil)
}

func ExibeTelaCadastro(c *gin.Context){
	c.HTML(http.StatusOK, "cadastro.html", nil)
}

func ExibeTelaRecSenha(c *gin.Context){
	c.HTML(http.StatusOK, "rec-senha", nil)
}

func ExibeTelaIndex(c *gin.Context) {
	usuario, _ := c.Get("Usuario")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Usuario": usuario,
	})
}

func ExibeTelaNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "erro.html", gin.H{
		"code": http.StatusNotFound,
		"message": "Nao encontramos esta pagina!",
	})
}

func ExibeTelaAjuda(c *gin.Context){
	usuario, _ := c.Get("Usuario")
	c.HTML(http.StatusOK, "ajuda.html", gin.H{
		"Usuario": usuario,
	})
}

func LoginAcess(c *gin.Context) {

	type LoginUser struct {
		Usuario    string
		Senha 	   string
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
		c.HTML(http.StatusNotFound, "login", gin.H{
			"Message": "Usuario inválido!",
			"Status":  false,
		})
		return
	}

	var session models.Session

	token := services.GenerateToken()
	session.Token = token
	session.UserID = user.ID
	session.Expired = time.Now().Add(time.Minute * 15)

	if err := database.DB.Create(&session).Error; err != nil {
		fmt.Println("Erro ao salvar sessão no banco:", err)
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Erro ao criar sessão.",
		})
		return
	}

	_ = godotenv.Load("/.env")
	ip_server := os.Getenv("IP_SERVER'")

	c.SetCookie("session_token", session.Token, 3600, "/", ip_server, false, true)
	c.Redirect(http.StatusFound, "/index")
}

func LogoutSession(c *gin.Context) {
	var session models.Session

	sessionToken, err := c.Cookie("session_token")

	_ = godotenv.Load("/.env")
	ip_server := os.Getenv("IP_SERVER'")

	if err != nil {
		c.SetCookie("session_token", "", -1, "/", ip_server, false, true)
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Erro ao resgatar o cookie de sessao",
		})
		return
	}

	database.DB.Where("token = ?", sessionToken).First(&session)

	id := session.ID

	if id == 0 {
		c.SetCookie("session_token", "", -1, "/", ip_server, false, true)
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	result := database.DB.Delete(&session, id)
	if result.Error != nil {
		c.SetCookie("session_token", "", -1, "/", ip_server, false, true)
		c.HTML(http.StatusInternalServerError, "erro.html", gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Erro ao excluir o cookie de sessao",
		})
		return
	}

	c.SetCookie("session_token", "", -1, "/", ip_server, false, true)
	c.Redirect(http.StatusSeeOther, "/")
}
