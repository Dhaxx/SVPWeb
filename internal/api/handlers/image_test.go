package handlers

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/api/repository"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateImage(t *testing.T) {
	repo := &repository.ImageRepositoryMock{}
	handler := NewImageHandler(repo)

	image := models.Image{ID: 1, Description: "Test Image"}
	imageJSON, _ := json.Marshal(image)
	req, err := http.NewRequest(http.MethodPost, "/images", bytes.NewBuffer(imageJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.CreateImage(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "Imagem criada com sucesso!")
}

func TestGetImageByID(t *testing.T) {
	repo := &repository.ImageRepositoryMock{}
	handler := NewImageHandler(repo)

	r := chi.NewRouter()
	r.Get("/images/{id}", handler.GetImageByID)

	req, err := http.NewRequest(http.MethodGet, "/images/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Test Image")
}

func TestUpdateImage(t *testing.T) {
	repo := &repository.ImageRepositoryMock{}
	handler := NewImageHandler(repo)

	image := models.Image{ID: 1, Description: "Updated Test Image"}
	imageJSON, _ := json.Marshal(image)
	req, err := http.NewRequest(http.MethodPut, "/images", bytes.NewBuffer(imageJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.UpdateImage(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Imagem atualizada com sucesso!")
}

func TestDeleteImage(t *testing.T) {
	repo := &repository.ImageRepositoryMock{}
	handler := NewImageHandler(repo)

	r := chi.NewRouter()
	r.Delete("/images/{id}", handler.DeleteImage)

	req, err := http.NewRequest(http.MethodDelete, "/images/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Imagem deletada com sucesso!")
}
