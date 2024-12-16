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

type UserRepositoryInterface interface {
	CreateUser(user models.User) error
	GetAllUser() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
}

func TestCreateUser(t *testing.T) {
	mockRepo := &repository.UserRepositoryMock{}  // Usando o mock do repositório

	// Instanciar o handler com o mock do repositório
	handler := &handlers.UserHandler{
		Repo: mockRepo,
	}

	// Criar um usuário para testar
	user := models.User{
		Name:     "John Doe",
		Password: "password123",
		Active:   '1',
		System:   1,
		Notice:   0,
		Multi:    '1',
		Control:  1,
		PassMD5: "somehashedpassword",
		Cad:      '1',
	}

	var buf bytes.Buffer
	// Criar uma requisição simulada
	req := httptest.NewRequest(http.MethodPost, "/usuarios", &buf)
	rr := httptest.NewRecorder()

	if err := json.NewEncoder(&buf).Encode(user); err != nil {
		t.Fatalf("erro ao codificar o usuário: %v", err)
	}

	// Chamar o handler
	handler.CreateUser(rr, req)

	// Verificar o código de status
	if rr.Code != http.StatusCreated {
		t.Errorf("esperado: %d, obtido: %d", http.StatusCreated, rr.Code)
	}

	// Verificar o corpo da resposta
	expectedBody := `{"message":"Usuário criado com sucesso"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado: %s, obtido: %s", expectedBody, rr.Body.String())
	}
}

func TestGetUserByID(t *testing.T) {
	// Instanciar o mock do repositório
	mockRepo := &repository.UserRepositoryMock{}

	// Instanciar o handler com o mock do repositório
	handler := &handlers.UserHandler{
		Repo: mockRepo,
	}

	// Criar uma requisição simulada para um ID existente
	req := httptest.NewRequest(http.MethodGet, "/usuarios/1", nil)

	// Usar chi para injetar o parâmetro "id" na rota
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1") // Simula o ID "1" na rota
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Criar um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()

	// Chamar o handler
	handler.GetUserByID(rr, req)

	// Verificar o status HTTP retornado
	if rr.Code != http.StatusOK {
		t.Errorf("esperado status %d, mas obteve %d", http.StatusOK, rr.Code)
	}

	// Verificar o corpo da resposta
	expectedBody := `{"ID":1,"Name":"Test User 1","Active":true}`
	if rr.Body.String() != expectedBody {
		t.Errorf("esperado body %s, mas obteve %s", expectedBody, rr.Body.String())
	}
}
