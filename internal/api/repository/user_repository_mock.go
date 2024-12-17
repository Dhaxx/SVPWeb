package repository

import (
	"SVPWeb/internal/api/models"
	"fmt"
)

// UserRepositoryMock simula a implementação de UserRepository para testes
type UserRepositoryMock struct{}

func (r *UserRepositoryMock) CreateUser(user models.User) error {
	// Aqui você pode verificar se os dados estão corretos, por exemplo
	if user.Name == "" || user.Password == "" {
		return fmt.Errorf("nome ou senha não podem ser vazios")
	}

	// Simulando a criação de um usuário com sucesso
	// Normalmente, o repositório real retornaria nil ou um erro se houvesse algum problema no banco de dados
	return nil
}

func (r *UserRepositoryMock) GetAllUser() ([]models.User, error) {
	// Simula a obtenção de todos os usuários
	return []models.User{
		{ID: 1, Name: "Test User 1", Active: `S`},
		{ID: 2, Name: "Test User 2", Active: `N`},
	}, nil
}

func (r *UserRepositoryMock) GetUserByID(id int) (*models.User, error) {
	// Simula o comportamento esperado ao obter um usuário pelo ID
	if id == 1 {
		return &models.User{ID: 1, Name: "Test User 1", Active: "S"}, nil
	} else if id == 2 {
		return &models.User{ID: 2, Name: "Test User 2", Active: "N"}, nil
	}
	return nil, fmt.Errorf("usuário com ID %d não encontrado", id)
}

func (r *UserRepositoryMock) UpdateUser(user models.User) error {
	// Simula a atualização do usuário
	if user.ID == 0 {
		return fmt.Errorf("usuário não encontrado")
	}
	return nil // Simula o sucesso na atualização
}

func (r *UserRepositoryMock) DeleteUser(id uint) error {
	// Simula a deleção do usuário
	if id == 0 {
		return fmt.Errorf("usuário não encontrado")
	}
	return nil // Simula a exclusão com sucesso
}
