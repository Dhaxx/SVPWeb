package database

import (
	"database/sql"
	"path/filepath"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/nakagami/firebirdsql"
)

func Connect() (*sql.DB, error) {
	currentDir, err := os.Getwd() // Pega o dir atual
	if err != nil {
		log.Fatalf("Erro ao obter diretório atual: %v", err)
	}

	rootPath := filepath.Join(currentDir, "..") // Sobe um diretório
	envPath := filepath.Join(rootPath, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	path := os.Getenv("DB_PATH")

	dsn := fmt.Sprintf("%s:%s@%s:%s/%s", user, password, host, port, path)

	db, err := sql.Open("firebirdsql", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	return db, err
}
