package main

import (
	"SVPWeb/internal/api/handlers"
	"SVPWeb/internal/api/repository"
	"SVPWeb/internal/database"
	"fmt"
	"log"
	"net/http"
	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := database.GetDB()

    // Inicializa o repositório
    userRepo := repository.NewUserRepository(db)
    systemRepo := repository.NewSystemRepository(db)
    serviceRepo := repository.NewServiceRepository(db)

    // Inicializa o handler
    userHandler := handlers.NewUserHandler(userRepo)
    systemHandler := handlers.NewSystemHandler(systemRepo)
    serviceHandler := handlers.NewServiceHandler(serviceRepo)

    // Cria o roteador
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    // Define as rotas
    r.Get("/colaboradores", userHandler.GetAllUser)
    r.Get("/sistemas", systemHandler.GetAllSystems)
    r.Get("/atendimentos", serviceHandler.GetAllServices)

    // Inicia o servidor
    fmt.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
