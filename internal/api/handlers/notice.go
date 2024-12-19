package handlers

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/api/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type NoticeRepository struct {
	Repo repository.NoticeRepositoryInterface
}

func NewNoticeRepository(repo repository.NoticeRepositoryInterface) *NoticeRepository {
	return &NoticeRepository{Repo: repo}
}

func (h *NoticeRepository) CreateNotice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
		return
	}

	var notice  models.Notice
	if err := json.NewDecoder(r.Body).Decode(&notice); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Repo.CreateNotice(notice); err != nil {
		http.Error(w, "Erro ao criar aviso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Aviso criado com sucesso!"})
}

func (h* NoticeRepository) GetAllNotices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
		return
	}

	allNotices, err := h.Repo.GetAllNotices()
	if err != nil {
		http.Error(w, "Erro ao obter todos avisos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allNotices); err != nil {
		http.Error(w, "Erro ao codificar todas as notícias: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NoticeRepository) GetNoticeByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	notice, err := h.Repo.GetNoticeByID(uint(id))
	if err != nil {
		http.Error(w, "Erro ao obter aviso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notice); err != nil {
		http.Error(w, "Erro ao codificar aviso: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NoticeRepository) UpdateNotice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	var notice models.Notice
	if err := json.NewDecoder(r.Body).Decode(&notice); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	notice.ID = id
	if err := h.Repo.UpdateNotice(notice); err != nil {
		http.Error(w, "Erro ao atualizar aviso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Aviso atualizado com sucesso!"})
}

func (h *NoticeRepository) DeleteNotice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteNotice(uint(id)); err != nil {
		http.Error(w, "Erro ao deletar aviso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Aviso deletado com sucesso!"})
}