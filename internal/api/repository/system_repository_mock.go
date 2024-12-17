package repository

import (
	"SVPWeb/internal/api/models"
	"fmt"
)

// SystemRepositoryMock simula a implementação de SystemRepository para testes
type SystemRepositoryMock struct{}

func (r *SystemRepositoryMock) CreateSystem(system models.System) error {
	// Aqui você pode verificar se os dados estão corretos, por exemplo
	if system.Name == "" {
		return fmt.Errorf("nome não pode ser vazios")
	}

	// Simulando a criação de um sistema com sucesso
	// Normalmente, o repositório real retornaria nil ou um erro se houvesse algum problema no banco de dados
	return nil
}

func (r *SystemRepositoryMock) GetAllSystems() ([]models.System, error) {
	// Simula a obtenção de todos os usuários
	return []models.System{
		{ID: 1, Name: "Test System 1", Obs: ``},
		{ID: 2, Name: "Test System 2", Obs: ``},
	}, nil
}

func (r *SystemRepositoryMock) GetSystemByID(id int) (*models.System, error) {
	// Simula o comportamento esperado ao obter um usuário pelo ID
	if id == 1 {
		return &models.System{ID: 1, Name: "Test System 1", Obs: ""}, nil
	} else if id == 2 {
		return &models.System{ID: 2, Name: "Test System 2", Obs: ""}, nil
	}
	return nil, fmt.Errorf("usuário com ID %d não encontrado", id)
}

func (r *SystemRepositoryMock) UpdateSystem(system models.System) error {
	// Simula a atualização do usuário
	if system.ID == 0 {
		return fmt.Errorf("usuário não encontrado")
	}
	return nil // Simula o sucesso na atualização
}

func (r *SystemRepositoryMock) DeleteSystem(id int) error {
	// Simula a deleção do usuário
	if id == 0 {
		return fmt.Errorf("usuário não encontrado")
	}
	return nil // Simula a exclusão com sucesso
}
