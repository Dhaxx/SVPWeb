package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

var secretKey []byte

func init() {
	currentDir, err := os.Getwd() // Pega o dir atual
	if err != nil {
		fmt.Printf("erro ao obter diretório atual: %v", err)
	}

	rootPath := filepath.Join(currentDir, "..") // Sobe um diretório
	envPath := filepath.Join(rootPath, "SVPWEB", ".env")

	if err := godotenv.Load(envPath); err != nil {
		fmt.Printf("erro ao carregar .env: %v\n", err)
		return
	}

	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	// Verifica se a chave secreta foi carregada com sucesso
	if len(secretKey) == 0 {
		fmt.Println("Erro: JWT_SECRET_KEY não encontrada no .env")
		return
	}

	fmt.Println("Chave secreta carregada com sucesso!")
	fmt.Print(secretKey)
}

func HashMD5WithSalt(password, salt string) string {
	hash := md5.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil))
  }
  
func GenerateSalt() string {
	// Usar rand.NewSource para gerar uma fonte de aleatoriedade com base no tempo
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	// Gerar um salt aleatório
	return fmt.Sprintf("%x", random.Intn(10000)) // Retorna um salt de 4 dígitos
}
