package repository

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/database"
	"database/sql"
	"fmt"
)

type NoticeRepositoryInterface interface {
	CreateNotice(models.Notice) error 
	GetAllNotices() ([]models.Notice, error)
	GetNoticeByID(uint) (*models.Notice, error)
	UpdateNotice(models.Notice) error
	DeleteNotice(uint) error
}

type NoticeRepository struct {
	DB *sql.DB
}

func NewNoticeRepository(db *sql.DB) *NoticeRepository {
	return &NoticeRepository{DB: database.GetDB()}
}

func (cnx *NoticeRepository) CreateNotice(notice models.Notice) error {
	query := "INSERT INTO AVISOS (id, titulo, sistema, usuario, data, tipo, caminho) values (?,?,?,?,?,?,?)"

	if _, err := cnx.DB.Exec(query, notice.Title, notice.System, notice.User, notice.Date, notice.Type, notice.Path); err != nil {
		return fmt.Errorf("erro ao inserir aviso: %v", err)
	}

	return nil
}

func (cnx *NoticeRepository) GetAllNotices() ([]models.Notice, error) {
	query := "SELECT id, titulo, sistema, usuario, data, tipo, caminho FROM AVISOS"

	rows, err := cnx.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter todos avisos")
	}
	defer rows.Close()

	var avisos []models.Notice
	for rows.Next() {
		var aviso models.Notice
		if err := rows.Scan(&aviso.ID, &aviso.Title, &aviso.System, &aviso.User, &aviso.Date, &aviso.Type, &aviso.Path); err != nil {
			return nil, fmt.Errorf("erro ao scanear avisos: %v", err)
		}
		avisos = append(avisos, aviso)
	}

	return avisos, nil
}

func (cnx *NoticeRepository) GetNoticeByID(id uint) (*models.Notice, error) {
	query := "SELECT id, titulo, sistema, usuario, data, tipo, caminho FROM AVISOS WHERE ID = ?"

	rows, err := cnx.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar avisos: %v", err)
	}
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("erro ao localizar aviso com ID: %d", id)
	}
	defer rows.Close()

	var notice models.Notice
	for rows.Next() {
		err := rows.Scan(&notice.ID, &notice.Title, &notice.System, &notice.User, &notice.Date, &notice.Type, &notice.Path)
		if err != nil {
			return nil, fmt.Errorf("erro ao scanear aviso: %v", err)
		}
	}
	return &notice, nil
}

func (cnx *NoticeRepository) UpdateNotice(notice models.Notice) error {
	query := "UPDATE AVISOS titulo = ?, sistema = ?, usuario = ?, tipo = ? WHERE id = ?"

	result, err := cnx.DB.Exec(query, notice.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar aviso: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("nenhum aviso com ID: %d encontrado", notice.ID)
	}

	return nil
}

func (cnx *NoticeRepository) DeleteNotice(id uint) error {
	query := "DELETE FROM AVISOS WHERE ID = ?"

	result, err := cnx.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar aviso: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("nenhum aviso com ID: %d encontrado", id)
	}

	return nil
}

func (cnx *NoticeRepository) CreateNoticeType(noticeType models.NoticeType) error {
	query := "INSERT INTO TIPO_AVISO (nome) values ?, ?"

	_, err := cnx.DB.Exec(query, noticeType.Name)
	if err != nil {
		return fmt.Errorf("erro ao criar tipo de aviso: %v", err)
	}

	return nil
}

func (cnx *NoticeRepository) GetAllNoticeType() ([]models.NoticeType, error) {
	query := "SELECT id, nome FROM TIPO_AVISO"

	rows, err := cnx.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro obter todos os tipos de aviso: %v", err)
	}
	defer rows.Close()

	var tiposAviso []models.NoticeType
	for rows.Next() {
		var tipoAviso models.NoticeType
		err := rows.Scan(&tipoAviso.ID, &tipoAviso.Name)
		if err != nil {
			return nil, fmt.Errorf("erro ao scanear tipos de aviso: %v", err)
		}
		tiposAviso = append(tiposAviso, tipoAviso)
	}
	return tiposAviso, nil
}

func (cnx *NoticeRepository) GetNoticeTypeById(id uint) (*models.NoticeType, error) {
	query := "SELECT id, nome FROM TIPO_AVISO WHERE ID = ?"

	row := cnx.DB.QueryRow(query, id)
	var noticeType models.NoticeType
	err := row.Scan(&noticeType.ID, &noticeType.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nenhum tipo de aviso com ID: %d encontrado", id)
		}
		return nil, fmt.Errorf("erro ao buscar tipo de aviso: %v", err)
	}

	return &noticeType, nil
}

func (cnx *NoticeRepository) UpdateNoticeType(noticeType models.NoticeType) error {
	query := "UPDATE TIPO_AVISO SET nome = ? WHERE id = ?"

	result, err := cnx.DB.Exec(query, noticeType.Name, noticeType.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar tipo de aviso: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("nenhum tipo de aviso com ID: %d encontrado", noticeType.ID)
	}

	return nil
}

func (cnx *NoticeRepository) DeleteNoticeType(id uint) error {
	query := "DELETE FROM TIPO_AVISO WHERE ID = ?"

	result, err := cnx.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar tipo de aviso: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("nenhum tipo de aviso com ID: %d encontrado", id)
	}

	return nil
}