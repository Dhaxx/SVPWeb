package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

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