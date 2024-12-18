package handlers_test

import (
	"SVPWeb/internal/api/handlers"
	"SVPWeb/internal/api/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

func TestGetAllServices(t *testing.T) {
	mockRepo := &repository.ServiceRepositoryMock{}

	handler := &handlers.ServiceHandler{
		Repo: mockRepo,
	}

	req := httptest.NewRequest(http.MethodGet, "/atendimentos/all", nil)
	rr := httptest.NewRecorder()

	handler.GetAllServices(rr, req)
	
	if rr.Code != http.StatusOK {
		t.Errorf("esperado: %d, obtido: %d", http.StatusCreated, rr.Code)	
	}

	expectedBody := `[{ "Client": 321, "StartDate": time.Now(), "EndDate": time.Now().Add(48 * time.Hour), "Requester": "Fulaninho", "Tel": "12345678", "Email": "johndoe@example.com", "Cell": "987654321", "Initial": "ERRO AO GERAR FASE 4", "Description": "SISTEMA ESTAVA DESATUALIZADO", "Obs": "", "Finished": 1, "User": 22, "Protocol": "", "System": 2, "UserAlteration": 0, "UserFinished": 0, "Origin": 0}]`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
	}
}

func TestGetServiceByID(t *testing.T) {
	mockRepo := &repository.ServiceRepositoryMock{}
	handler := &handlers.ServiceHandler{
		Repo: mockRepo,
	}

	req := httptest.NewRequest(http.MethodGet, "/atendimentos/1", nil)
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.GetServiceByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado: %d, obtido: %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{"Client": 321, "StartDate": time.Now(), "EndDate": time.Now().Add(48 * time.Hour), "Requester": "Fulaninho", "Tel": "12345678", "Email": "johndoe@example.com", "Cell": "987654321", "Initial": "ERRO AO GERAR FASE 4", "Description": "SISTEMA ESTAVA DESATUALIZADO", "Obs": "", "Finished": 1, "User": 22, "Protocol": "", "System": 2, "UserAlteration": 0, "UserFinished": 0, "Origin": 0}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
	}
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
