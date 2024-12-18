package handlers

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/api/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type SystemHandler struct {
	Repo repository.SystemRepositoryInterface
}

func NewSystemHandler(repo repository.SystemRepositoryInterface) *SystemHandler {
	return &SystemHandler{Repo: repo}
}

func (h *SystemHandler) CreateSystem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
		return
	}

	var system models.System
	if err := json.NewDecoder(r.Body).Decode(&system); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateSystem(system); err != nil {
		http.Error(w, "Erro ao cadastrar sistema: "+err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sistema cadastrado com sucesso!"})
}

func (h *SystemHandler) GetAllSystems(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
		return
	}

	allSystems, err := h.Repo.GetAllSystems()
	if err != nil {
		http.Error(w, "Erro ao obter lista de sistemas: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allSystems); err != nil {
		http.Error(w, "Erro ao codificar sistemas em formato JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SystemHandler) GetSystemByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método Inválido! ", http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro ao converter ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	system, err := h.Repo.GetSystemByID(idInt)
	if err != nil {
		http.Error(w, "Erro ao obter sistema com ID: "+id+", erro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(system); err != nil {
		http.Error(w, "Erro ao codificar sistema em JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SystemHandler) UpdateSystem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método Inválido", http.StatusMethodNotAllowed)
		return
	}

	var system models.System
	if err := json.NewDecoder(r.Body).Decode(&system); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if system.ID == 0 {
		http.Error(w, "ID do sistema inválido!", http.StatusBadRequest)
		return
	}

	if err := h.Repo.UpdateSystem(system); err != nil {
		http.Error(w, "Erro ao atualizar cadastro de sistema: "+err.Error(), http.StatusInternalServerError)
		return
	}	

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sistema atualizado com sucesso!"})
}

func (h *SystemHandler) DeleteSystem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método Inválido", http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro ao converter ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if intId == 0 {
		http.Error(w, "O ID do usuário é inválido", http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteSystem(intId); err != nil {
		http.Error(w, "Erro ao deletar registro de sistema: "+err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registro de sistema apagado com sucesso!"})
}