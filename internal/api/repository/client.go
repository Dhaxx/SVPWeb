package repository

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/database"
	"database/sql"
	"fmt"
)

type ClientRepository struct {
	DB *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{DB: database.GetDB()}
}

func (cnx *ClientRepository) CreateClient(client models.Client) error {
	query := "INSERT INTO CLIENTE (ENTIDADE, CIDADE, UF, TELEFONE, EMAIL) VALUES (?, ?, ?, ?, ?)"

	_, err := cnx.DB.Exec(query, client.Entity, client.City, client.Uf, client.Tel, client.Email)
	if err != nil {
		return fmt.Errorf("erro ao criar cliente: %v", err)
	}
	return nil
}

func (cnx *ClientRepository) GetAllClients() ([]models.Client, error) {
	query := "SELECT id, entidade, cidade, uf, telefone, email FROM CLIENTE"

	rows, err := cnx.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter todos os clientes: %v", err)
	}
	defer rows.Close()

	var clientes []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Entity, &client.City, &client.Uf, &client.Tel, &client.Email); err != nil {
			return nil, fmt.Errorf("erro ao scanear valores: %v", err)
		}
		clientes = append(clientes, client)
	}
	return clientes, nil
}

func (cnx *ClientRepository) GetClientById(id uint) (*models.Client, error) {
	query := "SELECT id, entidade, cidade, uf, telefone, email FROM CLIENTE WHERE ID = ?"

	rows, err := cnx.DB.Query(query, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("erro ao localizar cliente com ID: %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar cliente: %v", err)
	}
	defer rows.Close()

	var cliente models.Client
	for rows.Next() {
		if err := rows.Scan(&cliente.ID, &cliente.Entity, &cliente.City, &cliente.Uf, &cliente.Tel, &cliente.Email); err != nil {
			return nil, fmt.Errorf("erro ao scanear valores: %v", err)
		}
	}
	return &cliente, nil
}

func (cnx *ClientRepository) UpdateClient(client models.Client) error {
	query := "UDPATE CLIENTE SET entidade = ?, cidade = ?, uf = ?, tel = ?, email = ? WHERE id = ?"

	result, err := cnx.DB.Exec(query, client.Entity, client.City, client.Uf, client.Tel, client.Email)
	if err != nil {
		return fmt.Errorf("erro ao atualizar cliente: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao retornar linhas afetadas: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("nenhum cliente com id: %d encontrado", client.ID)
	}

	return nil
}

func (cnx *ClientRepository) DeleteClient(client models.Client) error {
	query := "DELETE FROM CLIENTE WHERE ID = ?"

	result, err := cnx.DB.Exec(query, client.ID)
	if err != nil {
		return fmt.Errorf("erro ao deletar cliente: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao retornar linhas afetadas: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("nenhum cliente com id: %d encontrado", client.ID)
	}

	return nil
}