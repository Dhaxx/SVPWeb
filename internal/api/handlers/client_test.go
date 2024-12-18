package handlers_test

import (
    "SVPWeb/internal/api/handlers"
    "SVPWeb/internal/api/models"
    "SVPWeb/internal/api/repository"
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "context"

    "github.com/go-chi/chi/v5"
)

func TestCreateClient(t *testing.T) {
    mockRepo := &repository.ClientRepositoryMock{}

    handler := &handlers.ClientHandler{
        Repo: mockRepo,
    }

    client := models.Client{
        Entity: "Empresa A",
        City:   "Cidade A",
        Uf:     "UF A",
        Tel:    "12345678",
        Email:  "empresaA@example.com",
    }

    var buf bytes.Buffer
    req := httptest.NewRequest(http.MethodPost, "/clientes", &buf)
    rr := httptest.NewRecorder()

    if err := json.NewEncoder(&buf).Encode(client); err != nil {
        t.Fatalf("erro ao codificar o cliente: %v", err)
    }

    handler.CreateClient(rr, req)

    if rr.Code != http.StatusCreated {
        t.Errorf("esperado: %d, obtido: %d", http.StatusCreated, rr.Code)
    }

    expectedBody := `{"message":"Cliente criado com sucesso!"}`
    if rr.Body.String() != expectedBody {
        t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
    }
}

func TestGetAllClients(t *testing.T) {
    mockRepo := &repository.ClientRepositoryMock{}

    handler := &handlers.ClientHandler{
        Repo: mockRepo,
    }

    req := httptest.NewRequest(http.MethodGet, "/clientes", nil)
    rr := httptest.NewRecorder()

    handler.GetAllClients(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("esperado: %d, obtido: %d", http.StatusOK, rr.Code)
    }

    expectedBody := `[{"ID":1,"Entity":"Empresa A","City":"Cidade A","Uf":"UF A","Tel":"12345678","Email":"empresaA@example.com"},{"ID":2,"Entity":"Empresa B","City":"Cidade B","Uf":"UF B","Tel":"87654321","Email":"empresaB@example.com"}]`
    if rr.Body.String() != expectedBody {
        t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
    }
}

func TestGetClientByID(t *testing.T) {
    mockRepo := &repository.ClientRepositoryMock{}

    handler := &handlers.ClientHandler{
        Repo: mockRepo,
    }

    req := httptest.NewRequest(http.MethodGet, "/clientes/1", nil)
    rr := httptest.NewRecorder()

    rctx := chi.NewRouteContext()
    rctx.URLParams.Add("id", "1")
    req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

    handler.GetClientByID(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("esperado: %d, obtido: %d", http.StatusOK, rr.Code)
    }

    expectedBody := `{"ID":1,"Entity":"Empresa A","City":"Cidade A","Uf":"UF A","Tel":"12345678","Email":"empresaA@example.com"}`
    if rr.Body.String() != expectedBody {
        t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
    }
}

func TestUpdateClient(t *testing.T) {
    mockRepo := &repository.ClientRepositoryMock{}

    handler := &handlers.ClientHandler{
        Repo: mockRepo,
    }

    client := models.Client{
        ID:     1,
        Entity: "Empresa A",
        City:   "Cidade A",
        Uf:     "UF A",
        Tel:    "12345678",
        Email:  "empresaA@example.com",
    }

    var buf bytes.Buffer
    req := httptest.NewRequest(http.MethodPut, "/clientes/1", &buf)
    rr := httptest.NewRecorder()

    if err := json.NewEncoder(&buf).Encode(client); err != nil {
        t.Fatalf("erro ao codificar o cliente: %v", err)
    }

    handler.UpdateClient(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("esperado: %d, obtido: %d", http.StatusOK, rr.Code)
    }

    expectedBody := `{"message":"Cliente Atualizado com sucesso!"}`
    if rr.Body.String() != expectedBody {
        t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
    }
}

func TestDeleteClient(t *testing.T) {
    mockRepo := &repository.ClientRepositoryMock{}

    handler := &handlers.ClientHandler{
        Repo: mockRepo,
    }

    req := httptest.NewRequest(http.MethodDelete, "/clientes/1", nil)
    rr := httptest.NewRecorder()

    rctx := chi.NewRouteContext()
    rctx.URLParams.Add("id", "1")
    req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

    handler.DeleteClient(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("esperado: %d, obtido: %d", http.StatusOK, rr.Code)
    }

    expectedBody := `{"message":"Cliente deletado com sucesso!"}`
    if rr.Body.String() != expectedBody {
        t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
    }
}