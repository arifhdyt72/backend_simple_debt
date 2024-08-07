package pembayaran

import (
	"backend_test_debt/initializer"
	"backend_test_debt/models"
)

type Repository interface {
	Create(data models.Pembayaran) (models.Pembayaran, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) Create(data models.Pembayaran) (models.Pembayaran, error) {
	err := initializer.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
