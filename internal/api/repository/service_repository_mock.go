package repository

import (
	"SVPWeb/internal/api/models"
	"fmt"
	"time"
)

// ServiceRepositoryMock simula a implementação de ServiceRepository para testes
type ServiceRepositoryMock struct{}

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

// GetAllServices simula a obtenção de todos os serviços
func (r *ServiceRepositoryMock) GetAllServices() ([]models.Service, error) {
	return []models.Service{
		{
			ID:          1,
			Client:      321,
			StartDate:   time.Now(),
			EndDate:     time.Now().Add(48 * time.Hour),
			Requester:   "Fulaninho",
			Tel:         "12345678",
			Email:       "johndoe@example.com",
			Cell:        "987654321",
			Initial:     "ERRO AO GERAR FASE 4",
			Description: "SISTEMA ESTAVA DESATUALIZADO",
			Obs:         "",
			Finished:    1,
			User:        22,
			Protocol:    "",
			System:      2,
			UserAlteration: 0,
			UserFinished:   0,
			Origin:         0,
		},
		{
			ID:          2,
			Client:      18,
			StartDate:   time.Now(),
			EndDate:     time.Now().Add(24 * time.Hour),
			Requester:   "Ciclano",
			Tel:         "87654321",
			Email:       "janedoe@example.com",
			Cell:        "987654321",
			Initial:     "DÚVIDA DE LICITAÇÃO",
			Description: "",
			Obs:         "",
			Finished:    1,
			User:        22,
			Protocol:    "",
			System:      2,
			UserAlteration: 0,
			UserFinished:   0,
			Origin:         0,
		},
	}, nil
}

// GetServiceByID simula a obtenção de um serviço pelo ID
func (r *ServiceRepositoryMock) GetServiceByID(id int) (*models.Service, error) {
	if id == 1 {
		return &models.Service{
			ID:          1,
			Client:      321,
			StartDate:   time.Now(),
			EndDate:     time.Now().Add(48 * time.Hour),
			Requester:   "Fulaninho",
			Tel:         "12345678",
			Email:       "johndoe@example.com",
			Cell:        "987654321",
			Initial:     "ERRO AO GERAR FASE 4",
			Description: "SISTEMA ESTAVA DESATUALIZADO",
			Obs:         "",
			Finished:    1,
			User:        22,
			Protocol:    "",
			System:      2,
			UserAlteration: 0,
			UserFinished:   0,
			Origin:         0,
		}, nil
	}
	if id == 2 {
		return &models.Service{
			ID:          2,
			Client:      18,
			StartDate:   time.Now(),
			EndDate:     time.Now().Add(24 * time.Hour),
			Requester:   "Ciclano",
			Tel:         "87654321",
			Email:       "janedoe@example.com",
			Cell:        "987654321",
			Initial:     "DÚVIDA DE LICITAÇÃO",
			Description: "",
			Obs:         "",
			Finished:    1,
			User:        22,
			Protocol:    "",
			System:      2,
			UserAlteration: 0,
			UserFinished:   0,
			Origin:         0,
		}, nil
	}
	return nil, fmt.Errorf("serviço com ID %d não encontrado", id)
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
