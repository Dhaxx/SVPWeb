package main

import (
	"SVPWeb/internal/api/handlers"
	"SVPWeb/internal/api/repository"
	"SVPWeb/internal/database"
	"SVPWeb/internal/service"
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
    clientRepo := repository.NewClientRepository(db)
    imageRepo := repository.NewImageRepository(db)
    noticeRepo := repository.NewNoticeRepository(db)

    // Inicializa o handler
    userHandler := handlers.NewUserHandler(userRepo)
    systemHandler := handlers.NewSystemHandler(systemRepo)
    serviceHandler := handlers.NewServiceHandler(serviceRepo)
    clientHandler := handlers.NewClientHandler(clientRepo)
    imageHandler := handlers.NewImageHandler(imageRepo)
    noticeHandler := handlers.NewNoticeRepository(noticeRepo)

    // Cria o roteador
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    // Rota de Autenticação
    r.Post("/SVPWeb/login", userHandler.Login)

    // Define as rotas
    r.Route("/SVPWeb", func(r chi.Router) {
		r.Use(service.JWTAuthMiddleware) // Middleware para verificar o token JWT

		// Rotas GET
		r.Get("/colaboradores", userHandler.GetAllUser)
		r.Get("/sistemas", systemHandler.GetAllSystems)
		r.Get("/atendimentos", serviceHandler.GetFilteredServices)
		r.Get("/clientes", clientHandler.GetFilteredClients)
		r.Get("/image/{id}", imageHandler.GetImageByID)
		r.Get("/notice", noticeHandler.GetAllNotices)
	})


    // Inicia o servidor
    fmt.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}