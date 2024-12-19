package repository

import (
	"SVPWeb/internal/api/models"
)

type ImageRepositoryMock struct{}

func (r *ImageRepositoryMock) CreateImage(img models.Image) error {
	return nil
}

func (r *ImageRepositoryMock) GetImageByID(id int) (*models.Image, error) {
	return &models.Image{ID: uint(id), Description: "Test Image"}, nil
}

func (r *ImageRepositoryMock) UpdateImage(img models.Image) error {
	return nil
}

func (r *ImageRepositoryMock) DeleteImage(id int) error {
	return nil
}
