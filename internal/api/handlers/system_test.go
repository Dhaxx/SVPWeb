package handlers_test

import (
	"SVPWeb/internal/api/handlers"
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/api/repository"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"context"

	"github.com/go-chi/chi/v5"
)

func TestCreateSystem(t *testing.T) {
	mockRepo := &repository.SystemRepositoryMock{}  // Usando o mock do repositório

	// Instanciar o handler com o mock do repositório
	handler := &handlers.SystemHandler{
		Repo: mockRepo,
	}

	// Criar um usuário para testar
	sys := models.System{
		Name: "STS",
		Obs: "Sistema de Terceiro Setor",
	}

	var buf bytes.Buffer
	// Criar uma requisição simulada
	req := httptest.NewRequest(http.MethodPost, "/sistemas", &buf)
	rr := httptest.NewRecorder()

	if err := json.NewEncoder(&buf).Encode(sys); err != nil {
		t.Fatalf("erro ao codificar o sistema: %v", err)
	}

	// Chamar o handler
	handler.CreateSystem(rr, req)

	// Verificar o código de status
	if rr.Code != http.StatusCreated {
		t.Errorf("esperado: %d, obtido: %d", http.StatusCreated, rr.Code)
	}

	// Verificar o corpo da resposta
	expectedBody := `{"message":"Sistema criado com sucesso"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
	}
}

func TestGetAllSystems(t *testing.T) {
	mockRepo := &repository.SystemRepositoryMock{}  // Usando o mock do repositório

	// Instanciar o handler com o mock do repositório
	handler := &handlers.SystemHandler{
		Repo: mockRepo,
	}

	req := httptest.NewRequest(http.MethodGet, "/sistemas/", nil)
	rr := httptest.NewRecorder()

	handler.GetAllSystems(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusOK {
		t.Errorf("esperava status %d, mas obteve %d", http.StatusOK, status)
	}

	expected := `[{"ID":1,"Name":"Test System 1","Obs":""},{"ID":2,"Name":"Test System 2","Obs":""}]`
	if body := rr.Body.String(); body != expected {
		t.Errorf("esperava body %s, mas obteve %s", expected, body)
	}
}

func TestGetSystemByID(t *testing.T) {
	mockRepo := &repository.SystemRepositoryMock{}

	handler := &handlers.SystemHandler{
		Repo: mockRepo,
	}

	req := httptest.NewRequest(http.MethodGet, "/sistemas/1", nil)
	
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1") // Simula o ID "1" na rota
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()

	handler.GetSystemByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status %d, mas obteve %d", http.StatusOK, rr.Code)
	}

	// Verificar o corpo da resposta
	expectedBody := `{"ID":1,"Name":"Test System 1","Obs":""}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado body %s, mas obteve %s", expectedBody, rr.Body.String())
	}
}

func TestUpdateSystem(t *testing.T) {
	mockRepo := &repository.SystemRepositoryMock{}

	handler := &handlers.SystemHandler{
		Repo: mockRepo,
	}

	body := `{
		"Id": 1,
		"Name": "STS",
		"Obs": "Sistema de Terceiro Setor"
	}`

	req := httptest.NewRequest(http.MethodPut, "/sistemas/1", strings.NewReader(body))

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1") // Simula o ID "1" na rota
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()

	handler.UpdateSystem(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status %d, mas obteve %d", http.StatusOK, rr.Code)
	}

	// Verificar o corpo da resposta
	expectedBody := `{"ID":1,"Name":"Test System 1","Obs":""}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado body %s, mas obteve %s", expectedBody, rr.Body.String())
	}
}

func TestDeleteSystem(t *testing.T) {
	mockRepo := &repository.SystemRepositoryMock{}

	handler := &handlers.SystemHandler{
		Repo: mockRepo,
	}

	userID := 1
	req := httptest.NewRequest(http.MethodDelete, "/sistemas/1", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.Itoa(userID)) // Simula o ID "1" na rota
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()

	handler.DeleteSystem(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status %d, mas obteve %d", http.StatusOK, rr.Code)
	}

	// Verificar o corpo da resposta
	expectedBody := `{"message": "Sistema deletado com sucesso!"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado body %s, mas obteve %s", expectedBody, rr.Body.String())
	}
}