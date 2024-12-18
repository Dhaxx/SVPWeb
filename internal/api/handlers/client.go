package handlers

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/api/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ClientHandler struct {
	Repo repository.ClientRepositoryInterface
}

func NewClientHandler(repo repository.ClientRepositoryInterface) *ClientHandler {
	return &ClientHandler{Repo: repo}
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
		return
	}

	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Erro ao decodificar JSON do cliente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Repo.CreateClient(client); err != nil {
		http.Error(w, "Erro ao registrar cliente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente criado com sucesso!"})
}

func (h *ClientHandler) GetAllClients(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método Inválido!",http.StatusMethodNotAllowed)
		return
	}

	var allClients []models.Client
	allClients, err := h.Repo.GetAllClients()
	if err != nil {
		http.Error(w, "Erro ao retornar todos os clientes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allClients); err != nil {
		http.Error(w, "Erro ao codificar lista de clientes em JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ClientHandler) GetClientByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método Inválido!",http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro ao converter ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if idInt == 0 {
		http.Error(w, "ID inválido: "+id, http.StatusBadRequest)
		return
	}

	client, err := h.Repo.GetClientByID(uint(idInt))
	if err != nil {
		http.Error(w, "Erro ao obter cliente com id: "+id+", erro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(client); err != nil {
		http.Error(w, "Erro ao codificar registro de cliente em JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método Inválido", http.StatusMethodNotAllowed)
		return
	}

	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Erro ao decodificar cliente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Repo.UpdateClient(client); err != nil {
		http.Error(w, "Erro ao atualizar cliente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente Atualizado com sucesso!"})
}

func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro ao converter ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err = h.Repo.DeleteClient(uint(idInt)); err != nil {
		http.Error(w, "Erro ao deletar cliente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente deletado com sucesso!"})
}