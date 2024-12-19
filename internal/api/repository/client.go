package repository

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/database"
	"database/sql"
	"fmt"
)

type ClientRepositoryInterface interface {
	CreateClient(models.Client) error
	GetFilteredClients(map[string]interface{}) ([]models.Client, error)
	UpdateClient(models.Client) error
	DeleteClient(uint) error
}

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

func (cnx *ClientRepository) GetFilteredClients(filters map[string]interface{}) ([]models.Client, error) {
	query := `SELECT
	id,
	entidade,
	cidade,
	uf,
	telefone,
	email
FROM
	CLIENTE
WHERE 1 = 1`
	args := []interface{}{}

	if v, ok := filters["id"]; ok {
		query += " AND id = ?"
		args = append(args, v)
	}
	if v, ok := filters["entidade"]; ok {
		query += " AND entidade containing ?"
		args = append(args, v)
	}
	if v, ok := filters["cidade"]; ok {
		query += " AND cidade containing ?"
		args = append(args, v)
	}
	if v, ok := filters["uf"]; ok {
		query += " AND uf containing ?"
		args = append(args, v)
	}
	if v, ok := filters["telefone"]; ok {
		query += " AND telefone containing ?"
		args = append(args, v)
	}
	if v, ok := filters["email"]; ok {
		query += " AND email containing ?"
		args = append(args, v)
	}

	rows, err := cnx.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter lista de clientes: %v", err.Error())
	}
	defer rows.Close()

	var allClients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Entity, &client.City, &client.Uf, &client.Tel, &client.Email); err != nil {
			return nil, fmt.Errorf("erro ao scanear clientes: %v", err)
		}
		allClients = append(allClients, client)
	}
	return allClients, nil
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

func (cnx *ClientRepository) DeleteClient(ID uint) error {
	query := "DELETE FROM CLIENTE WHERE ID = ?"

	result, err := cnx.DB.Exec(query, ID)
	if err != nil {
		return fmt.Errorf("erro ao deletar cliente: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao retornar linhas afetadas: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("nenhum cliente com id: %d encontrado", ID)
	}

	return nil
}