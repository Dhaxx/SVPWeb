package repository

import (
    "SVPWeb/internal/api/models"
    "SVPWeb/internal/database"
    "database/sql"
    "fmt"
)

type ImageRepositoryInterface interface {
    CreateImage(models.Image) error
    GetImageByID(id int) (*models.Image, error)
    UpdateImage(models.Image) error
    DeleteImage(id int) error
}

type ImageRepository struct {
    DB *sql.DB
}

func NewImageRepository(db *sql.DB) *ImageRepository {
    return &ImageRepository{DB: database.GetDB()}
}

func (cnx *ImageRepository) CreateImage(img models.Image) error {
    query := "INSERT INTO IMAGEM (atendimento, item, imagem, usuario, descricao, data, aviso) VALUES (?, ?, ?, ?, ?, ?, ?)"

    if _, err := cnx.DB.Exec(query, img.Service, img.Item, img.Image, img.User, img.Description, img.Date, img.Notice); err != nil {
        return fmt.Errorf("erro ao inserir imagem: %v", err)
    }

    return nil
}

func (cnx *ImageRepository) GetImageByID(id int) (*models.Image, error) {
    query := "SELECT id, atendimento, item, imagem, usuario, descricao, data, aviso FROM IMAGEM WHERE id = ?"

    row := cnx.DB.QueryRow(query, id)
    var img models.Image
    if err := row.Scan(&img.ID, &img.Service, &img.Item, &img.Image, &img.User, &img.Description, &img.Date, &img.Notice); err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("imagem com ID %d não encontrada", id)
        }
        return nil, fmt.Errorf("erro ao buscar imagem: %v", err)
    }

    return &img, nil
}

func (cnx *ImageRepository) UpdateImage(img models.Image) error {
    query := "UPDATE IMAGEM SET atendimento = ?, item = ?, imagem = ?, usuario = ?, descricao = ?, data = ?, aviso = ? WHERE id = ?"

    result, err := cnx.DB.Exec(query, img.Service, img.Item, img.Image, img.User, img.Description, img.Date, img.Notice, img.ID)
    if err != nil {
        return fmt.Errorf("erro ao atualizar imagem: %v", err)
    }

    affectedRows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("erro ao retornar linhas afetadas: %v", err)
    }

    if affectedRows == 0 {
        return fmt.Errorf("nenhuma imagem com id: %d encontrada", img.ID)
    }

    return nil
}

func (cnx *ImageRepository) DeleteImage(id int) error {
    query := "DELETE FROM IMAGEM WHERE id = ?"

    result, err := cnx.DB.Exec(query, id)
    if err != nil {
        return fmt.Errorf("erro ao deletar imagem: %v", err)
    }

    affectedRows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("erro ao retornar linhas afetadas: %v", err)
    }

    if affectedRows == 0 {
        return fmt.Errorf("nenhuma imagem com id: %d encontrada", id)
    }

    return nil
}