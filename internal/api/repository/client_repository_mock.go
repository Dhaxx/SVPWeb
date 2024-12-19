package repository

import (
    "SVPWeb/internal/api/models"
    "fmt"
)

// ClientRepositoryMock simula a implementação de ClientRepository para testes
type ClientRepositoryMock struct{
    GetFilteredClientsFunc func(filters map[string]interface{}) ([]models.Client, error)
}

// CreateClient simula a criação de um cliente
func (r *ClientRepositoryMock) CreateClient(client models.Client) error {
    if client.Entity == "" {
        return fmt.Errorf("entidade não pode ser vazia")
    }
    return nil
}

func (r *ClientRepositoryMock) GetFilteredClients(filters map[string]interface{}) ([]models.Client, error) {
	if r.GetFilteredClientsFunc != nil {
		return r.GetFilteredClientsFunc(filters)
	}
	return nil, fmt.Errorf("GetFilteredClientsFunc not implemented")
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