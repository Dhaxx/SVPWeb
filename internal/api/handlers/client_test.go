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
	"github.com/stretchr/testify/assert"
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

func TestGetFilteredClients(t *testing.T) {
	mockRepo := &repository.ClientRepositoryMock{}
	handler := &handlers.ClientHandler{
		Repo: mockRepo,
	}

	filteredClients := []models.Client{
		{
			ID:         1,
            Entity: "CAMARA DE LOS SANTOS",
            City: "LOS SANTOS",
            Uf: "CA",
            Tel: " ",
            Email: " ",
		},
	}

	// Mock do retorno do repositório
	mockRepo.GetFilteredClientsFunc = func(filters map[string]interface{}) ([]models.Client, error) {
		return filteredClients, nil
	}

	// Cria uma solicitação HTTP
	req := httptest.NewRequest(http.MethodGet, "/clients?solicitante=Solicitante+1", nil)
	rr := httptest.NewRecorder()

	// Chama o handler
	handler.GetFilteredClients(rr, req)

	// Verifica se a resposta tem o status 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verifica o conteúdo da resposta
	var response []models.Service
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, len(response), 1)
	assert.Equal(t, response[0].Requester, "Solicitante 1")
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