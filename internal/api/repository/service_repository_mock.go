package repository

import (
	"SVPWeb/internal/api/models"
	"database/sql"
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

func (r *ServiceRepositoryMock) GetFilteredServices(filters map[string]interface{}) ([]models.Service, error) {
	// Simula alguns serviços para retorno
	services := []models.Service{
		{
			ID:         1,
			Client:     1,
			Requester:  "Fulano",
			StartDate:  sql.NullTime{Time: time.Now().Add(-48 * time.Hour), Valid: true},
			EndDate:    sql.NullTime{Time: time.Now().Add(48 * time.Hour), Valid: true},
			Finished:   1,
			User:       1,
			Protocol:   sql.NullString{String: "", Valid: false},
			Initial:    "Descrição 1",
			Description: "Descrição do atendimento 1",
		},
		{
			ID:         2,
			Client:     2,
			Requester:  "Solicitante 2",
			StartDate:  sql.NullTime{Time: time.Now().Add(-72 * time.Hour), Valid: true},
			EndDate:    sql.NullTime{Time: time.Now().Add(72 * time.Hour), Valid: true},
			Finished:   0,
			User:       2,
			Protocol:   sql.NullString{String: "Protocol 2", Valid: true},
			Initial:    "Descrição 2",
			Description: "Descrição do atendimento 2",
		},
	}

	// Simula o filtro e retorna os serviços correspondentes
	var filteredServices []models.Service
	for _, service := range services {
		if v, ok := filters["solicitante"]; ok && service.Requester == v {
			filteredServices = append(filteredServices, service)
		}
	}
	return filteredServices, nil
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
