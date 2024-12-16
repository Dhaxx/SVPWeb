package repository

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/database"
	"database/sql"
	"fmt"
)

type UserRepositoryInterface interface {
	CreateUser(user models.User) error
	GetAllUser() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(*sql.DB) *UserRepository {
	return &UserRepository{DB: database.GetDB()}
}

func (cnx *UserRepository) CreatetUser(user models.User) error {
	query := "INSERT INTO USUARIO (nome, senha, ativo, sistema, aviso, multi, controle) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := cnx.DB.Exec(query, user.Name, user.Password, user.Active, user.System, user.Notice, user.Multi, user.Control)
	if err != nil {
		return fmt.Errorf("erro ao inserir usuário: %v", err)
	}
	return nil
}

func (cnx *UserRepository) GetAllUser() ([]models.User, error) {
	query := "SELECT id, nome, ativo, sistema, aviso, multi, controle FROM USUARIO"
	rows, err := cnx.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar todos usuários: %v", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Active, &user.System, &user.Notice, &user.Multi, &user.Control); err != nil {
			return nil, fmt.Errorf("erro ao scanear usuário: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (cnx *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := "SELECT id, nome, ativo, sistema, aviso, multi, controle FROM USUARIO WHERE ID = ?"
	rows, err := cnx.DB.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário com ID %d não encontrado", id)
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %v", err)
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Active, &user.System, &user.Notice, &user.Multi, &user.Control); err != nil {
			return nil, fmt.Errorf("erro ao scanear usuário: %v", err)
		}
	}
	return &user, nil
}

func (cnx *UserRepository) UpdateUser(user models.User) error {
	query := "UPDATE USUARIO SET nome = ?, ativo = ?, sistema = ?, multi = ?, controle = ? where id = ?"

	result, err := cnx.DB.Exec(query, user.Name, user.Active, user.System, user.Multi, user.Control, user.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas afetadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhum usuário com ID %d encontrado", user.ID)
	}

	return nil
}

func (cnx *UserRepository) DeleteUser(user models.User) error {
	query := "DELETE FROM USUARIO WHERE ID = ?"

	result, err := cnx.DB.Exec(query, user.ID)
	if err != nil {
		return fmt.Errorf("erro ao apagar usuários: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas deletadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhum usuário com ID %d encontrado", user.ID)
	}

	return nil
}