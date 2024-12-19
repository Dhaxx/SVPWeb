package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/nakagami/firebirdsql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = connect() // Inicializa a conexão apenas uma vez
		if err != nil {
			log.Fatalf("Erro ao inicializar conexão com o banco: %v", err)
		}
	})
	return db
}

func connect() (*sql.DB, error) {
	currentDir, err := os.Getwd() // Pega o dir atual
	if err != nil {
		return nil, fmt.Errorf("erro ao obter diretório atual: %v", err)
	}

	rootPath := filepath.Join(currentDir, "..") // Sobe um diretório
	// envPath := filepath.Join(rootPath, ".env")
	envPath := filepath.Join(rootPath, "\\SVPWEB\\.env")

	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("erro ao carregar .env: %v", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	path := os.Getenv("DB_PATH")

	dsn := fmt.Sprintf("%s:%s@%s:%s/%s", user, password, host, port, path)

	db, err := sql.Open("firebirdsql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com o banco de dados: %v", err)
	}

	// Testa a conexão
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao pingar o banco: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")

	return db, nil
}
