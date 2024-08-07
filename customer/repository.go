package customer

import (
	"backend_test_debt/initializer"
	"backend_test_debt/models"
)

type Repository interface {
	Create(data models.Customer) (models.Customer, error)
	Update(data models.Customer) (models.Customer, error)
	DeleteData(id int) (models.Customer, error)
	GetAllData(input DatatableInput) ([]models.Customer, error)
	GetCountData(input DatatableInput) (int, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) Create(data models.Customer) (models.Customer, error) {
	err := initializer.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) Update(data models.Customer) (models.Customer, error) {
	err := initializer.DB.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) DeleteData(id int) (models.Customer, error) {
	var customer models.Customer
	err := initializer.DB.First(&customer, id).Error
	if err != nil {
		return customer, err
	}

	err = initializer.DB.Delete(&customer, customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) GetAllData(input DatatableInput) ([]models.Customer, error) {
	if input.Rows == 0 {
		input.Rows = 10
	}

	if input.SortField == "" || input.SortField == "null" {
		input.SortField = "id"
	}

	if input.SortOrder == "1" {
		input.SortOrder = "ASC"
	} else {
		input.SortOrder = "DESC"
	}

	search := "%" + input.Filters + "%"

	var data []models.Customer
	err := initializer.DB.Where("name LIKE ? OR email LIKE ? OR phone_number LIKE ? OR address LIKE ?",
		search, search, search, search).Limit(input.Rows).Offset(input.First).
		Order(input.SortField + " " + input.SortOrder).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetCountData(input DatatableInput) (int, error) {
	if input.Rows == 0 {
		input.Rows = 10
	}

	if input.SortField == "" || input.SortField == "null" {
		input.SortField = "id"
	}

	if input.SortOrder == "1" {
		input.SortOrder = "ASC"
	} else {
		input.SortOrder = "DESC"
	}

	search := "%" + input.Filters + "%"
	var count int64
	var data []models.Customer
	err := initializer.DB.Where("name LIKE ? OR email LIKE ? OR phone_number LIKE ? OR address LIKE ?",
		search, search, search, search).Find(&data).Count(&count).Error
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}
