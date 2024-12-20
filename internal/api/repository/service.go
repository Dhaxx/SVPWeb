package repository

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/database"
	"database/sql"
	"fmt"
)

type ServiceRepositoryInterface interface {
	CreateService(models.Service) error
	GetFilteredServices(map[string]interface{}) ([]models.Service, error)
	UpdateService(models.Service) error
	DeleteService(int) error
}

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

func (cnx *ServiceRepository) GetFilteredServices(filters map[string]interface{}) ([]models.Service, error) {
	query := `SELECT
	A.id,
	A.cliente,
	A.dtinicio,
	A.dtfim,
	A.solicitante,
	A.finalizado,
	A.usuario,
	A.protocolo,
	A.inicial,
	A.desc_suporte,
	A.telefone,
	A.email,
	A.origem,
	A.sistema,
	A.usuario_alteracao,
	A.usuario_finalizacao
FROM
	ATENDIMENTO A
JOIN CLIENTE B ON
	A.CLIENTE = B.ID
WHERE
	1 = 1`
	args := []interface{}{}

	if v, ok := filters["id"]; ok {
		query += " AND A.id = ?"
		args = append(args, v)
	}
	if v, ok := filters["entidade"]; ok {
		query += " AND entidade = ?"
		args = append(args, v)
	}
	if v, ok := filters["cidade"]; ok {
		query += " AND cidade containing ?"
		args = append(args, v)
	}
	if v, ok := filters["cliente"]; ok {
		query += " AND cliente = ?"
		args = append(args, v)
	}
	if v, ok := filters["solicitante"]; ok {
		query += " AND solicitante = ?"
		args = append(args, v)
	}
	if v, ok := filters["descricao"]; ok {
		query += " AND inicial = ?"
		args = append(args, v)
	}
	if v, ok := filters["usuario"]; ok {
		query += " AND usuario = ?"
		args = append(args, v)
	}
	if v, ok := filters["dtinicio"]; ok {
		query += " AND dtinicio >= ?"
		args = append(args, v)
	}
	if v, ok := filters["dtfim"]; ok {
		query += " AND dtfim <= ?"
		args = append(args, v)
	}

	rows, err := cnx.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar query: %v", err)
	}
	defer rows.Close()

	var allAtendimentos []models.Service
	for rows.Next() {
		var atendimento models.Service
		if err := rows.Scan(&atendimento.ID, &atendimento.Client, &atendimento.StartDate, &atendimento.EndDate, &atendimento.Requester, &atendimento.Finished, &atendimento.User, &atendimento.Protocol, &atendimento.Initial, &atendimento.Description, &atendimento.Tel, &atendimento.Email, &atendimento.Origin, &atendimento.System, &atendimento.UserAlteration, &atendimento.UserFinished); err != nil {
			return nil, fmt.Errorf("erro ao scanear atendimento: %v", err)
		}
		allAtendimentos = append(allAtendimentos, atendimento)
	}
	return allAtendimentos, nil
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

func (cnx *ServiceRepository) DeleteService(id int) error {
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