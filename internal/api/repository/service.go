package repository

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/database"
	"database/sql"
	"fmt"
)

type ServiceRepository struct {
	DB *sql.DB
}

func NewServiceRepository(*sql.DB) *ServiceRepository {
	return &ServiceRepository{DB: database.GetDB()}
}

func (cnx *ServiceRepository) CreateService(service models.Service) error {
	query := "INSERT INTO ATENDIMENTO (cliente, dtinicio, dtfim, solicitante, telefone, email, celular, inicial, desc_suporte, obs, finalizado, usuario, protocolo, sistema, usuario_alteracao, usuario_finalizacao, origem) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := cnx.DB.Exec(query, service.Client, service.StartDate, service.EndDate, service.Requester, service.Tel, service.Email, service.Cell, service.Initial, service.Description, service.Obs, service.Finished, service.User, service.Protocol, service.System, service.UserAlteration, service.UserFinished, service.Origin)
	if err != nil {
		return fmt.Errorf("erro ao inserir atendimento: %v", err)
	}

	return nil
}

func (cnx *ServiceRepository)GetAllServices() ([]models.Service, error) {
	/*2 -- fiorilli
	0 -- em andamento
	1 -- finalizado*/
	query := "SELECT id, cliente, dtinicio, dtfim, solicitante, finalizado, usuario, protocolo, inicial, desc_suporte, telefone, email, origem FROM ATENDIMENTO"

	rows, err := cnx.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar todos atendimentos: %v", err)
	}
	defer rows.Close()

	var atendimentos []models.Service
	for rows.Next() {
		var atendimento models.Service
		if err := rows.Scan(&atendimento.ID, &atendimento.Client, &atendimento.StartDate, &atendimento.EndDate, &atendimento.Requester, &atendimento.Finished, &atendimento.User, &atendimento.Protocol, &atendimento.Initial, &atendimento.Description, &atendimento.Tel, &atendimento.Email, &atendimento.Origin); err != nil {
			return nil, fmt.Errorf("erro ao scanear atendimentos: %v", err)
		}
		atendimentos = append(atendimentos, atendimento)
	}
	return atendimentos, nil
}

func (cnx *ServiceRepository) GetServiceById(id uint) (*models.Service, error) {
	query := "SELECT id, cliente, dtinicio, dtfim, solicitante, finalizado, usuario, protocolo, inicial, desc_suporte, telefone, email, origem FROM ATENDIMENTO WHERE ID = ?"

	rows, err := cnx.DB.Query(query, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("erro ao localizar atendimento com ID: %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("erro ao obter atendimento: %v", err)
	}
	defer rows.Close()

	var atendimento models.Service
	for rows.Next() {
		if err := rows.Scan(&atendimento.ID, &atendimento.Client, &atendimento.StartDate, &atendimento.EndDate, &atendimento.Requester, &atendimento.Finished, &atendimento.User, &atendimento.Protocol, &atendimento.Initial, &atendimento.Description, &atendimento.Tel, &atendimento.Email, &atendimento.Origin); err != nil {
			return nil, fmt.Errorf("erro ao scanear atendimento: %v", err)
		}
	}
	return &atendimento, nil
}

func (cnx *ServiceRepository) UpdateService(service models.Service) error {
	query := "UPDATE ATENDIMENTO SET cliente = ?, dtinicio = ?, dtfim = ?, solicitante = ?, telefone = ?, email = ?, celular = ?, inicial = ?, desc_suporte = ?, obs = ?, finalizado = ?, usuario = ?, protocolo = ?, sistema = ?, usuario_alteracao = ?, usuario_finalizacao = ?, origem = ? WHERE id = ?"

	result, err := cnx.DB.Exec(query, service.Client, service.StartDate, service.EndDate, service.Requester, service.Tel, service.Email, service.Cell, service.Initial, service.Description, service.Obs, service.Finished, service.User, service.Protocol, service.System, service.UserAlteration, service.UserFinished, service.Origin, service.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar atendimento: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao retornar linhas afetadas: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("nenhum atendimento com id: %d encontrado", service.ID)
	}

	return nil
}

func (cnx *ServiceRepository) DeleteService(id uint) error {
	query := "DELETE FROM ATENDIMENTO WHERE id = ?"

	result, err := cnx.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar atendimento: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao retornar linhas afetadas: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("nenhum atendimento com id: %d encontrado", id)
	}

	return nil
}
