package handlers

import (
    "SVPWeb/internal/api/models"
    "SVPWeb/internal/api/repository"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
)

type ImageHandler struct {
    Repo repository.ImageRepositoryInterface
}

func NewImageHandler(repo repository.ImageRepositoryInterface) *ImageHandler {
    return &ImageHandler{Repo: repo}
}

func (h *ImageHandler) CreateImage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
        return
    }

    var image models.Image
    if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
        http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Repo.CreateImage(image); err != nil {
        http.Error(w, "Erro ao criar imagem: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Imagem criada com sucesso!"})
}

func (h *ImageHandler) GetImageByID(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método Inválido!", http.StatusMethodNotAllowed)
        return
    }

    id := chi.URLParam(r, "id")
    idInt, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Erro ao converter ID: "+err.Error(), http.StatusInternalServerError)
        return
    }

    image, err := h.Repo.GetImageByID(idInt)
    if err != nil {
        http.Error(w, "Erro ao obter imagem com ID: "+id+", erro: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(image); err != nil {
        http.Error(w, "Erro ao codificar imagem em JSON: "+err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *ImageHandler) UpdateImage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Método Inválido", http.StatusMethodNotAllowed)
        return
    }

    var image models.Image
    if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
        http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
        return
    }

    if image.ID == 0 {
        http.Error(w, "ID da imagem inválido!", http.StatusBadRequest)
        return
    }

    if err := h.Repo.UpdateImage(image); err != nil {
        http.Error(w, "Erro ao atualizar imagem: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Imagem atualizada com sucesso!"})
}

func (h *ImageHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Método Inválido", http.StatusMethodNotAllowed)
        return
    }

    id := chi.URLParam(r, "id")
    idInt, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Erro ao converter ID: "+err.Error(), http.StatusInternalServerError)
        return
    }

    if err := h.Repo.DeleteImage(idInt); err != nil {
        http.Error(w, "Erro ao deletar imagem: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Imagem deletada com sucesso!"})
}