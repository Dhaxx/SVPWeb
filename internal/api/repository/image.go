package repository

import (
	"SVPWeb/internal/api/models"
	"SVPWeb/internal/database"
	"database/sql"
	"fmt"
)

type ImageRepository struct {
	DB *sql.DB
}

func NewImageRepository(*sql.DB) *ImageRepository {
	return &ImageRepository{DB: database.GetDB()}
}

func (cnx *ImageRepository) CreateImage(img models.Image) error {
	query := "INSERT INTO IMAGEM (atendimento, item, imagem, usuario, descricao, data, aviso) VALUES (?, ?, ?, ?, ?, ?, ?)"

	if _, err := cnx.DB.Exec(query, img.Service, img.Item, img.Image, img.User, img.Description, img.Date, img.Notice); err != nil {
		return fmt.Errorf("erro ao inserir imagem: %v", err)
	}

	return nil
}