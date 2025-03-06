package database

import (
	"adpc-webserver/src/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// Função para conectar ao banco de dados com varias tentativas
func conectaDbRetry() error {
	dir, _ := os.Getwd()
	envPath := dir + "/.env"
	err = godotenv.Load(envPath)
	if err != nil {
		return fmt.Errorf("erro ao carregar o arquivo .env: %v", err)
	}

	strConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	for retries := 0; retries < 5; retries++ {
		DB, err = gorm.Open(postgres.Open(strConexao), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Erro ao conectar com o banco. Tentando novamente... (tentativa %d/5)", retries+1)
		time.Sleep(time.Duration(2<<retries) * time.Second)
	}

	if err != nil {
		return fmt.Errorf("não foi possível conectar ao banco de dados: %v", err)
	}

	return nil
}

// Função para rodar o AutoMigrate após a conexão com o banco
func ConectaDB() (*gorm.DB, error) {
	// Tenta conectar ao banco com retry
	err := conectaDbRetry()
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&models.Usuario{}, &models.Sistema{}, &models.Documento{})
	if err != nil {
		return nil, fmt.Errorf("erro ao rodar o AutoMigrate: %v", err)
	}

	return DB, nil
}
