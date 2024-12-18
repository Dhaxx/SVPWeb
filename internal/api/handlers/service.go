package handlers

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/api/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5")

type ServiceHandler struct {
	Repo repository.ServiceRepositoryInterface
}

func NewServiceHandler(repo repository.ServiceRepositoryInterface) *ServiceHandler {
	return &ServiceHandler{Repo: repo}
}

func (h *ServiceHandler) CreateService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
		return
	}

	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Erro ao decodificar JSON da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateService(service); err != nil {
		http.Error(w, "Erro ao criar atendimento: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Atendimento registrado com sucesso!"})
}

func (h *ServiceHandler) GetFilteredServices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método Inválido!", http.StatusInternalServerError)
		return
	}

	filters := r.URL.Query()
	filterMap := make(map[string]interface{})
	for key, values := range filters {
		if len(values) > 0 {
			filterMap[key] = values[0]
		}
	}
	
	services, err := h.Repo.GetFilteredServices(filterMap)
	if err != nil {
		http.Error(w, "Erro ao buscar atendimentos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func (h *ServiceHandler) UpdateSystem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método Inválido", http.StatusInternalServerError)
		return
	}

	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusInternalServerError)
		return	
	}

	if err := h.Repo.UpdateService(service); err != nil {
		http.Error(w, "Erro ao atualizar atendimento: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Atendimento atualizado com sucesso!"})
}

func (h *ServiceHandler) DeleteSystem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método Inválido", http.StatusInternalServerError)
		return
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro ao converter ID: "+id+", erro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Repo.DeleteService(idInt); err != nil {
		http.Error(w, "Erro ao deletar atendimento: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Atendimento excluído com sucesso!"})
}