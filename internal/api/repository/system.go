package repository

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/database"
	"database/sql"
	"fmt"
)

type SystemRepositoryInterface interface {
	CreateSystem(models.System) error
	GetAllSystems() ([]models.System, error)
	GetSystemByID(id int) (*models.System, error)
	UpdateSystem(models.System) error
	DeleteSystem(id int) error
}

type SystemRepository struct {
	DB *sql.DB
}

func NewSystemRepository(*sql.DB) *SystemRepository {
	return &SystemRepository{DB: database.GetDB()}
}

func (cnx *SystemRepository) CreateSystem(sys models.System) error {
	query := "INSERT INTO SISTEMA (nome, obs) VALUES ?, ?"
	_, err := cnx.DB.Exec(query, sys.Name, sys.Obs)
	if err != nil {
		return fmt.Errorf("erro ao inserir tipo de sistema %v", err)
	}
	return nil
}

func (cnx *SystemRepository) GetAllSystems() ([]models.System, error) {
	query := "SELECT ID, NOME, OBS FROM SISTEMA"
	rows, err := cnx.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar todos os sistemas: %v", err)
	}
	defer rows.Close()

	var systems []models.System
	for rows.Next() {
		var system models.System
		if err := rows.Scan(&system.ID, &system.Name, &system.Obs); err != nil {
			return nil, fmt.Errorf("erro ao scanear usuários: %v", err)
		}
		systems = append(systems, system)
	}
	return systems, nil
}

func (cnx *SystemRepository) GetSystemByID(id int) (*models.System, error) {
	query := "SELECT ID, NOME, OBS FROM SISTEMA WHERE ID = ?"
	rows, err := cnx.DB.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("erro ao buscar o sistema com ID: %d", id)
		}
		return nil, fmt.Errorf("erro ao buscar sistemas: %v", err)
	}
	defer rows.Close()

	var system models.System
	for rows.Next() {
		if err := rows.Scan(&system.ID, &system.Name, &system.Obs); err != nil {
			return nil, fmt.Errorf("erro ao scanear sistemas: %v", err)
		}
	}
	return &system, nil
}

func (cnx *SystemRepository) UpdateSystem(sistema models.System) error {
	query := "UPDATE SISTEMA SET NOME = ?, OBS = ? WHERE ID = ?"

	result, err := cnx.DB.Exec(query, sistema.Name, sistema.Obs, sistema.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar sistema: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao retornar linhas afetadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhum sistema com id: %d encontrado", sistema.ID)
	}

	return nil
}

func (cnx *SystemRepository) DeleteSystem(ID int) error {
	query := "DELETE FROM SISTEMA WHERE ID = ?"

	result, err := cnx.DB.Exec(query, ID)
	if err != nil {
		return fmt.Errorf("erro ao executar delete de sistema")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao retornar linhas afetadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhum sistema com id: %d localizado", ID)
	}

	return nil
}
