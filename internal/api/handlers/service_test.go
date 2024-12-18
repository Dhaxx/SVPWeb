package handlers_test

import (
	"SVPWeb/internal/api/handlers"
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/api/repository"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"context"

	"github.com/go-chi/chi/v5"
)

func TestCreateService(t *testing.T) {
	mockRepo := &repository.ServiceRepositoryMock{}

	handler := &handlers.ServiceHandler{
		Repo: mockRepo,
	}

	body := `{
		"Client":      321,
		"StartDate": "2024-12-17T15:04:05-03:00",
		"EndDate": "2024-12-18T15:04:05-03:00",
		"Requester":   "Fulaninho",
		"Tel":         "12345678",
		"Email":       "johndoe@example.com",
		"Cell":        "987654321",
		"Initial":     "ERRO AO GERAR FASE 4",
		"Description": "SISTEMA ESTAVA DESATUALIZADO",
		"Obs":         "",
		"Finished":    1,
		"User":        22,
		"Protocol":    "",
		"System":      2,
		"UserAlteration": 0,
		"UserFinished":   0,
		"Origin":         0
	}`

	req := httptest.NewRequest(http.MethodPost, "/atendimentos/", strings.NewReader(body))
	rr := httptest.NewRecorder()

	handler.CreateService(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("esperado: %d, obtido: %d", http.StatusCreated, rr.Code)	
	}

	expectedBody := `{"message": "Atendimento registrado com sucesso!"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
	}
}

func TestGetFilteredServices(t *testing.T) {
	mockRepo := &repository.ServiceRepositoryMock{}
	handler := &handlers.ServiceHandler{
		Repo: mockRepo,
	}

	// Cria um filtro de exemplo
	filters := map[string]interface{}{"solicitante": "Solicitante 1"}
	filteredServices := []models.Service{
		{
			ID:         1,
			Client:     1,
			Requester:  "Solicitante 1",
			StartDate:  sql.NullTime{Time: time.Now().Add(-48 * time.Hour), Valid: true},
			EndDate:    sql.NullTime{Time: time.Now().Add(48 * time.Hour), Valid: true},
			Finished:   1,
			User:       1,
			Protocol:   sql.NullString{String: "", Valid: false},
			Initial:    "Descrição 1",
			Description: "Descrição do atendimento 1",
		},
	}

	// Mock para simular o retorno do repositório
	handler.GetFilteredServices()
	mockRepo.GetFilteredServicesFunc = func(filters map[string]interface{}) ([]models.Service, error) {
		return filteredServices, nil
	}

	// Cria uma solicitação HTTP
	req := httptest.NewRequest(http.MethodGet, "/services?solicitante=Solicitante+1", nil)
	rr := httptest.NewRecorder()

	// Chama o handler
	handler.GetFilteredServices(rr, req)

	// Verifica se a resposta tem o status 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verifica o conteúdo da resposta
	var response []models.Service
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, len(response), 1)
	assert.Equal(t, response[0].Requester, "Solicitante 1")
}

func TestUpdateService(t *testing.T) {
	mockRepo := &repository.ServiceRepositoryMock{}
	handler := &handlers.ServiceHandler{
		Repo: mockRepo,
	}

	body := `{
		"Client":      321,
		"StartDate": "2024-12-17T15:04:05-03:00",
		"EndDate": "2024-12-18T15:04:05-03:00",
		"Requester":   "Fulaninho",
		"Tel":         "12345678",
		"Email":       "johndoe@example.com",
		"Cell":        "987654321",
		"Initial":     "ERRO AO GERAR FASE 4",
		"Description": "SISTEMA ESTAVA DESATUALIZADO",
		"Obs":         "",
		"Finished":    1,
		"User":        22,
		"Protocol":    "",
		"System":      2,
		"UserAlteration": 0,
		"UserFinished":   0,
		"Origin":         0
	}`

	req := httptest.NewRequest(http.MethodPost, "/atendimentos/update", strings.NewReader(body))
	rr := httptest.NewRecorder()

	handler.UpdateSystem(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado: %d, obtido: %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{"message": "Atendimento atualizado com sucesso!"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
	}
}

func TestDeleteService(t *testing.T) {
	mockRepo := &repository.ServiceRepositoryMock{}
	handler := &handlers.ServiceHandler{
		Repo: mockRepo,
	}

	req := httptest.NewRequest(http.MethodPost, "/atendimentos/delete/1", nil)
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.DeleteSystem(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado: %d, obtido: %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{"message": "Atendimento excluído com sucesso!"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
	}
}
