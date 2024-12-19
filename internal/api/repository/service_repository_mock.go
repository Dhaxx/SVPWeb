package repository

import (
	"SVPWeb/internal/api/models"
	"fmt"
)

// ServiceRepositoryMock simula a implementação de ServiceRepository para testes
type ServiceRepositoryMock struct{
	GetFilteredServicesFunc func(filters map[string]interface{}) ([]models.Service, error)
}

// CreateService simula a criação de um serviço
func (r *ServiceRepositoryMock) CreateService(service models.Service) error {
	// Validações simuladas
	if service.Client == 0 {
		return fmt.Errorf("cliente não pode ser vazio")
	}
	if service.Requester == "" {
		return fmt.Errorf("solicitante não pode ser vazio")
	}
	// Simula o sucesso na criação
	return nil
}

func (r *ServiceRepositoryMock) GetFilteredServices(filters map[string]interface{}) ([]models.Service, error) {
	if r.GetFilteredServicesFunc != nil {
		return r.GetFilteredServicesFunc(filters)
	}
	return nil, fmt.Errorf("GetFilteredServicesFunc not implemented")
}

// UpdateService simula a atualização de um serviço
func (r *ServiceRepositoryMock) UpdateService(service models.Service) error {
	if service.ID == 0 {
		return fmt.Errorf("serviço não encontrado")
	}
	return nil
}

// DeleteService simula a exclusão de um serviço
func (r *ServiceRepositoryMock) DeleteService(id int) error {
	if id == 0 {
		return fmt.Errorf("serviço não encontrado")
	}
	return nil
}
