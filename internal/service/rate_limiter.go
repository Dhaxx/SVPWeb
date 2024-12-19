package service

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// Limiter é o rate limiter global
var Limiter = rate.NewLimiter(rate.Every(1*time.Second), 15) // Limita a 5 requisições por segundo

// RateLimitingMiddleware aplica o rate limiting na API
func RateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verifica se o cliente está dentro do limite de requisições
		if !Limiter.Allow() {
			// Se o limite de requisições for excedido, retorna um erro 429 (Too Many Requests)
			http.Error(w, "Too many requests, please try again later.", http.StatusTooManyRequests)
			return
		}
		// Se estiver dentro do limite, continua com a requisição
		next.ServeHTTP(w, r)
	})
}
