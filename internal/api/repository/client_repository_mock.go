package repository

import (
    "SVPWeb/internal/api/models"
    "fmt"
)

// ClientRepositoryMock simula a implementação de ClientRepository para testes
type ClientRepositoryMock struct{}

// CreateClient simula a criação de um cliente
func (r *ClientRepositoryMock) CreateClient(client models.Client) error {
    if client.Entity == "" {
        return fmt.Errorf("entidade não pode ser vazia")
    }
    return nil
}

// GetAllClients simula a obtenção de todos os clientes
func (r *ClientRepositoryMock) GetAllClients() ([]models.Client, error) {
    return []models.Client{
        {ID: 1, Entity: "PREFEITURA DE LOS SANTOS", City: "LOS SANTOS", Uf: "SP", Tel: "12345678", Email: "empresaA@example.com"},
        {ID: 2, Entity: "CÂMARA MUNICIPAL DE GOTHAM CITY", City: "GOTHAM", Uf: "MG", Tel: "87654321", Email: "empresaB@example.com"},
    }, nil
}

// GetClientById simula a obtenção de um cliente pelo ID
func (r *ClientRepositoryMock) GetClientByID(id uint) (*models.Client, error) {
    if id == 1 {
        return &models.Client{ID: 1, Entity: "PREFEITURA DE LOS SANTOS", City: "LOS SANTOS", Uf: "SP A", Tel: "12345678", Email: "empresaA@example.com"}, nil
    }
    return nil, fmt.Errorf("cliente com ID %d não encontrado", id)
}

// UpdateClient simula a atualização de um cliente
func (r *ClientRepositoryMock) UpdateClient(client models.Client) error {
    if client.ID == 0 {
        return fmt.Errorf("cliente não encontrado")
    }
    return nil
}

// DeleteClient simula a exclusão de um cliente
func (r *ClientRepositoryMock) DeleteClient(id uint) error {
    if id == 0 {
        return fmt.Errorf("cliente não encontrado")
    }
    return nil
}