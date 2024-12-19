package handlers

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/api/repository"
	"SVPWeb/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Repo repository.UserRepositoryInterface// Utiliza da mesma conexão do repositório user
}

func NewUserHandler(repo repository.UserRepositoryInterface) *UserHandler {
	return &UserHandler{Repo: repo}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string  `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateUser(user); err != nil {
		http.Error(w, "Erro ao criar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário criado com sucesso"})
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	allUsers, err := h.Repo.GetAllUser()
	if err != nil {
		http.Error(w, "Erro ao obter todos os clientes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(allUsers); err != nil {
		http.Error(w, "Erro ao codificar usuários em JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro ao converter ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.Repo.GetUserByID(idInt)
	if err != nil {
		http.Error(w, "Erro ao obter usuario com ID: "+id+", erro : "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Erro ao codificar usuário em JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido ", http.StatusMethodNotAllowed)
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID == 0 {
		http.Error(w, "O ID do usuário é inválido", http.StatusBadRequest)
		return
	}

	err := h.Repo.UpdateUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao atualizar usuário: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário atualizado com sucesso!"})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro ao converter ID: "+id+". erro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if idInt == 0 {
		http.Error(w, "O ID do usuário é inválido", http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteUser(uint(idInt)); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao deletar usuário: %v", err), http.StatusInternalServerError)
		return
	} 

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário deletado com sucesso!"})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if r.Method != http.MethodPost {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados Inválidos", http.StatusBadRequest)
		return
	}

	user, err := service.ValidateUserCredentials(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Credenciais Inválidas", http.StatusUnauthorized)
		return
	}

	token, err := service.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}