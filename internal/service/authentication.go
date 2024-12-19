package service

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/api/repository"
	"SVPWeb/internal/database"
	"SVPWeb/internal/utils"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
}

func ValidateUserCredentials(username, password string) (*models.User, error) {
	userRepo := repository.NewUserRepository(database.GetDB())
	user, err := userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter usuário: %v", err)
	}

	passwordHashed := []byte(utils.HashMD5(password))
	passwordStored := []byte(user.Password)

	if !bytes.Equal(passwordHashed, passwordStored) {
		return nil, fmt.Errorf("senhas não coincidem: %v", err)
	}

	return user, nil
}

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id": user.ID,
		"username": user.Name,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Print(err)
		return "", fmt.Errorf("erro ao assinar string: %v", err)
	}

	return signedToken, nil
}
